package cloud

import (
	"appengine/datastore"
	"time"
	"github.com/crhym3/go-endpoints/endpoints"
)

//Id int `json:"id" endpoints:"d=0,min=0,max=1" datastore:"noindex"`
type Message struct {
	Id int64 `json:"id,string" datastore:"-"` //omitempty
  Date time.Time `json:"date"`
	Judge string `json:"judge"`
	Event string `json:"event" datastore:"-"`
	Division string `json:"division"`
	Heat int `json:"heat"`
	Rider string `json:"rider"`
	Trick string `json:"trick"`
	Score int `json:"score"`
	Attempt int `json:"attempt"`
}

type Entity struct {
	List []*Message `json:"list"`
}

type DataStore struct {
	Root *datastore.Key
	Entity *Entity
	Context endpoints.Context
}

func (s *DataStore) Put(u string) (err error) {
	for _,m := range s.Entity.List {
		key := datastore.NewKey(s.Context, m.Event, "", 0, s.Root)
		if (u!="Gert Cuykens") {m.Judge=u}
		m.Date=time.Now()
		key, err = datastore.Put(s.Context, key, m)
	}
	return nil
}

func (s *DataStore) Get(e string, id int64) (err error) {
	var m Message
	m.Event=e
	key := datastore.NewKey(s.Context, e, "", id, s.Root)
	err = datastore.Get(s.Context, key, &m)
	s.Entity.List = append(s.Entity.List, &m)
	return nil
}

func (s *DataStore) GetHeat(e string, d string, h int) (err error) {
	q := datastore.NewQuery(e).Ancestor(s.Root).Filter("Division =", d).Filter("Heat =", h).Order("-Date")
	for t := q.Run(s.Context);; {
		var m Message
		m.Event=e
		k, err := t.Next(&m)
		if err == datastore.Done {err=nil; break}
		if err != nil {break}
		m.Id = k.IntID()
		s.Entity.List = append(s.Entity.List, &m)
	}
	return nil
}

func (s *DataStore) GetFirst(e string) (err error) {
	//s.Context.Infof("============%s",e)
	q := datastore.NewQuery(e).Ancestor(s.Root).Order("-Date")
	t := q.Run(s.Context)
	var m Message
	m.Event=e
	k, err := t.Next(&m)
	if err != nil {return err}
	m.Id = k.IntID()
	s.Entity.List = append(s.Entity.List, &m)
	return nil
}

func (s *DataStore) Delete() (err error) {
	for _, m := range s.Entity.List {
		key := datastore.NewKey(s.Context, m.Event, "", m.Id, s.Root)
		datastore.Delete(s.Context, key)
	}
	return nil
}

func (s *DataStore) Truncate(e string) (err error) {
	q := datastore.NewQuery(e)
	var m []Message
	keys, err := q.GetAll(s.Context, &m)
	if err != nil {return err}
	for _, k := range keys {datastore.Delete(s.Context, k)}
	return nil
}

//s.Context.Infof("==========>%v",m)

/*
func (s *DataStore) Get2() (err error) {
	var m Message
	key := datastore.NewKey(s.Context, "tarifa", "current", 0, s.Root)
	err = datastore.Get(s.Context, key, &m)
	s.Entity.List = append(s.Entity.List, &m)
	return nil
}
*/

/*
func (s *DataStore) Put2(u string) (err error) {
	key := datastore.NewKey(s.Context, "tarifa", "current", 0, s.Root)
	for _,m := range s.Entity.List {
		m.Judge=u
		m.Date=time.Now()
		key, err = datastore.Put(s.Context, key, m)
	}
	return nil
}
*/
