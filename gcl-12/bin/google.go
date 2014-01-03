package bin

import (
	"io/ioutil"
	"appengine/urlfetch"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

func (s *Service) GoogleUser(r *http.Request, req *Request, resp *Response) error {
	c := endpoints.NewContext(r)
	g, err := endpoints.CurrentUser(c, SCOPES, AUDIENCES, CLIENTIDS)
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
