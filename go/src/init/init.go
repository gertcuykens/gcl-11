package init

import (
	"github.com/crhym3/go-endpoints/endpoints"
	"code.google.com/p/goauth2/oauth"
	"time"
)

type Service struct {}

const PACKAGE string = ""

const WEB_CLIENT_ID string = "522156758812-09f5qv0e4gqjdjqfocerqcud5m5jutau.apps.googleusercontent.com"
const ANDROID_CLIENT_ID string = ""

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
	api, err := endpoints.RegisterService(service, "rest", "v0", "API", true)
	if err != nil {panic(err.Error())}
	rpc1(api)
	rpc2(api)
	rpc3(api)
	endpoints.HandleHttp()
}

func rpc1(api *endpoints.RpcService) {
	scope := []string{"https://www.googleapis.com/auth/userinfo.email"}
	info1 := api.MethodByName("GoogleCallback").Info()
	info1.Name, info1.HttpMethod, info1.Path, info1.Desc, info1.Scopes = "google.callback", "POST", "google/callback", "Oauth callback.", scope
	info2 := api.MethodByName("GoogleRevoke").Info()
	info2.Name, info2.HttpMethod, info2.Path, info2.Desc = "google.revoke", "POST", "google/revoke", "Oauth revoke."
	info3 := api.MethodByName("TwitterOauth").Info()
	info3.Name, info3.HttpMethod, info3.Path, info3.Desc = "twitter.oauth", "GET", "twitter/oauth", "Oauth url."
	info4 := api.MethodByName("TwitterCallback").Info()
	info4.Name, info4.HttpMethod, info4.Path, info4.Desc = "twitter.callback", "GET", "twitter/callback", "Oauth callback."
	info5 := api.MethodByName("TwitterOauthOob").Info()
	info5.Name, info5.HttpMethod, info5.Path, info5.Desc = "twitter.oauth.oob", "GET", "twitter/oauth/oob", "Oauth url oob."
	info6 := api.MethodByName("TwitterCallbackOob").Info()
	info6.Name, info6.HttpMethod, info6.Path, info6.Desc = "twitter.callback.oob", "POST", "twitter/callback/oob", "Oauth callback oob."
	//info7 := api.MethodByName("LinkedInOauth").Info()
	//info7.Name, info7.HttpMethod, info7.Path, info7.Desc = "linkedin.oauth", "GET", "linkedin/oauth", "Oauth url."
	//info8 := api.MethodByName("LinkedInCallback").Info()
	//info8.Name, info8.HttpMethod, info8.Path, info8.Desc = "linkedin.callback", "GET", "linkedin/callback", "Oauth callback."
	info9 := api.MethodByName("FacebookCallback").Info()
	info9.Name, info9.HttpMethod, info9.Path, info9.Desc = "facebook.callback", "POST", "facebook/callback", "Oauth callback."
	//info10 := api.MethodByName("Register").Info()
	//info10.Name, info10.HttpMethod, info10.Path, info10.Desc = "register", "POST", "register", "Register user."
	//info11 := api.MethodByName("CheckSum").Info()
	//info11.Name, info11.HttpMethod, info11.Path, info11.Desc = "checksum", "POST", "checkSum", "Check token."
}

func rpc2(api *endpoints.RpcService) {
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

func rpc3(api *endpoints.RpcService) {
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
