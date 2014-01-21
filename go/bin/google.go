package bin

import (
	"io/ioutil"
	"appengine/urlfetch"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
	"cloud"
	"log"
)

type NoRequest struct {}

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

var clientids = []string{WEB_CLIENT_ID, ANDROID_CLIENT_ID, endpoints.ApiExplorerClientId}
var audiences = []string{WEB_CLIENT_ID}
var gscope = []string{SCOPE1}

func (s *Service) GoogleUserService(r *http.Request, req *NoRequest, resp *Response) error {
	e := endpoints.NewContext(r)
    g, err := endpoints.CurrentUser(e, gscope, audiences, clientids);
	if err != nil {return err}
	resp.Message = g.Email
	return nil
}

func (s *Service) GoogleRevokeService(r *http.Request, req *NoRequest, resp *Response) (err error) {
	t := r.Header.Get("authorization")
	log.Print("---------"+t[7:])
	buf, err := urlfetch.Client(endpoints.NewContext(r)).Get("https://accounts.google.com/o/oauth2/revoke?token="+t[7:])
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	resp.Message=string(b)
	return err
}

func (s *Service) GooglePurchasesService(r *http.Request, req *NoRequest, resp *Response) error {
	e := endpoints.NewContext(r)
	TRANS.Transport=urlfetch.Client(e).Transport

	c := &cloud.Publisher{
		Package:"com.appspot",
		Product:"gas",
		Recete:R,
	}

	c.New(TRANS.Client())
	c.Get()
	resp.Message = "Done."
	return nil
}

func (s *Service) GoogleStorageService(r *http.Request, req *NoRequest, resp *Response) (err error) {
	e := endpoints.NewContext(r)
    TRANS.Transport=urlfetch.Client(e).Transport

	c := &cloud.Storage{
	    BucketName: "gcl-storage",
		ObjectName: "test.txt",
	}

	c.New(TRANS.Client())
	c.Set("gert.cuykens.2@gmail.com")
	resp.Message = "ACL is set."
	return nil
}
