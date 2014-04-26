package deprecated

import (
	//"github.com/crhym3/go-endpoints/endpoints"
	//"appengine/urlfetch"
	//"net/http"
	"io/ioutil"
    "encoding/xml"
)

var transport = &Transport{
	Token: &Token{},
	ClientId:     "77snhv7ncbvs5f",
	ClientSecret: "",
	RedirectURL:  "http://localhost:8080/_ah/api/rest/v0/linkedin/callback",
	AuthURL:      "https://www.linkedin.com/uas/oauth2/authorization",
	TokenURL:     "https://www.linkedin.com/uas/oauth2/accessToken",
	Scopes: []string{"r_basicprofile", "r_emailaddress"},
}
/*
func (s *Service) LinkedInOauth(r *http.Request, req *NoRequest, resp *ResponseOauth) error {
	resp.Url=transport.AuthCodeURL("LinkedIn")
	return nil
}

func (s *Service) LinkedInCallback(r *http.Request, req *RequestCallback, resp *Token) (err error) {
	c := endpoints.NewContext(r)
	transport.TokenCache=&Cache{Context: c, Key: "linkedin"}
	transport.Token.Context= c
	transport.Token.Client= urlfetch.Client(c)
	err = transport.Exchange(req.Code)
	err = LinkedInUser(transport)
	resp.Email=transport.Token.Email
	return err
}
*/
func LinkedInUser(transport *Transport) (err error) {
	r, err := transport.Token.Client.Get("https://api.linkedin.com/v1/people/~/email-address?oauth2_access_token="+transport.Token.Access)
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {return err}
	var u struct {
		XMLName xml.Name `xml:"email-address"`
		Email string `xml:",chardata"`
	}
	err = xml.Unmarshal(b, &u)
	if err != nil {return err}
	transport.Token.Email=u.Email
	return nil
}
