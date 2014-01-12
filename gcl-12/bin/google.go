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

func (s *Service) GoogleCallback(r *http.Request, req *NoRequest, resp *Token) error {
	c := endpoints.NewContext(r)
	g, err := endpoints.CurrentUser(c, google_scopes, audiences, clientids);
	if err != nil {return err}
	resp.Email=g.String()
	return nil
}

func (s *Service) GoogleRevoke(r *http.Request, req *Token, resp *Token) (err error) {
	req.Client = urlfetch.Client(endpoints.NewContext(r))
	if err = Revoke(req); err != nil {return err}
	resp = req
	return nil
}

func Revoke(t *Token) (err error){
	resp, err := t.Client.Get("https://accounts.google.com/o/oauth2/revoke?token="+t.Access)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	t.Status=string(b)
	return err
}

func GoogleUser(t *Token) (error){
	g, err := endpoints.CurrentUser(t.Context, google_scopes, audiences, clientids);
	if err != nil {return err}
	t.Email = g.Email
	return err
}
