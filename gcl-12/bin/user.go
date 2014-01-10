package bin

import (
	"appengine/datastore"
	"crypto/sha1"
	"time"
	"encoding/hex"
)

type User struct {
	Key *datastore.Key `datastore:"-"`
	Type string `datastore:"type"`
	Refresh []byte `datastore:"refresh"`
	Extra []Property `datastore:"extra"`
	Token *Token `datastore:"-"`
}

func (u *User) Error() string {
	return u.Token.Status
}

func (u *User) Store() error{
	if u.Key == nil {u.Token.Status="Datastore no Key!"; return u}
	key, err := datastore.Put(u.Token.Context, u.Key, u);
	if err != nil {u.Token.Status="Datastore put error! "+err.Error(); return u}
	u.Token.Status="Stored "+key.StringID()+"."
	return nil
}

func (u *User) Get() (err error){
	if u.Key == nil {u.Token.Status="Datastore no Key!"; return u}
	if err = datastore.Get(u.Token.Context, u.Key, u); err != nil {u.Token.Status="Datastore get error! "+err.Error(); return u}
	u.Token.Status="Fetched "+u.Key.StringID()+"."
	return nil
}

func (u *User) Init() (err error){
	if u.Token == nil {u.Token.Status="No token!"; return u}
	u.Key= datastore.NewKey(u.Token.Context, u.Token.Type, "", u.Token.Id, nil)
	u.Type= u.Token.Type
	u.Extra=u.Token.Extra
	u.Refresh=[]byte(u.Token.Refresh)
	h := sha1.New()
	e := time.Now().Add(time.Duration(3600)*time.Second)
	a := string(byte(u.Token.Id))+e.String()+SERVER_SECRET
	s := hex.EncodeToString(h.Sum([]byte(a)))
	u.Token.Access = s
	u.Token.Expiry = e
	u.Token.Status="OK"
	return nil
}

func (u *User) Logout() error{
	u.Token.Status="Out."
	return nil
}

func (u *User) Login() error {
	b:=[]byte(u.Token.Refresh)
	if u.Refresh == nil {u.Token.Status="No refresh token!"; return u}
	if len(u.Refresh) != len(b) {
		u.Token.Status="Refresh not equal!";
		return u
	}
	for i, v := range u.Refresh {
		if v != b[i] {
			u.Token.Status="Refresh not equal!";
			return u
		}
	}
	return nil
}
