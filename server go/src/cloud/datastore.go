package cloud

import (
	"github.com/crhym3/go-endpoints/endpoints"
	"appengine/datastore"
	"time"
)

//Id int `json:"id" endpoints:"d=0,min=0,max=1" datastore:"noindex"`
type Message struct {
	User string `json:"user"`
	Message string `json:"message"`
    Date time.Time `json:"date"`
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
	key := datastore.NewKey(s.Context, "message", "", 0, s.Root)
	for _,m := range s.Entity.List {
		//s.Context.Infof("==========>%s",m)
		m.User=u
		m.Date=time.Now()
		key, err = datastore.Put(s.Context, key, m)
	}
	return
}

func (s *DataStore) Get() (err error) {
	q := datastore.NewQuery("message").Order("-Date")  //Ancestor(s.Root)
	for t := q.Run(s.Context);; {
		var m Message
		_, err := t.Next(&m)
		if err == datastore.Done {err=nil; break}
		if err != nil {break}
		//s.Context.Infof("<==========%s",m)
		s.Entity.List = append(s.Entity.List, &m)
	}
	return
}

func (s *DataStore) Delete() (err error) {
	s.Context.Infof("==========delete")
	q := datastore.NewQuery("message")
	var m []Message
	keys,err := q.GetAll(s.Context, &m)
	if err != nil {return err}
	for i, k := range keys {
		s.Context.Infof("<==========%d",i)
		datastore.Delete(s.Context,k)
	}
	return
}
