package bin

import (
	"time"
    //"appengine/datastore"
)

type Request struct {
  Id int `json:"id" endpoints:"d=0,min=0,max=1"`
}

type Response struct {
  Message string `json:"message"`
}

type ResponseList struct {
  Items []*Response `json:"items"`
}

type Multiply struct {
  Times int `json:"times"`
  Message string `json:"message"`
}

type Greeting struct {
  //Key     *datastore.Key `json:"id" datastore:"-"`
  Author  string         `json:"author"`
  Content string         `json:"content" datastore:", noindex" endpoints:"req"`
  Date    time.Time      `json:"date"`
}

type NoRequest struct {
}

type GreetingService struct {
}


