package init

import (
	"net/http"
	"appengine/datastore"
	"cloud"
	"github.com/crhym3/go-endpoints/endpoints"
)

func (s *Service) GetHeat(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)
	//c, err := appengine.Namespace(c2, "")
	//if err != nil {return err}

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: &cloud.Entity{},
		Context: c,
	}
	d.GetHeat(m.List[0].Event, m.List[0].Division, m.List[0].Heat)

	*resp = *d.Entity
	return nil
}

func (s *Service) GetFirst(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)
	//c, err := appengine.Namespace(c2, "")
	//if err != nil {return err}

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: &cloud.Entity{},
		Context: c,
	}
	d.GetFirst(m.List[0].Event)

	*resp = *d.Entity
	return nil
}

func (s *Service) Get(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)
	//c, err := appengine.Namespace(c2, "")
	//if err != nil {return err}

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: &cloud.Entity{},
		Context: c,
	}
	d.Get(m.List[0].Event, m.List[0].Id)

	*resp = *d.Entity
	return nil
}

func (s *Service) Put(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)
	//c, err := appengine.Namespace(c2, "")
	//if err != nil {return err}

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

func (s *Service) Delete(r *http.Request, m *cloud.Entity, _ *cloud.Entity) error {
	c := endpoints.NewContext(r)
	//c, err := appengine.Namespace(c2, "")
	//if err != nil {return err}

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
	//c, err := appengine.Namespace(c2, "")
	//if err != nil {return err}

	s.Status="no authentication"
	var u = &User{}
	u.set(c,r)
	if(u.Name!="Gert Cuykens"){return s}
	s.Status="ok"

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: m,
		Context: c,
	}
	d.Truncate(m.List[0].Event)

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

/*
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
*/

/*
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
*/
