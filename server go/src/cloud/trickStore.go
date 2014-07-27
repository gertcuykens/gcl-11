package cloud

import (
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
)

type Trick struct {
	Name string `json:"name"`
	Difficulty int `json:"difficulty"`
}

type TrickList struct {
	List []*Trick `json:"list"`
}

type TrickStore struct {
	Root *datastore.Key
	Entity *TrickList
	Context endpoints.Context
}

func (s *TrickStore) PutTrickName() (err error) {
	for _,m := range s.Entity.List {
		key := datastore.NewKey(s.Context, "Trick list", m.Name, 0, s.Root)
		key, err = datastore.Put(s.Context, key, m)
	}
	return nil
}

func (s *TrickStore) GetTrickList() (err error) {
	q := datastore.NewQuery("Trick list").Ancestor(s.Root).Order("Name")
	for t := q.Run(s.Context);; {
		var m Trick
		k, err := t.Next(&m)
		if err == datastore.Done {err=nil; break}
		if err != nil {break}
		m.Name = k.StringID()
		s.Entity.List = append(s.Entity.List, &m)
	}
	return nil
}

func (s *TrickStore) DeleteTrickName() (err error) {
	for _, m := range s.Entity.List {
		key := datastore.NewKey(s.Context, "Trick list", m.Name, 0, s.Root)
		datastore.Delete(s.Context, key)
	}
	return nil
}
