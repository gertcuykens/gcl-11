package rpc

import (
	"appengine/urlfetch"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"
)

/*****************************************************************************/

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
	client := urlfetch.Client(a.Context)
	req, err := http.NewRequest("GET", "https://graph.facebook.com/me/accounts", nil) //&access_token=
	if err != nil {return nil}
	req.Header = map[string][]string{"Authorization": {a.Authorization}}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	err = json.Unmarshal(b,a)
	return err
}

func (a *Accounts) Editor() bool {
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

/*****************************************************************************/

type User struct {
	Id int64 `json:"id,string"`
	Name string `json:"name"`
	Email string `json:"email"`
	Authorization string `json:"-"`
	Context endpoints.Context `json:"-"`
}

func (u *User) set() error {
	client := urlfetch.Client(u.Context)
	req, err := http.NewRequest("GET", "https://graph.facebook.com/me?fields=id,name,email", nil) //&access_token=
	req.Header = map[string][]string{"Authorization": {u.Authorization}}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	err = json.Unmarshal(b,u)
	return err
}

/*****************************************************************************/

func (s *Service) Facebook(r *http.Request, _*Q, _*Q) error {
	var c = endpoints.NewContext(r)
	var t = r.Header.Get("Authorization")
	//c.Infof("FACEBOOK TOKEN ============\n%s\n============",t[7:])
	var a = &Accounts{
		Authorization: t,
		Context: c,
	}
	a.set()
	var u = &User{
		Authorization: t,
		Context: c,
	}
	u.set()
	c.Infof("FACEBOOK USER ============\n%v\n============",u)
	return nil
}
