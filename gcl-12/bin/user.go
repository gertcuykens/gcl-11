package bin

import (
	"appengine"
	"appengine/datastore"
	"crypto/sha1"
	"time"
	"encoding/hex"
	"net/http"
	"fmt"
	"encoding/json"
	//"log"
)

type User struct {
	Key *datastore.Key `datastore:"-"`
	Group []byte `datastore:"group"`
	Status string `datastore:"-"`
	Token *Token `datastore:"-"`
}

func (u *User) Error() string {
	return u.Status
}

func (u *User) Store(c appengine.Context) error{
	if u.Key == nil {u.Status="Datastore no Key!"; return u}
	key, err := datastore.Put(c, u.Key, u);
	if err != nil {u.Status="Datastore put error! "+err.Error(); return u}
	u.Status="Stored "+key.StringID()+"."
	return nil
}

func (u *User) Get(c appengine.Context) (err error){
	if u.Key == nil {u.Status="Datastore no Key!"; return u}
	err = datastore.Get(c, u.Key, u);
	if err != nil {u.Status="Datastore get error! "+err.Error(); return u}
	err = json.Unmarshal(u.Group,&u.Token.Extra)
	if err != nil {u.Status="Datastore get error! "+err.Error(); return u}
	u.Status="Fetched "+u.Key.StringID()+"."
	return nil
}

func (u *User) Login(c appengine.Context) (err error){
	if u.Token == nil {u.Status="No token!"; return u}
	u.Key= datastore.NewKey(c, "User", u.Token.Extra["key"], 0, nil)
	u.Group, err = json.Marshal(u.Token.Extra)
	if err != nil {u.Status="Login error! "+err.Error(); return u}
	h := sha1.New()
	e := time.Now().Add(time.Duration(3600)*time.Second)
	a := string(u.Group)+e.String()+SERVER_SECRET
	s := hex.EncodeToString(h.Sum([]byte(a)))
	u.Token.AccessToken = s
	u.Token.Expiry = e
	u.Status="In."
	return nil
}

func (u *User) Logout() error{
	u.Status="Out."
	u.Token=nil
	return nil
}

func (u *User) CheckSum() error {
	if u.Token == nil {u.Status="No token!"; return u}
	b, err := json.Marshal(u.Token.Extra)
	if err != nil {u.Status="Token error! "+err.Error(); return u}
	h := sha1.New()
	a := string(b)+u.Token.Expiry.String()+SERVER_SECRET
	s := hex.EncodeToString(h.Sum([]byte(a)))
	if u.Token.Expired() {u.Status="Token expired!"; return u}
	if u.Token.AccessToken != s {u.Status="Token checkSum error!"; return u}
	return nil
}

func Test(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := new(User)
	u.Token= &Token{Extra:map[string]string{"key":"gert","group":"admin"}}
	u.Login(c)
	u.Store(c)
	u.Get(c)
	u.CheckSum()
	//u.Logout()
	//log.Print(u.Key.StringID())
	//log.Print(string(u.Group))
	t, _ :=json.Marshal(u.Token)
	w.Header().Set("Content-type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "User status %+v</br>Token %s</br>The time is now %v", u, string(t), time.Now())
}

/*
func welcome(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        url, _ := user.LoginURL(c, "/")
        fmt.Fprintf(w, `<a href="%s">Sign in google</a>`, url)
        return
    }
    url, _ := user.LogoutURL(c, "/")
    fmt.Fprintf(w, `Welcome, %s! (<a href="%s">Sign out google</a>)`, u, url)
}

func welcome2(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u, err := user.CurrentOAuth(c, "https://www.googleapis.com/auth/userinfo.email")
    if err != nil {
        http.Error(w, "OAuth Authorization header required", http.StatusUnauthorized)
        return
    }
    if !u.Admin {
        http.Error(w, "Admin login only", http.StatusUnauthorized)
        return
    }
    fmt.Fprintf(w, `Welcome, %s!`, u)
}
*/

/*
func init() {
    http.HandleFunc("/_ah/login_required", openIdHandler)
}

func openIdHandler(w http.ResponseWriter, r *http.Request) {
    // ...
}
*/
