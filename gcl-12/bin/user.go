package bin

import (
	"crypto/sha1"
	"time"
	"fmt"
	"net/http"
	"encoding/hex"
	//"errors"
	//"log"
)


type UserError struct {
	Description string
}

func (e UserError) Error() string {
	return e.Description
}

func User(id string, group string) Token{
	h := sha1.New()
	e := time.Now().Add(time.Duration(3600)*time.Second)
	a :=id+group+e.String()+SERVER_SECRET
	s := hex.EncodeToString(h.Sum([]byte(a)))
	//log.Print(a)
	t := Token{
		AccessToken: s,
		Expiry: e,
		Extra : map[string]string{"user":id, "group":group},
	}
	return t
}

func (t *Token) CheckSum() (bool, error) {
	h := sha1.New()
	a := t.Extra["user"]+t.Extra["group"]+t.Expiry.String()+SERVER_SECRET
	s := hex.EncodeToString(h.Sum([]byte(a)))
	//log.Print(a)
	if t.Expired() {return false, UserError{Description:"Error: Expired"}}
	if t.AccessToken != s {return false, UserError{Description:"Error: CheckSum"}}
	return true, nil
}

func Test(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "text/html; charset=utf-8")
	u:=User("gert","admin")
	b, err := u.CheckSum()
    fmt.Fprintf(w, "Token %+v</br> is %t, %s!</br> The time is now %v", u, b, err, time.Now())
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
