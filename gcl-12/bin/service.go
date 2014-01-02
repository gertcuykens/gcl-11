package bin

import (
	"net/http"
    "encoding/json"
	"github.com/crhym3/go-endpoints/endpoints"
    "appengine/urlfetch"
    "io/ioutil"
)

func (s *Service) Oauth(r *http.Request, req *NoRequest, resp *ResponseOauth) error {
	c := endpoints.NewContext(r)
	consumer.HttpClient=urlfetch.Client(c)
	requestToken, url, err := TwitterOauth()
	resp.RequestToken=requestToken
	resp.Url=url
	//resp.MessageOauth="{token:\""+requestToken.Token+"\", secret:\""+requestToken.Secret+"\", url:\""+url+"\"}"
	return err
}

func (s *Service) Twitter(r *http.Request, req *RequestOauth, resp *Response) error {
	c := endpoints.NewContext(r)
	consumer.HttpClient=urlfetch.Client(c)
	b, err := TwitterApi(req.RequestToken, req.VerificationCode)
	resp.Message=b
	return err
}

func GetUser(c endpoints.Context, access_token string) (*User, error) {
    var u *User
    httpClient := urlfetch.Client(c)
    resp, err := httpClient.Get("https://graph.facebook.com/me?access_token="+access_token)
    b, err := ioutil.ReadAll(resp.Body)
    err = json.Unmarshal(b, &u)
    resp.Body.Close()
    return u, err
}

func RevokeUser(c endpoints.Context, access_token string) (string, error){
    httpClient := urlfetch.Client(c)
    resp, err := httpClient.Get("https://accounts.google.com/o/oauth2/revoke?token="+access_token)
	defer resp.Body.Close()
    b, err := ioutil.ReadAll(resp.Body)
    return string(b),err
}

func (s *Service) Welcome(r *http.Request, req *Request, resp *Response) error {
    c := endpoints.NewContext(r)
    g, err := endpoints.CurrentUser(c, SCOPES, AUDIENCES, CLIENTIDS)
    f, err := GetUser(c, req.Access_token)
    if err == nil {resp.Message="Google:"+g.String()+" Facebook:"+f.Name}
    return err
}

func (s *Service) Bye(r *http.Request, req *Request, resp *Response) error {
    c := endpoints.NewContext(r)
    b, err := RevokeUser(c, req.Access_token)
    if err == nil {resp.Message=b}
    return err
}

func init() {
	consumer.Debug(true)
	service := &Service{}
	api, err := endpoints.RegisterService(service, "rest", "v0", "Login API", true)
	if err != nil {panic(err.Error())}
	info1 := api.MethodByName("Welcome").Info()
	info1.Name, info1.HttpMethod, info1.Path, info1.Desc, info1.Scopes = "login", "POST", "welcome", "Log in.", SCOPES
	info2 := api.MethodByName("Bye").Info()
	info2.Name, info2.HttpMethod, info2.Path, info2.Desc = "logout", "POST", "bye", "Log out."
	info3 := api.MethodByName("Oauth").Info()
	info3.Name, info3.HttpMethod, info3.Path, info3.Desc = "oauth", "GET", "oauth", "Oauth url."
	info4 := api.MethodByName("Twitter").Info()
	info4.Name, info4.HttpMethod, info4.Path, info4.Desc = "twitter", "POST", "twitter", "Twitter."
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
