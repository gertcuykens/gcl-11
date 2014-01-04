package bin

import (
	"io/ioutil"
	"appengine/urlfetch"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

const WEB_CLIENT_ID string = "1034966141188-b4cup6jccsjqpdc14c9218fhb488e515.apps.googleusercontent.com"
const ANDROID_CLIENT_ID_d string = ""
const ANDROID_CLIENT_ID_r string = ""

var clientids = []string{WEB_CLIENT_ID, ANDROID_CLIENT_ID_d, ANDROID_CLIENT_ID_r, endpoints.ApiExplorerClientId}
var audiences = []string{WEB_CLIENT_ID}
var google_scopes = []string{"https://www.googleapis.com/auth/userinfo.email"}

func (s *Service) GoogleCallback(r *http.Request, req *Request, resp *Response) error {
	c := endpoints.NewContext(r)
	g, err := endpoints.CurrentUser(c, google_scopes, audiences, clientids)
	if err == nil {resp.Message="GoogleOauth2: "+g.String()}
	return err
}

func (s *Service) GoogleRevoke(r *http.Request, req *Request, resp *Response) error {
	c := endpoints.NewContext(r)
	b, err := Revoke(c, req.Access_token)
	if err == nil {resp.Message=b}
	return err
}

func Revoke(c endpoints.Context, access_token string) (string, error){
	httpClient := urlfetch.Client(c)
	resp, err := httpClient.Get("https://accounts.google.com/o/oauth2/revoke?token="+access_token)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	return string(b),err
}
