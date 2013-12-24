package bin

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

func (gs *GreetingService) Multiply(r *http.Request, req *Multiply, resp *Response) error {
    resp.Message = strings.Repeat(req.Message, req.Times)
    return nil
}

func (gs *GreetingService) Id(r *http.Request, req *Request, resp *Response) error {
    greets := []*Response{}
    greets = append(greets, &Response{Message: "hello"})
    greets = append(greets, &Response{Message: "goodbye"})
    resp.Message = greets[req.Id].Message
    return nil
}

func (gs *GreetingService) List(r *http.Request, req *NoRequest, resp *ResponseList) error {
    greets := []*Response{}
    greets = append(greets, &Response{Message: "hello"})
    greets = append(greets, &Response{Message: "goodbye"})
    resp.Items = greets
    return nil
}

func (gs *GreetingService) Login(r *http.Request, req *NoRequest, resp *Response) error {
    c := endpoints.NewContext(r)
    u, err := endpoints.CurrentUser(c, SCOPES, AUDIENCES, CLIENTIDS)
    if err != nil {return err}
    if u != nil {resp.Message="hello "+u.String()}
    return nil
}

func (gs *GreetingService) Soap(r *http.Request, req *NoRequest, resp *Response) error {
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

func (gs *GreetingService) Store(r *http.Request, req *NoRequest, resp *Response) error {
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

func init() {
  greetService := &GreetingService{}
  api, err := endpoints.RegisterService(greetService, "rest1", "0", "Greetings API", true)
  if err != nil {panic(err.Error())}

  info1 := api.MethodByName("List").Info()
  info1.Name, info1.HttpMethod, info1.Path, info1.Desc = "greetings.listGreeting", "GET", "response", "List of greetings."

  info2 := api.MethodByName("Store").Info()
  info2.Name, info2.HttpMethod, info2.Path, info2.Desc, info2.Scopes = "greetings.datastore", "GET", "greetings/datastore", "Datastore.", SCOPES

  info3 := api.MethodByName("Soap").Info()
  info3.Name, info3.HttpMethod, info3.Path, info3.Desc = "greetings.soap", "GET", "greetings/soap", "Soap."

  info4 := api.MethodByName("Login").Info()
  info4.Name, info4.HttpMethod, info4.Path, info4.Desc, info4.Scopes= "greetings.authed", "POST", "greetings/authed", "Login.", SCOPES

  info5 := api.MethodByName("Id").Info()
  info5.Name, info5.HttpMethod, info5.Path, info5.Desc= "greetings.getGreeting", "GET", "response/{id}", "Get greeting."

  info6 := api.MethodByName("Multiply").Info()
  info6.Name, info6.HttpMethod, info6.Path, info6.Desc= "greetings.multiply", "POST", "response", "Multiply greeting."

  endpoints.HandleHttp()
  http.HandleFunc("/", welcome)
}

/*
  if req.Limit <= 0 {req.Limit = 10}
  c := endpoints.NewContext(r)
  q := datastore.NewQuery("Greeting").Order("-Date").Limit(req.Limit)
  greets := make([]*Greeting, 0, req.Limit)
  keys, err := q.GetAll(c, &greets)
  if err != nil {return err}
  for i, k := range keys {greets[i].Key = k}
  resp.Items = greets
*/
