package bin

import (
	"crypto/sha1"
	"encoding/hex"
	"time"
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"
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
	a := t.Extra[0].Value+t.Expiry.String()+SERVER_SECRET
	s := hex.EncodeToString(h.Sum([]byte(a)))
	if t.Expired() {t.Status="Token expired!"; return t}
	if t.Access != s {t.Status="Token checkSum error!"; return t}
	return nil
}

func (t *Token) SelectId() (err error) {
	transport.Token = t
	transport.TokenCache=&Cache{Context: t.Context, Key: t.Type}
	switch t.Type {
		case "facebook": FacebookUser(t)
	    case "google": GoogleUser(t)
	    case "twitter": TwitterUser(t)
	    case "linkedin": LinkedInUser(transport)
		//case "server": Server(transport)
		default: t.Status = "Unrecognized Id Type!"; return t
	}
	return nil
}
