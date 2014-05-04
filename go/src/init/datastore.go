package init

import (
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"
	"appengine/datastore"
	"cloud"
	"appengine/urlfetch"
	"encoding/json"
	"io/ioutil"
)

type Accounts struct {
	Data []Data `json:"data"`
}

type Data struct {
	Name string `json:"name"`
	Perms []string `json:"perms"`
}

func (s *Service) List(r *http.Request, _, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)
	t := r.Header.Get("Authorization")
	c.Infof("============%s",t[7:])

	client := urlfetch.Client(c)
	req, err := http.NewRequest("GET", "https://graph.facebook.com/me/accounts", nil) //&access_token=
	if err != nil {return err}
	req.Header = map[string][]string{"Authorization": {r.Header.Get("Authorization")}}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	var a = &Accounts{}
	err = json.Unmarshal(b,a)
	if err != nil {return err}
	c.Infof("============%v",a)
	s.Status="no authentication"
	err = s
	for i,x := range a.Data {
		if x.Name=="Gcl-11" {
			for _,y := range a.Data[i].Perms {
				if y == "CREATE_CONTENT" {
					err=nil
				}
			}
		}
	}
	if err != nil {return err}

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: &cloud.Entity{},
		Context: c,
	}
	d.Get()
	//c.Infof("<============%s",d.Entity.List[0].Message)
	*resp = *d.Entity
	return err
}

func (s *Service) Submit(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)
	t := r.Header.Get("Authorization")
	c.Infof("============%s",t)

	client := urlfetch.Client(c)
	req, err := http.NewRequest("GET", "https://graph.facebook.com/me/accounts", nil) //&access_token=
	if err != nil {return err}
	req.Header = map[string][]string{"Authorization": {r.Header.Get("Authorization")}}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	var a = &Accounts{}
	err = json.Unmarshal(b,a)
	if err != nil {return err}
	c.Infof("============%v",a)
	s.Status="no authentication"
	err = s
	for i,x := range a.Data {
		if x.Name=="Gcl-11" {
			for _,y := range a.Data[i].Perms {
				if y == "CREATE_CONTENT" {
					err=nil
				}
			}
		}
	}
	if err != nil {return err}

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: m,
		Context: c,
	}
	d.Put()
	resp = d.Entity
	return nil
}

/*
greets := []*Response{}
greets = append(greets, &Response{Message: "hello"})
greets = append(greets, &Response{Message: "goodbye"})
resp.Items = greets
return nil
*/
/*
req.Header = map[string][]string{
"Authorization": {t[7:]},
}
*/
/*
{
	"data": [
		{
			"category": "Community",
			"name": "Gcl-11",
			"access_token": "",
			"perms": [
				"ADMINISTER",
				"EDIT_PROFILE",
				"CREATE_CONTENT",
				"MODERATE_CONTENT",
				"CREATE_ADS",
				"BASIC_ADMIN"
			],
			"id": "269460066558605"
		}
	],
	"paging": {
		"next": "..."
	}
}
*/
