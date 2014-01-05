package bin

import (
	"github.com/crhym3/go-endpoints/endpoints"
	//"appengine"
	"appengine/urlfetch"
	"net/http"
	"io/ioutil"
	//"encoding/json"
	"log"
    "encoding/xml"
)

type UserL struct {
	XMLName xml.Name `xml:"email-address"`
	Email string `xml:",chardata"`
}

var transport = &Transport{
	ClientId:     "77snhv7ncbvs5f",
	ClientSecret: LINKEDIN_SECRET,
	RedirectURL:  "http://localhost:8080/_ah/api/rest/v0/linkedin/callback",
	AuthURL:      "https://www.linkedin.com/uas/oauth2/authorization",
	TokenURL:     "https://www.linkedin.com/uas/oauth2/accessToken",
	Scopes: []string{"r_basicprofile", "r_emailaddress"},
}

func (s *Service) LinkedInOauth(r *http.Request, req *NoRequest, resp *ResponseOauth) error {
	resp.Url=transport.AuthCodeURL("LinkedIn")
	return nil
}

func (s *Service) LinkedInCallback(r *http.Request, req *RequestCallback, resp *Response) error {
	context := endpoints.NewContext(r)
	transport.Transport=&urlfetch.Transport{
		Context: context,
		Deadline: 0,
		AllowInvalidServerCertificate: false,
	}
	transport.TokenCache=&Cache{
		Context: context,
		Key: "oauth2_basicprofile_emailaddress",
	}
	transport.Client= urlfetch.Client(context)
	transport.Context=context

	//if err := transport.FetchToken(); err != nil {return err}
	token, err := transport.Exchange(req.Code)
	transport.Token = token
	u, err := LinkedinUser(transport)
	resp.Message="LinkedIn email: "+u.Email
	return err
}

func LinkedinUser(transport *Transport) (u *UserL, err error) {
	r, err := transport.Client.Get("https://api.linkedin.com/v1/people/~/email-address?oauth2_access_token="+transport.Token.AccessToken)
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	log.Print(string(b))
	err = xml.Unmarshal(b, &u)
	return
}
