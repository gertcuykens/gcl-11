package bin

import (
	"net/http"
	"appengine"
	"encoding/json"
	"fmt"
	"time"
)

func (s *Service) UserCreate(r *http.Request, req *Token, resp *Token) (err error) {
	p:=Property{Key:"group",Value:"user"}
	u := new(User)
	u.Context = appengine.NewContext(r)
	u.Token = &Token{Id:req.Id, Extra:[]Property{p}}
	u.Refresh=[]byte(req.Refresh)
	if err = u.Init(); err !=nil {return}
	if err = u.Store(); err !=nil {return}
	resp = u.Token
	return err
}

func (s *Service) UserRefresh(r *http.Request, req *Token, resp *Token) (err error) {
	p := Property{Key:"group", Value:"user"}
	u := new(User)
	u.Context = appengine.NewContext(r)
	u.Token= &Token{Id:req.Id, Extra:[]Property{p}}
	if err = u.Init(); err !=nil {return}
	if err = u.Get(); err !=nil {return}
	if err = u.Login([]byte(req.Refresh)); err !=nil {return}
	resp = u.Token
	return err
}

func (s *Service) UserToken(r *http.Request, req *Token, resp *Response) (err error) {
	u := new(User)
	u.Context = appengine.NewContext(r)
	u.Token = &Token{
		Access: req.Access,
		Expiry: req.Expiry,
		Extra: req.Extra,
    }
	if err = u.Token.CheckSum(); err !=nil {return}
	resp.Message="OK"
	return err
}

func Test(w http.ResponseWriter, r *http.Request) {
	p := Property{Key:"group", Value:"user"}
	u := new(User)
	u.Context = appengine.NewContext(r)
	u.Token= &Token{Id:"gert", Extra:[]Property{p}}
	u.Refresh=[]byte("password")
	u.Init()
	u.Store()
	u.Get()
	u.Login([]byte("password"))
	u.Token.CheckSum()
	//u.Logout()
	//log.Print(u.Key.StringID())
	//log.Print(string(u.Group))
	//log.Print(u.Equals([]byte("password")))
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
