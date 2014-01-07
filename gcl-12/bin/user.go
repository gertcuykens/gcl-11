package bin

import (
	"appengine/datastore"
	"appengine"
	"encoding/json"
	"crypto/sha1"
	"time"
	"encoding/hex"
)

type User struct {
	Key *datastore.Key `datastore:"-"`
	Group []byte `datastore:"group"`
	Refresh []byte `datastore:"refresh_token"`
	Status string `datastore:"-"`
	Token *Token `datastore:"-"`
	Context appengine.Context `datastore:"-"`
}

func (u *User) Error() string {
	return u.Status
}

func (u *User) Store() error{
	if u.Key == nil {u.Status="Datastore no Key!"; return u}
	key, err := datastore.Put(u.Context, u.Key, u);
	if err != nil {u.Status="Datastore put error! "+err.Error(); return u}
	u.Status="Stored "+key.StringID()+"."
	return nil
}

func (u *User) Get() (err error){
	if u.Key == nil {u.Status="Datastore no Key!"; return u}
	if err = datastore.Get(u.Context, u.Key, u); err != nil {u.Status="Datastore get error! "+err.Error(); return u}
	if err = json.Unmarshal(u.Group, &u.Token.Extra); err != nil {u.Status="Datastore get error! "+err.Error(); return u}
	u.Status="Fetched "+u.Key.StringID()+"."
	return nil
}

func (u *User) Init() (err error){
	if u.Token == nil {u.Status="No token!"; return u}
	u.Key= datastore.NewKey(u.Context, "User", u.Token.Id, 0, nil)
	if u.Group, err = json.Marshal(u.Token.Extra); err != nil {u.Status="Login error! "+err.Error(); return u}
	h := sha1.New()
	e := time.Now().Add(time.Duration(3600)*time.Second)
	a := string(u.Group)+e.String()+SERVER_SECRET
	s := hex.EncodeToString(h.Sum([]byte(a)))
	u.Token.Access = s
	u.Token.Expiry = e
	u.Token.Status="OK"
	u.Status="OK"
	return nil
}

func (u *User) Logout() error{
	u.Status="Out."
	u.Token=nil
	return nil
}

func (u *User) Login(b []byte) error {
	if u.Refresh == nil {u.Status="No refresh token!"; return u}
	if len(u.Refresh) != len(b) {
		u.Status="Refresh not equal!";
		return u
	}
	for i, v := range u.Refresh {
		if v != b[i] {
			u.Status="Refresh not equal!";
			return u
		}
	}
	return nil
}
