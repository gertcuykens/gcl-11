package bin

import "github.com/crhym3/go-endpoints/endpoints"

type Service struct {}

func init() {
	service := &Service{}
	api, err := endpoints.RegisterService(service, "rest", "v0", "API", true)
	if err != nil {panic(err.Error())}
	info1 := api.MethodByName("GoogleUserService").Info()
	info1.Name, info1.HttpMethod, info1.Path, info1.Desc, info1.Scopes = "google.user", "POST", "google/user", "Oauth2 google user.", google_scopes
	info2 := api.MethodByName("GoogleRevokeService").Info()
	info2.Name, info2.HttpMethod, info2.Path, info2.Desc = "google.revoke", "POST", "google/revoke", "Oauth2 google revoke."
	info3 := api.MethodByName("GooglePurchasesService").Info()
	info3.Name, info3.HttpMethod, info3.Path, info3.Desc, info3.Scopes= "google.purchases", "POST", "google/purchases", "Oauth2 google purchases.", google_scopes
	info4 := api.MethodByName("GoogleStorageService").Info()
	info4.Name, info4.HttpMethod, info4.Path, info4.Desc, info4.Scopes = "google.storage", "GET", "google/storage", "Oauth2 google storage.", google_scopes
	endpoints.HandleHttp()
}

/*

	info3 := api.MethodByName("TwitterOauth").Info()
	info3.Name, info3.HttpMethod, info3.Path, info3.Desc = "twitter.oauth", "GET", "twitter/oauth", "Oauth url."
	info4 := api.MethodByName("TwitterCallback").Info()
	info4.Name, info4.HttpMethod, info4.Path, info4.Desc = "twitter.callback", "GET", "twitter/callback", "Oauth callback."
	info5 := api.MethodByName("TwitterOauthOob").Info()
	info5.Name, info5.HttpMethod, info5.Path, info5.Desc = "twitter.oauth.oob", "GET", "twitter/oauth/oob", "Oauth url oob."
	info6 := api.MethodByName("TwitterCallbackOob").Info()
	info6.Name, info6.HttpMethod, info6.Path, info6.Desc = "twitter.callback.oob", "POST", "twitter/callback/oob", "Oauth callback oob."
	info7 := api.MethodByName("LinkedInOauth").Info()
	info7.Name, info7.HttpMethod, info7.Path, info7.Desc = "linkedin.oauth", "GET", "linkedin/oauth", "Oauth url."
	info8 := api.MethodByName("LinkedInCallback").Info()
	info8.Name, info8.HttpMethod, info8.Path, info8.Desc = "linkedin.callback", "GET", "linkedin/callback", "Oauth callback."
	info9 := api.MethodByName("FacebookCallback").Info()
	info9.Name, info9.HttpMethod, info9.Path, info9.Desc = "facebook.callback", "GET", "facebook/callback", "Oauth callback."
	info10 := api.MethodByName("Register").Info()
	info10.Name, info10.HttpMethod, info10.Path, info10.Desc = "register", "POST", "register", "Register user."
	info11 := api.MethodByName("CheckSum").Info()
	info11.Name, info11.HttpMethod, info11.Path, info11.Desc = "checksum", "POST", "checkSum", "Check token."

	http.HandleFunc("/bin/test", Test)
import (
    "fmt"
    "net/http"
)

func init() {
    //http.HandleFunc("/", redirect("/home/index.html"))
    //http.HandleFunc("/favicon.ico", redirect("/home/favicon.ico"))
    http.HandleFunc("/bin/hello", hello)
    http.HandleFunc("/bin/contact", contact)
    http.HandleFunc("/bin/welcome", welcome)
    
}

func redirect(path string) func(http.ResponseWriter, *http.Request) {
  return func (w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, path, http.StatusMovedPermanently)
  }
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}
//c.Infof("-------------------")
*/
