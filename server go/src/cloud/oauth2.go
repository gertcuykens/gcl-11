package cloud

import (
	"log"
	"net/http"
	oauth2 "code.google.com/p/google-api-go-client/oauth2/v2"
)

type Oauth2 struct {
	Email string
	Service *oauth2.Service
}

func (s *Oauth2) New(c *http.Client) (err error) {
	log.Print("-------NEW----------")
	s.Service, err = oauth2.New(c)
	return err
}

func (s *Oauth2) Get() (err error) {
	result, err := s.Service.Userinfo.V2.Me.Get().Do()
	log.Printf("---------------Error oauth2:\n%v", err)
	if err != nil {return err}
	log.Printf("---------------Result oauth2:\n%v", result)
	s.Email=result.Email
	return nil
}

/*
func (s *Service) GoogleServerUser (r *http.Request, req *NoRequest, resp *Response) error {
	e := endpoints.NewContext(r)
	TRANS.Transport=urlfetch.Client(e).Transport

	c := &cloud.Oauth2{}
	c.New(TRANS.Client())
	c.Get()

	resp.Message=c.Email
	return nil
}
*/


