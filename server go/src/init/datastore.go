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

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email []string `json:"email"`
}

func (a *Accounts) editor(c endpoints.Context, r *http.Request) bool {
	//t := r.Header.Get("Authorization")
	//c.Infof("============%s",t[7:])

	client := urlfetch.Client(c)
	req, err := http.NewRequest("GET", "https://graph.facebook.com/me/accounts", nil) //&access_token=
	if err != nil {return false}
	req.Header = map[string][]string{"Authorization": {r.Header.Get("Authorization")}}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	err = json.Unmarshal(b,a)
	if err != nil {return false}
	c.Infof("============%v",a)
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

func (u *User) set(c endpoints.Context, r *http.Request) error {
    //t := r.Header.Get("Authorization")
	//c.Infof("============%s",t[7:])

	client := urlfetch.Client(c)
	req, err := http.NewRequest("GET", "https://graph.facebook.com/me?fields=name", nil) //&access_token=
	req.Header = map[string][]string{"Authorization": {r.Header.Get("Authorization")}}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	err = json.Unmarshal(b,u)
	if err != nil {}
	c.Infof("============%v",u)

	return nil
}

func (s *Service) List(r *http.Request, _, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)

	s.Status="no authentication"
	var a = &Accounts{}
	if !a.editor(c,r) {return s}
	s.Status="ok"

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: &cloud.Entity{},
		Context: c,
	}
	d.Get()

	*resp = *d.Entity
	return nil
}

func (s *Service) Submit(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)

	s.Status="no authentication"
	var a = &Accounts{}
	if !a.editor(c,r) {return s}
	s.Status="ok"

	var u = &User{}
	u.set(c,r)

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: m,
		Context: c,
	}
	d.Put(u.Name)

	*resp = *d.Entity
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

//req.Id, err = strconv.ParseInt(f.Id, 10, 64)
