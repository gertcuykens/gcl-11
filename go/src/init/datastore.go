package init

import (
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"
	"appengine/datastore"
	"cloud"
	"appengine/urlfetch"
)

func (s *Service) List(r *http.Request, _, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)
	t := r.Header.Get("Authorization")
	c.Infof("============%s",t[7:])

	client := urlfetch.Client(c)

/*
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {return err}

	req.Header = map[string][]string{
		"Accept-Language": {"en-US"},
		"User-Agent" : {"Mozilla/5.0"},
	}
*/

	//res, err := client.Do(req)

	buf, err := client.Get("https://graph.facebook.com/me/accounts&access_token="+t[7:])
	c.Infof("============%v", buf)
    //f := FacebookAccounts()

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)

	d := cloud.DataStore {
		Root:  k,
		Entity: &cloud.Entity{},
		Context: c,
	}
	d.Get()
	//c.Infof("<============%s",d.Entity.List[0].Message)
	*resp = *d.Entity
	return err
}

func (s *Service) Submit(r *http.Request, m *cloud.Entity, resp *cloud.Entity) error {
	c := endpoints.NewContext(r)
	a := r.Header.Get("Authorization")
	c.Infof("============%s",a)
	//f := FacebookAccounts()

	//c.Infof("==========>%s",m.List[0].Message)

	k := datastore.NewKey(c, "feed", "gcl11", 0, nil)
	d := cloud.DataStore {
		Root:  k,
		Entity: m,
		Context: c,
	}
	d.Put()
	resp = d.Entity
	return nil
}

/*
greets := []*Response{}
greets = append(greets, &Response{Message: "hello"})
greets = append(greets, &Response{Message: "goodbye"})
resp.Items = greets
return nil
*/

/*
client := urlfetch.Client(c)

req, err := http.NewRequest("GET", uri.String(), nil)
if err != nil {return nil, err}

req.Header = map[string][]string{
	"Accept-Language": {"en-US"},
	"User-Agent" : {"Mozilla/5.0"},
}

res, err := client.Do(req)
*/
