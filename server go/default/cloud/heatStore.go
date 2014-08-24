package cloud

import (
	"net/http"
	"appengine/datastore"
	"time"
	"github.com/crhym3/go-endpoints/endpoints"
	"errors"
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
	Entity *Entity
	Request *http.Request
}

func (s *DataStore) Put() (err error) {
	var c = endpoints.NewContext(s.Request)
	var t = s.Request.Header.Get("Authorization")
	var a = &Accounts{
		Authorization: t,
		Context: c,
	}
	if !a.Editor() {return errors.New("No authentication.")}
	var u = &User{
		Authorization: t,
		Context: c,
	}
	u.set()
	root := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	for _,m := range s.Entity.List {
    if (u.Name!="Gert Cuykens") {m.Judge=u.Name}
		m.Date=time.Now()
		key := datastore.NewKey(c, m.Event, "", 0, root)
		go datastore.Put(c, key, m)
	}
	return nil
}

func (s *DataStore) Get(e string, id int64) (err error) {
  var c = endpoints.NewContext(s.Request)
  root := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	var m Message
	m.Event=e
	key := datastore.NewKey(c, e, "", id, root)
	err = datastore.Get(c, key, &m)
	s.Entity.List = append(s.Entity.List, &m)
	return nil
}

func (s *DataStore) GetHeat(e string, d string, h int) (err error) {
	var c = endpoints.NewContext(s.Request)
	root := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	q := datastore.NewQuery(e).Ancestor(root).Filter("Division =", d).Filter("Heat =", h).Order("-Date")
	for t := q.Run(c);; {
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
	var c = endpoints.NewContext(s.Request)
	root := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	q := datastore.NewQuery(e).Ancestor(root).Order("-Date")
	t := q.Run(c)
	var m Message
	m.Event=e
	k, err := t.Next(&m)
	if err != nil {return err}
	m.Id = k.IntID()
	s.Entity.List = append(s.Entity.List, &m)
	return nil
}

func (s *DataStore) Delete() (err error) {
	var c = endpoints.NewContext(s.Request)
	var t = s.Request.Header.Get("Authorization")
	var a = &Accounts{
		Authorization: t,
		Context: c,
	}
	if !a.Editor() {return errors.New("No authentication.")}
	root := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	for _, m := range s.Entity.List {
		key := datastore.NewKey(c, m.Event, "", m.Id, root)
		go datastore.Delete(c, key)
	}
	return nil
}

func (s *DataStore) Truncate(e string) (err error) {
	var c = endpoints.NewContext(s.Request)
	var t = s.Request.Header.Get("Authorization")
	var u = &User{
		Authorization: t,
		Context: c,
	}
	u.set()
	if (u.Name!="Gert Cuykens") {return errors.New("No authentication.")}
	q := datastore.NewQuery(e)
	var m []Message
	keys, err := q.GetAll(c, &m)
	if err != nil {return err}
	for _, k := range keys {go datastore.Delete(c, k)}
	return nil
}

func (s *DataStore) Editor() (err error) {
	var c = endpoints.NewContext(s.Request)
	var t = s.Request.Header.Get("Authorization")
  var a = &Accounts{
	  Authorization: t,
	  Context: c,
  }
  if !a.Editor() {return errors.New("No authentication.")}
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
