package bin

import (
	"crypto/sha1"
	"encoding/hex"
	"time"
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/url"
	"mime"
	"io/ioutil"
	"encoding/json"
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

func (t *Token) Error() string {
	return t.Status
}

func (t *Token) Expired() bool {
	if t.Expiry.IsZero() {return false}
	return t.Expiry.Before(time.Now())
}

func (t *Token) CheckSum() error {
	if t.Extra == nil {t.Status="No group set!"; return t}
	h := sha1.New()
	a := t.Extra[0].Value+t.Extra[1].Value+t.Expiry.String()+SERVER_SECRET
	s := hex.EncodeToString(h.Sum([]byte(a)))
	if t.Expired() {t.Status="Token expired!"; return t}
	if t.Access != s {t.Status="Token checkSum error!"; return t}
	return nil
}

/*func (t *Token) SelectId() (err error) {
	transport.Token = t
	transport.TokenCache=&Cache{Context: t.Context, Key: t.Type}
	switch t.Type {
		case "facebook": FacebookUser(t)
	    case "google": GoogleUserToken(t)
	    case "twitter": TwitterUser(t)
	    case "linkedin": LinkedInUser(transport)
		default: t.Status = "Unrecognized Id Type!"; return t
	}
	return nil
}*/

func (t *Token) Renew() error {
	var v = url.Values{}
	v.Set("client_id", WEB_CLIENT_ID)
	v.Set("client_secret", CLIENT_SECRET)
	v.Set("refresh_token", REFRESH_TOKEN)
	v.Set("grant_type","refresh_token")

	r, err := t.Client.PostForm("https://accounts.google.com/o/oauth2/token", v)
	if err != nil {return err}
	defer r.Body.Close()
    if r.StatusCode != 200 {t.Status="Error Refresh Token "+r.Status; return t}

	content, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	switch content {
		case "application/x-www-form-urlencoded", "text/plain":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {return err}
			vals, err := url.ParseQuery(string(body))
			if err != nil {return err}
			t.Access = vals.Get("access_token")
			t.Type = vals.Get("token_type")
			t.ExpiresIn, _ = time.ParseDuration(vals.Get("expires_in") + "s")
		default:
			if err = json.NewDecoder(r.Body).Decode(&t); err != nil {return err}
			t.ExpiresIn *= time.Second
	}

	if t.ExpiresIn == 0 {t.Expiry = time.Time{}} else { t.Expiry = time.Now().Add(t.ExpiresIn) }
	return nil
}

/*
POST /o/oauth2/token HTTP/1.1
Host: accounts.google.com
Content-Type: application/x-www-form-urlencoded

client_id=8819981768.apps.googleusercontent.com&
client_secret={client_secret}&
refresh_token=1/6BMfW9j53gdGImsiyUH5kU5RsR4zwI9lUVX-tqf8JXQ&
grant_type=refresh_token

{
  "access_token":"1/fFBGRNJru1FQd44AzqT3Zg",
  "expires_in":3920,
  "token_type":"Bearer",
}
*/

//t.Refresh = vals.Get("refresh_token")
//t.Id = vals.Get("id_token")
//if len(t.Refresh) > 0 {t.Refresh = t.Refresh}
