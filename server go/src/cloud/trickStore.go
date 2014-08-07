package cloud

import (
	"net/http"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
	"errors"
)

type Trick struct {
	Name string `json:"name" datastore:"-"`
	Difficulty int `json:"difficulty"`
}

type TrickList struct {
	List []*Trick `json:"list"`
}

type TrickStore struct {
	Entity *TrickList
	Request *http.Request
}

func (s *TrickStore) PutTrickName() (err error) {
	var c = endpoints.NewContext(s.Request)
	var t = s.Request.Header.Get("Authorization")
	var a = &Accounts{
		Context: c,
		Authorization: t,
	}
	if !a.Editor() {return errors.New("No authentication.")}
	root := datastore.NewKey(c, "feed", "gcl11", 0, nil)

	/*
	var key=[]*datastore.Key{}
	for _,m := range s.Entity.List {
		if (m.Name == "") {m.Name = "Unknown"}
		key = append(key, datastore.NewKey(c, "Trick list", m.Name, 0, root))
	}
	key, err = datastore.PutMulti(c, key, s.Entity.List)
	*/

	for _,m := range s.Entity.List {
		if (m.Name == "") {m.Name = "Unknown"}
		key := datastore.NewKey(c, "Trick list", m.Name, 0, root)
		go datastore.Put(c, key, m)
	}

	return nil
}

func (s *TrickStore) GetTrickList() (err error) {
	var c = endpoints.NewContext(s.Request)
	root := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	q := datastore.NewQuery("Trick list").Ancestor(root).Order("__key__")
	for t := q.Run(c);; {
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
	var c = endpoints.NewContext(s.Request)
	var t = s.Request.Header.Get("Authorization")
	var a = &Accounts{
		Context: c,
		Authorization: t,
	}
	if !a.Editor() {return errors.New("No authentication.")}
	root := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	for _, m := range s.Entity.List {
		key := datastore.NewKey(c, "Trick list", m.Name, 0, root)
		datastore.Delete(c, key)
	}
	return nil
}

/*
func (s *Trick) Load(c <-chan datastore.Property) error {
		return datastore.LoadStruct(s, c)
}

func (s *Trick) Save(c chan<- datastore.Property) error {
		return datastore.SaveStruct(s, c)
}
*/
