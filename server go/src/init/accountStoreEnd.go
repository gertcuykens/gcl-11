package init

import (
	"net/http"
	"cloud"
	"appengine/urlfetch"
	"encoding/json"
	"io/ioutil"
	"github.com/crhym3/go-endpoints/endpoints"
)

type Data struct {
	Name string `json:"name"`
	Perms []string `json:"perms"`
}

type Accounts struct {
	Data []Data `json:"data"`
}

func (a *Accounts) set(c endpoints.Context, r *http.Request) error {
	//t := r.Header.Get("Authorization")
	//c.Infof("============%s",t[7:])

	client := urlfetch.Client(c)
	req, err := http.NewRequest("GET", "https://graph.facebook.com/me/accounts", nil) //&access_token=
	if err != nil {return nil}
	req.Header = map[string][]string{"Authorization": {r.Header.Get("Authorization")}}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	err = json.Unmarshal(b,a)
	if err != nil {return nil}
	return nil
}

func (a *Accounts) editor(c endpoints.Context, r *http.Request) bool {
	a.set(c,r)
	for i,x := range a.Data {
		if x.Name=="Gcl-11" {
			for _,y := range a.Data[i].Perms {
				if y == "CREATE_CONTENT" {
					return true
				}
			}
		}
	}
	return false
}

func (s *Service) Editor(r *http.Request,  _ *cloud.Entity, _ *cloud.Entity) error {
	c := endpoints.NewContext(r)
	//c, err := appengine.Namespace(c2, "")
	//if err != nil {return err}

	s.Status="no authentication"
	var a = &Accounts{}
	if !a.editor(c,r) {return s}
	s.Status="ok"
	return nil
}
