package cloud

import (
	"log"
	"net/http"
	"code.google.com/p/google-api-go-client/storage/v1beta2"
	"io"
)

type Cloud struct {
	Bucket string
	Object *storage.Object
	Service *storage.Service
}

func (s *Cloud) New(c *http.Client) (err error) {
	s.Service, err = storage.New(c)
	return err
}

func (s *Cloud) Insert(f io.Reader) (err error) {
	s.Object, err = s.Service.Objects.Insert(s.Bucket, s.Object).Media(f).Do()
	if err != nil {}
	log.Printf("Got storage.Object, err: %#v, %v", s.Object, err)
	return err
}

func (s *Cloud) Get() (err error) {
	s.Object, err = s.Service.Objects.Get(s.Bucket, s.Object.Name).Do()
	log.Printf("Got storage.Object, err: %#v", s.Object)
	return nil
}
