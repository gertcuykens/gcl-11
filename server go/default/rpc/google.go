package rpc

import (
	"io/ioutil"
	"appengine/urlfetch"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
	"cloud"
	publisher "code.google.com/p/google-api-go-client/androidpublisher/v1.1"
	"strconv"
)

type Message struct {Message string `json:"message"`}

type Iab struct {
	Product string `json:"product"`
	Receipt string `json:"receipt"`
}

func (s *Service) GoogleCallback(r *http.Request, _, resp *Message) error {
	e := endpoints.NewContext(r)
	g, err := endpoints.CurrentUser(e, SCOPES, AUDIENCES, CLIENTIDS);
	if err != nil {return err}
	resp.Message = g.Email
	return nil
}

func (s *Service) GoogleRevoke(r *http.Request, _, resp *Message) (err error) {
	e := endpoints.NewContext(r)
	t := r.Header.Get("authorization")
	e.Infof("============="+t[7:])
	buf, err := urlfetch.Client(e).Get("https://accounts.google.com/o/oauth2/revoke?token="+t[7:])
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	resp.Message=string(b)
	return err
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

//req, _ := http.NewRequest("GET", urls, body)
//req.URL.Path =
//googleapi.SetOpaque(req.URL)
//req.Header.Set("User-Agent", "google-api-go-client/0.5")
//res, err := TRANS.Client().Do(req)

