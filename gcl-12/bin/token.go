package bin

import (
	"encoding/json"
	"crypto/sha1"
	"encoding/hex"
	"time"
)

type Token struct {
	AccessToken string
	RefreshToken string
	Expiry time.Time
	Extra map[string]string
	Status string
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
	b, err := json.Marshal(t.Extra)
	if err != nil {t.Status="Token error! "+err.Error(); return t}
	h := sha1.New()
	a := string(b)+t.Expiry.String()+SERVER_SECRET
	s := hex.EncodeToString(h.Sum([]byte(a)))
	if t.Expired() {t.Status="Token expired!"; return t}
	if t.AccessToken != s {t.Status="Token checkSum error!"; return t}
	return nil
}
