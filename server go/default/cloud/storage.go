package cloud

import (
	"log"
	"net/http"
	storage2 "code.google.com/p/google-api-go-client/storage/v1beta2"
)

type Storage struct {
	BucketName string
	ObjectName string
	Object *storage2.Object
	Service *storage2.Service
}

func (s *Storage) New(c *http.Client) (err error) {
	log.Print("-------NEW----------")
	s.Service, err = storage2.New(c)
	return err
}

func (s *Storage) Set(entity string) (err error) {
	log.Print("-------SET----------")
	acl := &storage2.ObjectAccessControl{
		Bucket: s.BucketName,
		Entity: "user-"+entity,
		Object: s.ObjectName,
		Role: "READER",
	}
	result, err := s.Service.ObjectAccessControls.Insert(s.BucketName, s.ObjectName, acl).Do()
	log.Printf("---------------Error ACL for %s/%s:\n%v", s.BucketName, s.ObjectName, err)
	if err != nil {return err}
	log.Printf("---------------Result ACL for %s/%s:\n%v", s.BucketName, s.ObjectName, result)
	return nil
}

func (s *Storage) Get() (err error) {
	s.Object, err = s.Service.Objects.Get(s.BucketName, s.ObjectName).Do()
	log.Printf("Got storage.Object, err: %#v", s.Object)
	return nil
}

/*
func (s *Storage) Insert(f io.Reader) (err error) {
	s.Object, err = s.Service.Objects.Insert(s.Bucket, s.Object).Media(f).Do()
	if err != nil {}
	log.Printf("Got storage.Object, err: %#v, %v", s.Object, err)
	return err
}
*/
