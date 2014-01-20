package bin

import (
	"io/ioutil"
	"appengine/urlfetch"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
	"log"
	"cloud"
	"code.google.com/p/goauth2/oauth"
)

const WEB_CLIENT_ID string = "1093123836049-ilqfjb2s2tqal2fobuduj8b790hnnfju.apps.googleusercontent.com"
const ANDROID_CLIENT_ID string = "1093123836049-g84t42bajg9dbu7q13m6eotii0hpdpd8.apps.googleusercontent.com"

var clientids = []string{WEB_CLIENT_ID, ANDROID_CLIENT_ID, endpoints.ApiExplorerClientId}
var audiences = []string{WEB_CLIENT_ID}
var google_scopes = []string{"https://www.googleapis.com/auth/userinfo.email"}

type NoRequest struct {}

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

func (s *Service) GoogleUserService(r *http.Request, req *NoRequest, resp *Response) error {
	c := endpoints.NewContext(r)
	g, err := endpoints.CurrentUser(c, google_scopes, audiences, clientids);
	if err != nil {return err}
	resp.Message=g.String()
	return nil
}

func (s *Service) GoogleRevokeService(r *http.Request, req *NoRequest, resp *Response) (err error) {
	t := r.Header.Get("authorization")
	buf, err := urlfetch.Client(endpoints.NewContext(r)).Get("https://accounts.google.com/o/oauth2/revoke?token="+t)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	resp.Message=string(b)
	return err
}

func (s *Service) GooglePurchasesService(r *http.Request, req *NoRequest, resp *Response) error {
	e := endpoints.NewContext(r)

	t := &oauth.Transport{
		Token:     IAB_TOKEN,
		Config:    IAB_CONFIG,
		Transport: urlfetch.Client(e).Transport,
	}

	c := &cloud.Publisher{
		Package:"com.appspot",
		Product:"gas",
		Recete:R,
	}

	c.New(t.Client())
	c.Get()
	resp.Message = "Done."
	return nil
}

func (s *Service) GoogleStorageService(r *http.Request, req *NoRequest, resp *Response) (err error) {
	e := endpoints.NewContext(r)

	t := &oauth.Transport{
		Token:     STORAGE_TOKEN,
		Config:    STORAGE_CONFIG,
		Transport: urlfetch.Client(e).Transport,
	}

	g, err := endpoints.CurrentUser(e, google_scopes, audiences, clientids);
	if err != nil {return err}
	log.Print("------LOGIN---------"+g.String())

	c := &cloud.Storage{
	    BucketName: "gcl-storage",
		ObjectName: "test.txt",
	}

	c.New(t.Client())
	c.Set("gert.cuykens.2@gmail.com")
	resp.Message = "ACL is set."
	return nil
}
