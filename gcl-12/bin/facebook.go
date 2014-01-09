package bin

import (
	"appengine/urlfetch"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"
)

func (s *Service) FacebookCallback(r *http.Request, t *Token, v *Token) error {
	t.Client = urlfetch.Client(endpoints.NewContext(r))
	if err := FacebookUser(t); err != nil {t.Status="Facebook User error!"; return err}
	*v = *t
	return nil
}

func  FacebookUser(req *Token) (err error) {
	resp, err := req.Client.Get("https://graph.facebook.com/me?access_token="+req.Access)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &req)
	return err
}
