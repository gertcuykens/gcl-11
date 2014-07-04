package cloud

import (
	"github.com/crhym3/go-endpoints/endpoints"
	"appengine/datastore"
	"time"
)

//Id int `json:"id" endpoints:"d=0,min=0,max=1" datastore:"noindex"`
type Message struct {
	Id int64 `json:"id,string" datastore:"-"` //omitempty
    Date time.Time `json:"date"`
	Judge string `json:"judge"`
	Event string `json:"event"`
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
	key := datastore.NewKey(s.Context, "tarifa", "", 0, s.Root)
	for _,m := range s.Entity.List {
		m.Judge=u
		m.Date=time.Now()
		key, err = datastore.Put(s.Context, key, m)
	}
	return nil
}

func (s *DataStore) Put2(u string) (err error) {
	key := datastore.NewKey(s.Context, "tarifa", "current", 0, s.Root)
	for _,m := range s.Entity.List {
		m.Judge=u
		m.Date=time.Now()
		key, err = datastore.Put(s.Context, key, m)
	}
	return nil
}

func (s *DataStore) Get() (err error) {
	for _,m := range s.Entity.List {
		key := datastore.NewKey(s.Context, "tarifa", "", m.Id, s.Root)
		err = datastore.Get(s.Context, key, m)
	}
	return nil
}

func (s *DataStore) Get2() (err error) {
	for _,m := range s.Entity.List {
		key := datastore.NewKey(s.Context, "tarifa", "current", 0, s.Root)
		err = datastore.Get(s.Context, key, m)
	}
	return nil
}

func (s *DataStore) GetAll() (err error) {
	q := datastore.NewQuery("tarifa").Ancestor(s.Root).Order("-Date")
	for t := q.Run(s.Context);; {
		var m Message
		k, err := t.Next(&m)
		if err == datastore.Done {err=nil; break}
		if err != nil {break}
		m.Id = k.IntID()
		s.Entity.List = append(s.Entity.List, &m)
	}
	return nil
}

func (s *DataStore) Delete() (err error) {
	for _, m := range s.Entity.List {
		key := datastore.NewKey(s.Context, "tarifa", "", m.Id, s.Root)
		datastore.Delete(s.Context,key)
	}
	return nil
}

func (s *DataStore) Truncate() (err error) {
	q := datastore.NewQuery("tarifa")
	var m []Message
	keys, err := q.GetAll(s.Context, &m)
	if err != nil {return err}
	for _, k := range keys {datastore.Delete(s.Context, k)}
	return nil
}

//s.Context.Infof("==========>%v",m)
