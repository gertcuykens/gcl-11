package cloud

import (
	"log"
	"net/http"
	storage "code.google.com/p/google-api-go-client/storage/v1beta2"
)

type Storage struct {
	BucketName string
	ObjectName string
	ObjectAcl *storage.ObjectAccessControl
	Service *storage.Service
}

func (s *Storage) New(c *http.Client) (err error) {
	s.Service, err = storage.New(c)
	return err
}

func (s *Storage) Set(entity string) (err error) {
	s.ObjectAcl = &storage.ObjectAccessControl{
		Bucket: s.BucketName,
		Entity: entity,
		Object: s.ObjectName,
		Role: "READER",
	}
	result, err := s.Service.ObjectAccessControls.Insert(s.BucketName, s.ObjectName, s.ObjectAcl).Do()
	log.Printf("---------------Result of inserting ACL for %s/%s:\n%v", s.BucketName, s.ObjectName, result)
	return nil
}

/*
func (s *Storage) Get() (err error) {
	s.Object, err = s.Service.Objects.Get(s.Bucket, s.Object.Name).Do()
	log.Printf("Got storage.Object, err: %#v", s.Object)
	return nil
}

func (s *Storage) Insert(f io.Reader) (err error) {
	s.Object, err = s.Service.Objects.Insert(s.Bucket, s.Object).Media(f).Do()
	if err != nil {}
	log.Printf("Got storage.Object, err: %#v, %v", s.Object, err)
	return err
}
*/
