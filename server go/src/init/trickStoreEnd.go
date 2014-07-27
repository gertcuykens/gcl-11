package init

import (
	"net/http"
	"appengine/datastore"
	"cloud"
	"github.com/crhym3/go-endpoints/endpoints"
)

func (s *Service) GetTrickList(r *http.Request, m *cloud.TrickList, resp *cloud.TrickList) error {
	c := endpoints.NewContext(r)
	//c, err := appengine.Namespace(c2, "")
	//if err != nil {return err}

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.TrickStore {
		Root:  k,
		Entity: &cloud.TrickList{},
		Context: c,
	}
	d.GetTrickList()

	*resp = *d.Entity
	return nil
}

func (s *Service) PutTrickName(r *http.Request, m *cloud.TrickList, resp *cloud.TrickList) error {
	c := endpoints.NewContext(r)
	//c, err := appengine.Namespace(c2, "")
	//if err != nil {return err}

	s.Status="no authentication"
	var a = &Accounts{}
	if !a.editor(c,r) {return s}
	s.Status="ok"

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.TrickStore {
		Root:  k,
		Entity: m,
		Context: c,
	}
	d.PutTrickName()

	*resp = *d.Entity
	return nil
}

func (s *Service) DeleteTrickName(r *http.Request, m *cloud.TrickList, resp *cloud.TrickList) error {
	c := endpoints.NewContext(r)
	//c, err := appengine.Namespace(c2, "")
	//if err != nil {return err}

	s.Status="no authentication"
	var a = &Accounts{}
	if !a.editor(c,r) {return s}
	s.Status="ok"

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.TrickStore {
		Root:  k,
		Entity: m,
		Context: c,
	}
	d.DeleteTrickName()

	*resp = *d.Entity
	return nil
}
