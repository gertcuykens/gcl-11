package rpc

import "github.com/crhym3/go-endpoints/endpoints"

type Q struct{}

type Service struct {Status string `json:"error"`}

func (s *Service) Error() string {return s.Status}

func init() {
	service := &Service{}
	api, err := endpoints.RegisterService(service, "Oauth", "v0", "API", true)
	if err != nil {panic(err.Error())}
	scope := []string{"https://www.googleapis.com/auth/userinfo.email","https://www.googleapis.com/auth/plus.me"}
	rpc(api, scope, "Google", "Google login.")
	rpc(api, []string{}, "Facebook", "Facebook login.")
	endpoints.HandleHttp()
}

func rpc(api *endpoints.RpcService, scope []string, a ...string) {
	info := api.MethodByName(a[0]).Info()
	info.HttpMethod = "POST"
	info.Name = a[0]
	info.Path = a[0]
	info.Desc = a[1]
	info.Scopes = scope
}

/*
import (
	"github.com/crhym3/go-endpoints/endpoints"
	"code.google.com/p/goauth2/oauth"
	oauth2 "code.google.com/p/google-api-go-client/oauth2/v2"
	"time"
)

const PACKAGE string = ""

const WEB_CLIENT_ID string = "522156758812-09f5qv0e4gqjdjqfocerqcud5m5jutau.apps.googleusercontent.com"
const ANDROID_CLIENT_ID string = "522156758812-29jkcaiofrismobslc4ioop1dvfhhgoi.apps.googleusercontent.com"

const SERVER_CLIENT_ID string = ""
const SERVER_SECRET string = ""
const SERVER_REFRESH string = ""

const SCOPE1 string = "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/plus.me"
const SCOPE2 string = "https://www.googleapis.com/auth/devstorage.full_control"
const SCOPE3 string = "https://www.googleapis.com/auth/androidpublisher"

var CLIENTIDS = []string{WEB_CLIENT_ID, ANDROID_CLIENT_ID, endpoints.ApiExplorerClientId}
var AUDIENCES = []string{WEB_CLIENT_ID}
var SCOPES = []string{SCOPE1}

var CONFIG = &oauth.Config{
	ClientId: SERVER_CLIENT_ID,
	ClientSecret: SERVER_SECRET,
	Scope: SCOPE1+" "+SCOPE2+" "+SCOPE3,
	AuthURL: "https://accounts.google.com/o/oauth2/auth",
	TokenURL: "https://accounts.google.com/o/oauth2/token",
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

func (s *Service) Google3(r *http.Request, _ *Q, _ *Q) error {
	c := endpoints.NewContext(r)
	t := r.Header.Get("authorization")
	c.Infof("GOOGLE TOKEN ============\n%s\n============",t[7:])
	TRANS.Transport=urlfetch.Client(c).Transport
	service, err := oauth2.New(TRANS.Client())
	result, err := service.Userinfo.V2.Me.Get().Do()
	c.Infof("GOOGLE USER ============\n%v\n============", result)
	return err
}

func (s *Service) Google2(r *http.Request, _ *Q, _ *Q) error {
	c := endpoints.NewContext(r)
	g, err := endpoints.CurrentUser(c, SCOPES, AUDIENCES, CLIENTIDS);
	if err != nil {return err}
	c.Infof("GOOGLE USER ============\n%v\n============",g)
	return nil
}
*/
