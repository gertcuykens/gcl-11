package init

import (
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"
	"appengine/datastore"
	"cloud"
)

func (s *Service) List(r *http.Request, _, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)
	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: &cloud.Entity{},
		Context: c,
	}
	d.Get()
	//c.Infof("<============%s",d.Entity.List[0].Message)
	*resp = *d.Entity
	return nil
}

func (s *Service) Submit(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)
	//c.Infof("==========>%s",m.List[0].Message)
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
