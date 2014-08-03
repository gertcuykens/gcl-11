package cloud

import (
	"net/http"
	"appengine/urlfetch"
	"encoding/json"
	"io/ioutil"
	"github.com/crhym3/go-endpoints/endpoints"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email []string `json:"email"`
	Authorization string `json:"-"`
	Context endpoints.Context `json:"-"`
}

func (u *User) set() error {
  //t := r.Header.Get("Authorization")
	//c.Infof("============%s",t[7:])
	client := urlfetch.Client(u.Context)
	req, err := http.NewRequest("GET", "https://graph.facebook.com/me?fields=name", nil) //&access_token=
	req.Header = map[string][]string{"Authorization": {u.Authorization}}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	err = json.Unmarshal(b,u)
	if err != nil {}

	return nil
}
