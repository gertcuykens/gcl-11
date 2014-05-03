package init

import (
	"appengine/urlfetch"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"
	"strconv"
	"time"
)

type Property struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type Token struct {
	Id int64 `json:"id_token"`
	Name string `json:"name_token"`
	Email string `json:"email_token"`
	Type string `json:"type_token"`
	Access string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	ExpiresIn time.Duration `json:"expires_in"`
	Expiry time.Time `json:"expiry"`
	Extra []Property `json:"extra"`
	Status string `json:"status"`
	Message string `json:"message"`
	Client *http.Client `json:"-"`
	Context endpoints.Context `json:"-"`
	Oauth_token string `json:"oauth_token"`
	Oauth_verifier string `json:"oauth_verifier"`
}

func (s *Service) FacebookCallback(r *http.Request, t *Token, v *Token) error {
	c := endpoints.NewContext(r)
	//c.Infof("------------------: %v", "-------------------")
	t.Client = urlfetch.Client(c)
	if err := FacebookUser(t); err != nil {t.Status="Facebook User error!"; return err}
	*v = *t
	return nil
}

func  FacebookUser(req *Token) (err error) {
	var f struct {
		Id string `json:"id"`
		Email string `json:"email"`
	}
	resp, err := req.Client.Get("https://graph.facebook.com/me?fields=email&access_token="+req.Access)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &f)
	req.Id, err = strconv.ParseInt(f.Id,10,64)
	req.Email = f.Email
	return err
}

type Accounts struct {
	Data []Data `json:"data"`
}

type Data struct {
	Name string `json:"name"`
	Perms []string `json:"perms"`
}

func  FacebookAccounts(req *Token) (err error) {
	resp, err := req.Client.Get("https://graph.facebook.com/me?fields=accounts?access_token="+req.Access)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &Accounts{})
	return err
}

/*
{
"data": [
{
"category": "Community",
"name": "Gcl-11",
"access_token": "",
"perms": [
"ADMINISTER",
"EDIT_PROFILE",
"CREATE_CONTENT",
"MODERATE_CONTENT",
"CREATE_ADS",
"BASIC_ADMIN"
],
"id": "269460066558605"
}
],
"paging": {
"next": "https://graph.facebook.com/v2.0/100001676421210/accounts?limit=5000&offset=5000&__after_id=enc_AeyDT65z8_Y-uBUnYu7-82gjWe3EGIDgVs5S-6y2bUUV4mRmjNZaUQ9M07o8iJfFJ_J4OAhtMIJVDiAfQfYkCce1"
}
}
*/
