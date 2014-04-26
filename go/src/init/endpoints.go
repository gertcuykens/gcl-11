package init

import (
	"time"
	"strings"
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"
	"appengine/datastore"
	"appengine/urlfetch"
	//"appengine/user"
	//"appengine"
	"bytes"
	"encoding/xml"
	"io/ioutil"
)

type NoRequest struct {
}

type Request struct {
  Id int `json:"id" endpoints:"d=0,min=0,max=1"`
}

type Response struct {
  Message string `json:"message"`
}

type ResponseList struct {
  Items []*Response `json:"items"`
}

type Multiply struct {
  Times int `json:"times"`
  Message string `json:"message"`
}

type Greeting struct {
  //Key     *datastore.Key `json:"id" datastore:"-"`
  Author  string         `json:"author"`
  Content string         `json:"content" datastore:", noindex" endpoints:"req"`
  Date    time.Time      `json:"date"`
}

func (gs *Service) Multiply(r *http.Request, req *Multiply, resp *Response) error {
	resp.Message = strings.Repeat(req.Message, req.Times)
	return nil
}

func (gs *Service) Id(r *http.Request, req *Request, resp *Response) error {
	greets := []*Response{}
	greets = append(greets, &Response{Message: "hello"})
	greets = append(greets, &Response{Message: "goodbye"})
	resp.Message = greets[req.Id].Message
	return nil
}

func (gs *Service) List(r *http.Request, req *NoRequest, resp *ResponseList) error {
	greets := []*Response{}
	greets = append(greets, &Response{Message: "hello"})
	greets = append(greets, &Response{Message: "goodbye"})
	resp.Items = greets
	return nil
}

func (gs *Service) Login(r *http.Request, req *NoRequest, resp *Response) error {
	c := endpoints.NewContext(r)
	u, err := endpoints.CurrentUser(c, SCOPES, AUDIENCES, CLIENTIDS)
	if err != nil {return err}
	if u != nil {resp.Message="hello "+u.String()}
	return nil
}

func (gs *Service) Soap(r *http.Request, req *NoRequest, resp *Response) error {
	c := endpoints.NewContext(r)
	httpClient := urlfetch.Client(c)
	respx, err := httpClient.Post(SERVER, "text/xml; charset=utf-8", bytes.NewBufferString(QUERY))
	if err != nil {return err}
	b, err := ioutil.ReadAll(respx.Body)
	if err != nil {return err}
	in := string(b)
	var envelope SoapEnvelope
	parser := xml.NewDecoder(bytes.NewBufferString(in))
	err = parser.DecodeElement(&envelope, nil)
	if err != nil {return err}
	resp.Message = envelope.Body.VerifyEmailResponse.VerifyEmailResult.ResponseText
	respx.Body.Close()
	return nil
}

func (gs *Service) Store(r *http.Request, req *NoRequest, resp *Response) error {
	c := endpoints.NewContext(r)
	u, err := endpoints.CurrentUser(c, SCOPES, AUDIENCES, CLIENTIDS)

	c.Infof("-------------------")
	a:=""
	if u != nil {
		a=u.String()
		c.Infof("Hello, %v\n", u)
	}
	c.Infof("-------------------")

	g := Greeting {
		Author:  a,
		Content: "test",
		Date:    time.Now(),
	}

	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "employee", nil), &g)
	if err != nil {return err}

	err = datastore.Get(c, key, &g);
	if err != nil {return err}

	resp.Message="hello stored"
	return nil
}



