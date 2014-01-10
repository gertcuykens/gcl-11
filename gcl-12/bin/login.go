package bin

import (
	"net/http"
	"encoding/json"
	"fmt"
	"time"
	"appengine/urlfetch"
	"github.com/crhym3/go-endpoints/endpoints"
)

func (s *Service) UserCreate(r *http.Request, req *Token, resp *Token) (err error) {
	c := endpoints.NewContext(r)
	//g, err := endpoints.CurrentUser(c, google_scopes, audiences, clientids);
	//if err != nil {return}
	//......check for admin..........
	u := new(User)
	u.Token = req
	u.Token.Context = c
	u.Token.Client = urlfetch.Client(c)
	if err = u.Token.SelectId(); err !=nil {return}
	if err = u.Init(); err !=nil {return}
	if err = u.Store(); err !=nil {return}
	*resp = *u.Token
	return err
}

func (s *Service) UserRefresh(r *http.Request, req *Token, resp *Token) (err error) {
	p := Property{Key:"group", Value:"user"}
	u := new(User)
	u.Token.Context = endpoints.NewContext(r)
	u.Token= &Token{Id:req.Id, Extra:[]Property{p}}
	if err = u.Init(); err !=nil {return}
	if err = u.Get(); err !=nil {return}
    if u.Token.Refresh != req.Refresh {u.Token.Status="Wrong refresh token!"; return u}
	*resp = *u.Token
	return err
}

func (s *Service) UserToken(r *http.Request, req *Token, resp *Token) (err error) {
	u := new(User)
	u.Token = req
	u.Token.Context = endpoints.NewContext(r)
	if err = u.Token.CheckSum(); err !=nil {return err}
	u.Token.Status="OK"
	*resp = *u.Token
	return err
}

func Test(w http.ResponseWriter, r *http.Request) {
	p := Property{Key:"group", Value:"user"}
	u := new(User)
	u.Token.Context = endpoints.NewContext(r)
	u.Token= &Token{Id:1, Type:"test", Refresh:"password", Extra:[]Property{p}}
	u.Init()
	u.Store()
	u.Get()
	u.Token.CheckSum()
	//u.Logout()
	//log.Print(u.Key.StringID())
	//log.Print(string(u.Group))
	//log.Print(u.Equals([]byte("password")))
	j, _ :=json.Marshal(u)
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Status: %+v</br> %s</br>The time is now %v", u, string(j), time.Now())
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
