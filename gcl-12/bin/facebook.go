package bin

import (
	"appengine/urlfetch"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"
)

type UserF struct {
	Name string
	Id string
}

func (s *Service) FacebookCallback(r *http.Request, req *Request, resp *Response) error {
	c := endpoints.NewContext(r)
	httpClient := urlfetch.Client(c)
	f, err := FacebookUser(httpClient, req.Access_token)
	if err == nil {resp.Message="Facebook user: "+f.Name}
	return err
}

func  FacebookUser(httpClient *http.Client, access_token string) (u *UserF, err error) {
	resp, err := httpClient.Get("https://graph.facebook.com/me?access_token="+access_token)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &u)
	return u, err
}
