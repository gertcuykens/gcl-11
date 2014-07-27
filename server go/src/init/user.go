package init

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
}

func (u *User) set(c endpoints.Context, r *http.Request) error {
  //t := r.Header.Get("Authorization")
	//c.Infof("============%s",t[7:])

	client := urlfetch.Client(c)
	req, err := http.NewRequest("GET", "https://graph.facebook.com/me?fields=name", nil) //&access_token=
	req.Header = map[string][]string{"Authorization": {r.Header.Get("Authorization")}}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	err = json.Unmarshal(b,u)
	if err != nil {}

	return nil
}
