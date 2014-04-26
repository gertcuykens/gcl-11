package cloud

import (
	"log"
	"net/http"
	publisher "code.google.com/p/google-api-go-client/androidpublisher/v1.1"
)

type Iab publisher.InappPurchase

type Publisher struct {
	Package string
	Product string
	Recete string
	Result *publisher.InappPurchase
	Service *publisher.Service
}

func (s *Publisher) New(c *http.Client) (err error) {
	log.Print("-------NEW----------")
	s.Service, err = publisher.New(c)
	return err
}

func (s *Publisher) Get() (err error) {
	s.Result, err = s.Service.Inapppurchases.Get(s.Package,s.Product,s.Recete).Do()
	log.Printf("---------------Error for %s/%s:\n%v", s.Package, s.Product, s.Recete, err)
	if err != nil {return err}
	log.Printf("---------------Result for %s/%s:\n%v", s.Package, s.Product, s.Recete, s.Result)
	return nil
}
