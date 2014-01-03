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
	api, err := endpoints.RegisterService(service, "rest", "v0", "Login API", true)
	if err != nil {panic(err.Error())}
	info1 := api.MethodByName("GoogleUser").Info()
	info1.Name, info1.HttpMethod, info1.Path, info1.Desc, info1.Scopes = "login", "POST", "welcome", "Log in.", SCOPES
	info2 := api.MethodByName("GoogleRevoke").Info()
	info2.Name, info2.HttpMethod, info2.Path, info2.Desc = "logout", "POST", "bye", "Log out."
	info3 := api.MethodByName("TwitterOauth").Info()
	info3.Name, info3.HttpMethod, info3.Path, info3.Desc = "oauth", "GET", "oauth", "Oauth url."
	info4 := api.MethodByName("TwitterCallback").Info()
	info4.Name, info4.HttpMethod, info4.Path, info4.Desc = "callback", "GET", "oauth_callback", "Oauth callback url."
	info5 := api.MethodByName("TwitterOauthOob").Info()
	info5.Name, info5.HttpMethod, info5.Path, info5.Desc = "oauth_oob", "GET", "oauth_oob", "Oauth url."
	info6 := api.MethodByName("TwitterCallbackOob").Info()
	info6.Name, info6.HttpMethod, info6.Path, info6.Desc = "callback_oob", "POST", "oauth_callback_oob", "Oauth callback url."
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
