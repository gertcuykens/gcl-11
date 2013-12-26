package bin

import (
	"net/http"
    "encoding/json"
	"github.com/crhym3/go-endpoints/endpoints"
    "appengine/urlfetch"
    "io/ioutil"
)

func GetUser(c endpoints.Context, token string) (*User, error) {
    var u *User
    httpClient := urlfetch.Client(c)
    resp, err := httpClient.Get("https://graph.facebook.com/me?access_token="+token)
    if err != nil {}
    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {}
    err = json.Unmarshal(b, &u)
    if err != nil {}
    resp.Body.Close()
    return u, err
}

func (s *Service) Welcome(r *http.Request, req *Request, resp *Response) error {
    c := endpoints.NewContext(r)
    g, err := endpoints.CurrentUser(c, SCOPES, AUDIENCES, CLIENTIDS)
    f, err := GetUser(c, req.Token)
    if err == nil {resp.Message="Google:"+g.String()+" Facebook:"+f.Name}
    //c.Infof("-------------------")
    return err
}

func init() {
  service := &Service{}
  api, err := endpoints.RegisterService(service, "rest", "v0", "Login API", true)
  if err != nil {panic(err.Error())}
  info1 := api.MethodByName("Welcome").Info()
  info1.Name, info1.HttpMethod, info1.Path, info1.Desc, info1.Scopes = "login", "POST", "welcome", "Login.", SCOPES
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
*/
