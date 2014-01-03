package bin

import (
	"github.com/crhym3/go-endpoints/endpoints"
	"github.com/mrjones/oauth"
)

type Service struct {}
type NoRequest struct {}
var consumer = oauth.NewConsumer(TWITTER_ID, TWITTER_SECRET, TWITTER_SERVER)

func init() {
	consumer.Debug(true)
	service := &Service{}
	api, err := endpoints.RegisterService(service, "rest", "v0", "API", true)
	if err != nil {panic(err.Error())}
	info1 := api.MethodByName("GoogleUser").Info()
	info1.Name, info1.HttpMethod, info1.Path, info1.Desc, info1.Scopes = "google.user", "POST", "google/user", "Sign in.", SCOPES
	info2 := api.MethodByName("GoogleRevoke").Info()
	info2.Name, info2.HttpMethod, info2.Path, info2.Desc = "google.revoke", "POST", "google/revoke", "Sign out."
	info3 := api.MethodByName("TwitterOauth").Info()
	info3.Name, info3.HttpMethod, info3.Path, info3.Desc = "twitter.oauth", "GET", "twitter/oauth", "Oauth url."
	info4 := api.MethodByName("TwitterCallback").Info()
	info4.Name, info4.HttpMethod, info4.Path, info4.Desc = "twitter.callback", "GET", "twitter/callback", "Oauth callback url."
	info5 := api.MethodByName("TwitterOauthOob").Info()
	info5.Name, info5.HttpMethod, info5.Path, info5.Desc = "twitter.oob", "GET", "twitter/oob", "Oauth url oob."
	info6 := api.MethodByName("TwitterCallbackOob").Info()
	info6.Name, info6.HttpMethod, info6.Path, info6.Desc = "twitter.callback.oob", "POST", "twitter/callback", "Oauth callback url oob."
	/*info7 := api.MethodByName("LinkedInOauth2").Info()
	info7.Name, info7.HttpMethod, info7.Path, info7.Desc = "oauth2", "GET", "oauth2", "Oauth url."
	info8 := api.MethodByName("LinkedInCallback2").Info()
	info8.Name, info8.HttpMethod, info8.Path, info8.Desc = "oauth2", "GET", "oauth2_callback", "Oauth callback url."*/
	info9 := api.MethodByName("FacebookUser").Info()
	info9.Name, info9.HttpMethod, info9.Path, info9.Desc = "facebook.user", "GET", "facebook/user", "Facebook user."
	endpoints.HandleHttp()
}

/*
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
