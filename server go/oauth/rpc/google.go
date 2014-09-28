package rpc

import (
	"io/ioutil"
	"appengine/urlfetch"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
	"encoding/json"
)

type GoogleUser struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Authorization string `json:"-"`
	Context endpoints.Context `json:"-"`
}

func (u *GoogleUser) set() error {
	client := urlfetch.Client(u.Context)
	req, err := http.NewRequest("GET", "https://www.googleapis.com/userinfo/v2/me", nil)
	req.Header = map[string][]string{"Authorization": {u.Authorization}}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	u.Context.Infof("GOOGLE USER ============\n%s\n============",string(b))
	err = json.Unmarshal(b,u)
	return err
}

func (s *Service) Google(r *http.Request, _*Q, _*Q) error {
	var c = endpoints.NewContext(r)
	var t = r.Header.Get("Authorization")
	//c.Infof("GOOGLE TOKEN ============\n%s\n============",t[7:])
	var u = &GoogleUser{
		Authorization: t,
		Context: c,
	}
	u.set()
	c.Infof("GOOGLE USER ============\n%v\n============",u)
	return nil
}

/*
func (s *Service) Revoke(r *http.Request, _*Q, _*Q) (err error) {
	c := endpoints.NewContext(r)
	t := r.Header.Get("authorization")
	c.Infof("============\n%s\n============",t[7:])
	buf, err := urlfetch.Client(c).Get("https://accounts.google.com/o/oauth2/revoke?token="+t[7:])
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	c.Infof("============\n%s\n============",string(b))
	return err
}
*/

/*
	//publisher "code.google.com/p/google-api-go-client/androidpublisher/v1.1"
	//"strconv"


type Message struct {Message string `json:"message"`}

type Iab struct {
	Product string `json:"product"`
	Receipt string `json:"receipt"`
}

func (s *Service) GooglePurchases(r *http.Request, req *Iab, resp *publisher.InappPurchase) error {
	e := endpoints.NewContext(r)
	TRANS.Transport=urlfetch.Client(e).Transport

	p := &cloud.Publisher{
		Package:PACKAGE,
		Product:req.Product,
		Recete:req.Receipt,
	}

	p.New(TRANS.Client())
	p.Get()

	*resp = *p.Result
	return nil
}

func (s *Service) GoogleStorage(r *http.Request, req *Iab, resp *Message) (err error) {
	e := endpoints.NewContext(r)
	TRANS.Transport=urlfetch.Client(e).Transport

	g, err := endpoints.CurrentUser(e, SCOPES, AUDIENCES, CLIENTIDS);
	if err != nil {return err}

	p := &cloud.Publisher{
		Package:PACKAGE,
		Product:req.Product,
		Recete:req.Receipt,
	}
	p.New(TRANS.Client())
	p.Get()

	c := &cloud.Storage{
		BucketName: "gcl-storage",
		ObjectName: "test.txt",
	}
	c.New(TRANS.Client())
	c.Get()

	buf , err := TRANS.Client().Get(c.Object.MediaLink)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	resp.Message=string(b)+" => "+g.Email+" => "+strconv.FormatInt(p.Result.PurchaseTime,10)
	return nil
}
*/

//req, _ := http.NewRequest("GET", urls, body)
//req.URL.Path =
//googleapi.SetOpaque(req.URL)
//req.Header.Set("User-Agent", "google-api-go-client/0.5")
//res, err := TRANS.Client().Do(req)

/*
import (
	"log"
	"net/http"
	oauth2 "code.google.com/p/google-api-go-client/oauth2/v2"
)

type Oauth2 struct {
	Email string
	Service *oauth2.Service
}

func (s *Oauth2) New(c *http.Client) (err error) {
	log.Print("-------NEW----------")
	s.Service, err = oauth2.New(c)
	return err
}

func (s *Oauth2) Get() (err error) {
	result, err := s.Service.Userinfo.V2.Me.Get().Do()
	log.Printf("---------------Error oauth2:\n%v", err)
	if err != nil {return err}
	log.Printf("---------------Result oauth2:\n%v", result)
	s.Email=result.Email
	return nil
}
*/

/*
func (s *Service) GoogleServerUser (r *http.Request, req *NoRequest, resp *Response) error {
	e := endpoints.NewContext(r)
	TRANS.Transport=urlfetch.Client(e).Transport

	c := &cloud.Oauth2{}
	c.New(TRANS.Client())
	c.Get()

	resp.Message=c.Email
	return nil
}
*/
