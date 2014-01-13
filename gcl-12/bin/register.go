package bin

import (
	"net/http"
	"time"
	"appengine/urlfetch"
	"github.com/crhym3/go-endpoints/endpoints"
	"crypto/sha1"
	"encoding/hex"
	"appengine/datastore"
	"strconv"
)

func (s *Service) Register(r *http.Request, req *Token, resp *Token) (err error) {
	c := endpoints.NewContext(r)

	u := new(User)
	u.Token = req
	u.Token.Context = c
	u.Token.Client = urlfetch.Client(c)
	if err = u.Token.SelectId(); err !=nil {return}

	var id = Property{}
	if u.Token.Id != 0 {
		u.Key= datastore.NewKey(u.Token.Context, u.Token.Type, "", u.Token.Id, nil)
		id.Key="id"
		id.Value= strconv.FormatInt(u.Token.Id,10)
	}
	if u.Token.Email != "" {
		u.Key= datastore.NewKey(u.Token.Context, u.Token.Type, u.Token.Email, 0, nil)
		id.Key="id"
		id.Value=u.Token.Email
	}

	if err = u.Get(); err != nil && err != datastore.ErrNoSuchEntity {return}
	err=nil

    if u.Group == nil {
		var g = []Property{}
		var group = Property{
			Key:"group",
			Value:"user",
		}
		g = append(g, id)
		g = append(g, group)
		u.Group=g
		if err = u.Put(); err !=nil {return}
	}
	u.Token.Extra = u.Group

	h := sha1.New()
	e := time.Now().Add(time.Duration(3600)*time.Second)
	a := u.Token.Extra[0].Value+u.Token.Extra[1].Value+e.String()+SERVER_SECRET
	x := hex.EncodeToString(h.Sum([]byte(a)))
	u.Token.Access = x
	u.Token.Expiry = e
	u.Token.Status = "OK"
	*resp = *u.Token
	return
}

func (s *Service) CheckSum(r *http.Request, req *Token, resp *Token) (err error) {
	u := new(User)
	u.Token = req
	u.Token.Context = endpoints.NewContext(r)
	if err = u.Token.CheckSum(); err !=nil {return err}
	u.Token.Status="OK"
	*resp = *u.Token
	return err
}

/*
	"encoding/json"
	"fmt"

func Test(w http.ResponseWriter, r *http.Request) {
	p := Property{Key:"group", Value:"user"}
	u := new(User)
	u.Token.Context = endpoints.NewContext(r)
	u.Token= &Token{Id:1, Type:"test", Refresh:"password", Extra:[]Property{p}}
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

	func (u *User) Init() (err error){
	if u.Token == nil {u.Token.Status="No token!"; return u}

	u.Extra = u.Token.Extra
	//u.Refresh=[]byte(u.Token.Refresh)

	return nil
}

//Refresh []byte `datastore:"refresh"`

	//g, err := endpoints.CurrentUser(c, google_scopes, audiences, clientids);
	//if err != nil {return}
	//......check for admin..........

func (s *Service) UserRefresh(r *http.Request, req *Token, resp *Token) (err error) {
	p := Property{Key:"group", Value:"user"}
	u := new(User)
	u.Token.Context = endpoints.NewContext(r)
	u.Token = &Token{Id:req.Id, Extra:[]Property{p}}
	if err = u.Init(); err != nil {return}
	if err = u.Get(); err != nil {return}
    if u.Token.Refresh != req.Refresh {u.Token.Status="Wrong refresh token!"; return u}
	*resp = *u.Token
	return err
}

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

/*
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
*/

