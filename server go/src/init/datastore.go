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

type Data struct {
	Name string `json:"name"`
	Perms []string `json:"perms"`
}

type Accounts struct {
	Data []Data `json:"data"`
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

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email []string `json:"email"`
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

	return nil
}

func (s *Service) Editor(r *http.Request,  _ *cloud.Entity, _ *cloud.Entity) error {
	c := endpoints.NewContext(r)
	s.Status="no authentication"
	var a = &Accounts{}
	if !a.editor(c,r) {return s}
	s.Status="ok"
	return nil
}

func (s *Service) GetAll(r *http.Request, _ *cloud.Entity, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)

	//s.Status="No authentication."
	//var a = &Accounts{}
	//if !a.editor(c,r) {return s}
	//s.Status="ok"

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: &cloud.Entity{},
		Context: c,
	}
	d.GetAll()

	*resp = *d.Entity
	return nil
}

func (s *Service) Get(r *http.Request, _ *cloud.Entity, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)

	//s.Status="no authentication"
	//var a = &Accounts{}
	//if !a.editor(c,r) {return s}
	//s.Status="ok"

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

func (s *Service) Get2(r *http.Request, _ *cloud.Entity, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)

	//s.Status="no authentication"
	//var a = &Accounts{}
	//if !a.editor(c,r) {return s}
	//s.Status="ok"

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: &cloud.Entity{},
		Context: c,
	}
	d.Get2()

	*resp = *d.Entity
	return nil
}

func (s *Service) Put(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
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

func (s *Service) Put2(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
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
	d.Put2(u.Name)

	*resp = *d.Entity
	return nil
}

func (s *Service) Delete(r *http.Request, m *cloud.Entity, _ *cloud.Entity) error {
	c := endpoints.NewContext(r)

	s.Status="no authentication"
	var a = &Accounts{}
	if !a.editor(c,r) {return s}
	s.Status="ok"

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: m,
		Context: c,
	}
	d.Delete()

	return nil
}

func (s *Service) Truncate(r *http.Request, m *cloud.Entity, _ *cloud.Entity) error {
	c := endpoints.NewContext(r)

	s.Status="no authentication"
	var a = &Accounts{}
	if !a.editor(c,r) {return s}
	s.Status="ok"

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: m,
		Context: c,
	}
	d.Truncate()

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
