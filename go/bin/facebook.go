package bin

import (
	"appengine/urlfetch"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"
	"strconv"
)

func (s *Service) FacebookCallback(r *http.Request, t *Token, v *Token) error {
	t.Client = urlfetch.Client(endpoints.NewContext(r))
	if err := FacebookUser(t); err != nil {t.Status="Facebook User error!"; return err}
	*v = *t
	return nil
}

func  FacebookUser(req *Token) (err error) {
	var f struct {
		Id string `json:"id"`
		Name string `json:"name"`
	}
	resp, err := req.Client.Get("https://graph.facebook.com/me?access_token="+req.Access)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &f)
	req.Id, err = strconv.ParseInt(f.Id,10,64)
	req.Name = f.Name
	return err
}
