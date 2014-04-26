package deprecated

import (
	"time"
	"appengine"
	"appengine/memcache"
)

type Cache struct {
	Context appengine.Context
	Key     string
}

func (m Cache) Token() (*Token, error) {
	item, err := memcache.Get(m.Context, m.Key)
	if err != nil {return nil, err}
	return &Token{Access: string(item.Value), Expiry: time.Now().Add(item.Expiration),}, nil
}

func (m Cache) PutToken(tok *Token) error {
	return memcache.Set(m.Context, &memcache.Item{
		Key: m.Key,
		Value: []byte(tok.Access),
		Expiration: tok.Expiry.Sub(time.Now()),
	})
}

/*
type Cache interface {
	Token() (*Token, error)
	PutToken(*Token) error
}
*/
