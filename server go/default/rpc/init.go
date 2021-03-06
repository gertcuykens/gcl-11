package rpc

import (
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"
	"code.google.com/p/goauth2/oauth"
	"time"
	"cloud"
)

type Service struct {Status string `json:"error"`}

func (s *Service) Error() string {return s.Status}

func (s *Service) GetHeat(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
	d := cloud.DataStore {Entity: &cloud.Entity{}, Request: r}
	d.GetHeat(m.List[0].Event, m.List[0].Division, m.List[0].Heat)
	*resp = *d.Entity
	return nil
}

func (s *Service) GetFirst(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
	d := cloud.DataStore {Entity: &cloud.Entity{}, Request: r}
	d.GetFirst(m.List[0].Event)
	*resp = *d.Entity
	return nil
}

func (s *Service) Get(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
	d := cloud.DataStore {Entity: &cloud.Entity{}, Request: r}
	d.Get(m.List[0].Event, m.List[0].Id)
	*resp = *d.Entity
	return nil
}

func (s *Service) Put(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
	d := cloud.DataStore {Entity: m, Request: r}
	d.Put()
	*resp = *d.Entity
	return nil
}

func (s *Service) Delete(r *http.Request, m *cloud.Entity, _ *cloud.Entity) error {
	d := cloud.DataStore {Entity: m, Request: r}
	d.Delete()
	return nil
}

func (s *Service) Truncate(r *http.Request, m *cloud.Entity, _ *cloud.Entity) error {
	d := cloud.DataStore {Entity: m, Request: r}
	d.Truncate(m.List[0].Event)
	return nil
}

func (s *Service) Editor(r *http.Request,  _ *cloud.Entity, _ *cloud.Entity) error {
	d := cloud.DataStore {Request: r}
	if err := d.Editor(); err !=nil {return err}
	return nil
}

func (s *Service) GetTrickList(r *http.Request, m *cloud.TrickList, resp *cloud.TrickList) error {
	d := cloud.TrickStore {Entity: &cloud.TrickList{}, Request: r}
	d.GetTrickList()
	*resp = *d.Entity
	return nil
}

func (s *Service) PutTrickName(r *http.Request, m *cloud.TrickList, resp *cloud.TrickList) error {
	d := cloud.TrickStore {Entity: m, Request: r}
	d.PutTrickName()
	*resp = *d.Entity
	return nil
}

func (s *Service) DeleteTrickName(r *http.Request, m *cloud.TrickList, resp *cloud.TrickList) error {
	d := cloud.TrickStore {Entity: m, Request: r}
	d.DeleteTrickName()
	*resp = *d.Entity
	return nil
}

const PACKAGE string = ""

const WEB_CLIENT_ID string = "522156758812-09f5qv0e4gqjdjqfocerqcud5m5jutau.apps.googleusercontent.com"
const ANDROID_CLIENT_ID string = "522156758812-29jkcaiofrismobslc4ioop1dvfhhgoi.apps.googleusercontent.com"

const SERVER_CLIENT_ID string = ""
const SERVER_SECRET string = ""
const SERVER_REFRESH string = ""

const SCOPE1 string = "https://www.googleapis.com/auth/userinfo.email"
const SCOPE2 string = "https://www.googleapis.com/auth/devstorage.full_control"
const SCOPE3 string = "https://www.googleapis.com/auth/androidpublisher"

var CLIENTIDS = []string{WEB_CLIENT_ID, ANDROID_CLIENT_ID, endpoints.ApiExplorerClientId}
var AUDIENCES = []string{WEB_CLIENT_ID}
var SCOPES = []string{SCOPE1}

var CONFIG = &oauth.Config{
	ClientId: SERVER_CLIENT_ID,
	ClientSecret: SERVER_SECRET,
	Scope:        SCOPE1+" "+SCOPE2+" "+SCOPE3,
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://accounts.google.com/o/oauth2/token",
}

var TOKEN = &oauth.Token{
	AccessToken: "",
	RefreshToken: SERVER_REFRESH,
	Expiry: time.Now(),
	Extra: nil,
}

var TRANS = &oauth.Transport{
	Token: TOKEN,
	Config: CONFIG,
}

func init() {
	service := &Service{}
	api, err := endpoints.RegisterService(service, "service", "v0", "API", true)
	if err != nil {panic(err.Error())}
	scope := []string{}
	rpc(api, scope, "GetTrickList", "trickstore.getTrickList", "POST", "trickstore/getTrickList", "Get trick list.")
	rpc(api, scope, "PutTrickName", "trickstore.putTrickName", "POST", "trickstore/putTrickName", "Put trick name.")
	rpc(api, scope, "DeleteTrickName", "trickstore.deleteTrickName", "POST", "trickstore/deleteTrickName", "Delete trick name.")

	rpc(api, scope, "GetHeat", "datastore.getHeat", "POST", "datastore/getHeat", "Get heat.")
	rpc(api, scope, "GetFirst", "datastore.getFirst", "POST", "datastore/getFirst", "Get first.")
	rpc(api, scope, "Get", "datastore.get", "POST", "datastore/get", "Get.")
	rpc(api, scope, "Put", "datastore.put", "POST", "datastore/put", "Put.")
	rpc(api, scope, "Delete", "datastore.delete", "POST", "datastore/delete", "Delete.")
	rpc(api, scope, "Truncate", "datastore.truncate", "POST", "datastore/truncate", "Truncate.")

	rpc(api, scope, "Editor", "editor", "POST", "editor", "Editor.")
	endpoints.HandleHttp()
}

func rpc(api *endpoints.RpcService, scope []string, a ...string) {
	info := api.MethodByName(a[0]).Info()
	info.Name, info.HttpMethod, info.Path, info.Desc, info.Scopes= a[1], a[2], a[3], a[4], scope
}

//scope := []string{"https://www.googleapis.com/auth/userinfo.email"}
//rpc(api, scope, "GoogleCallback", "google.callback", "POST", "google/callback", "Oauth callback.")
//rpc(api, scope, "GoogleRevoke", "google.revoke", "POST", "google/revoke", "Oauth revoke.")
//rpc(api, scope, "TwitterOauth", "twitter.oauth", "GET", "twitter/oauth", "Oauth url.")
//rpc(api, scope, "TwitterCallback", "twitter.callback", "GET", "twitter/callback", "Oauth callback.")
//rpc(api, scope, "TwitterOauthOob", "twitter.oauth.oob", "GET", "twitter/oauth/oob", "Oauth url oob.")
//rpc(api, scope, "TwitterCallbackOob", "twitter.callback.oob", "POST", "twitter/callback/oob", "Oauth callback oob.")
//rpc(api, scope, "LinkedInOauth", "linkedin.oauth", "GET", "linkedin/oauth", "Oauth url.")
//rpc(api, scope, "LinkedInCallback". "linkedin.callback", "GET", "linkedin/callback", "Oauth callback.")
//rpc(api, scope, "FacebookCallback", "facebook.callback", "POST", "facebook/callback", "Oauth callback.")

/*
func rpc3(api *endpoints.RpcService) {
	scope := []string{"https://www.googleapis.com/auth/userinfo.email"}
	info1 := api.MethodByName("List").Info()
	info1.Name, info1.HttpMethod, info1.Path, info1.Desc = "greetings.listGreeting", "GET", "response", "List of greetings."
	info2 := api.MethodByName("Store").Info()
	info2.Name, info2.HttpMethod, info2.Path, info2.Desc, info2.Scopes = "greetings.datastore", "GET", "greetings/datastore", "Datastore.", scope
	info3 := api.MethodByName("Soap").Info()
	info3.Name, info3.HttpMethod, info3.Path, info3.Desc = "greetings.soap", "GET", "greetings/soap", "Soap."
	info4 := api.MethodByName("Login").Info()
	info4.Name, info4.HttpMethod, info4.Path, info4.Desc, info4.Scopes= "greetings.authed", "POST", "greetings/authed", "Login.", scope
	info5 := api.MethodByName("Id").Info()
	info5.Name, info5.HttpMethod, info5.Path, info5.Desc= "greetings.getGreeting", "GET", "response/{id}", "Get greeting."
	info6 := api.MethodByName("Multiply").Info()
	info6.Name, info6.HttpMethod, info6.Path, info6.Desc= "greetings.multiply", "POST", "response", "Multiply greeting."
}

func rpc4(api *endpoints.RpcService) {
	scope := []string{"https://www.googleapis.com/auth/userinfo.email"}
	//info1 := api.MethodByName("GoogleUserService").Info()
	//info1.Name, info1.HttpMethod, info1.Path, info1.Desc, info1.Scopes = "google.user", "POST", "google/user", "Oauth2 google user.", scope
	//info2 := api.MethodByName("GoogleRevokeService").Info()
	//info2.Name, info2.HttpMethod, info2.Path, info2.Desc = "google.revoke", "POST", "google/revoke", "Oauth2 google revoke."
	info3 := api.MethodByName("GooglePurchases").Info()
	info3.Name, info3.HttpMethod, info3.Path, info3.Desc, info3.Scopes= "google.purchases", "POST", "google/purchases", "Oauth2 google purchases.", scope
	info4 := api.MethodByName("GoogleStorage").Info()
	info4.Name, info4.HttpMethod, info4.Path, info4.Desc, info4.Scopes = "google.storage", "GET", "google/storage", "Oauth2 google storage.", scope
}
*/

/*
import (
    "fmt"
    "net/http"
    "fmt"
    "appengine"
    "appengine/user"
)

func init() {
    //http.HandleFunc("/", redirect("/home/index.html"))
    //http.HandleFunc("/favicon.ico", redirect("/home/favicon.ico"))
    //http.HandleFunc("/_ah/login_required", openIdHandler)
    http.HandleFunc("/bin/hello", hello)
    http.HandleFunc("/bin/contact", contact)
    http.HandleFunc("/bin/welcome", welcome)
    http.HandleFunc("/", welcome)
}

func redirect(path string) func(http.ResponseWriter, *http.Request) {
  return func (w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, path, http.StatusMovedPermanently)
  }
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	url, _ := user.LogoutURL(c, "/")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
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

func openIdHandler(w http.ResponseWriter, r *http.Request) {
    // ...
}
*/

//rpc(api, scope, "Get2", "datastore.get2", "POST", "datastore/get2", "Get2.")
//rpc(api, scope, "Put2", "datastore.put2", "POST", "datastore/put2", "Put2.")
