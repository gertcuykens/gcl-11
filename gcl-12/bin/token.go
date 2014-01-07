package bin

import (
	"crypto/sha1"
	"encoding/hex"
	"time"
)

type Property struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type Token struct {
	Access string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	ExpiresIn time.Duration `json:"expires_in"`
	Expiry time.Time `json:"expiry"`
	Id string `json:"id_token"`
	Extra []Property `json:"extra"`
	Status string `json:"status"`
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
