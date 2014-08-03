package cloud

import (
	"net/http"
	"appengine/urlfetch"
	"encoding/json"
	"io/ioutil"
	"github.com/crhym3/go-endpoints/endpoints"
)

type Data struct {
	Name string `json:"name"`
	Perms []string `json:"perms"`
}

type Accounts struct {
	Data []Data `json:"data"`
	Authorization string `json:"-"`
	Context endpoints.Context `json:"-"`
}

func (a *Accounts) set() error {
	//t := r.Header.Get("Authorization")
	//c.Infof("============%s",t[7:])
	client := urlfetch.Client(a.Context)
	req, err := http.NewRequest("GET", "https://graph.facebook.com/me/accounts", nil) //&access_token=
	if err != nil {return nil}
	req.Header = map[string][]string{"Authorization": {a.Authorization}}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	err = json.Unmarshal(b,a)
	if err != nil {return nil}
	return nil
}

func (a *Accounts) Editor() bool {
	a.set()
	for i,x := range a.Data {
		if x.Name=="Gcl-11" {
			for _,y := range a.Data[i].Perms {
				if y == "CREATE_CONTENT" {
					return true
				}
			}
		}
	}
	return false
}

/*
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type ContextHandler func(appengine.Context, http.ResponseWriter, *http.Request)

func (f ContextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(appengine.NewContext(r), w, r)
	c.Save()
}

func init() {
	http.Handle("/test", ContextHandler(testHandler))
}

//Then to test testHandler:

func TestHandler(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatalf("Could not get a context - %v", err)
	}
	defer c.Close()
	r, _ := http.NewRequest("GET", "/test", nil)

	w := httptest.NewRecorder()
	testHandler(c, r, w)
	// test things that testHandler is supposed to do
}
*/
