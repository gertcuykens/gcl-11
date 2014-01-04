package bin

import (
	"log"
	"github.com/mrjones/oauth"
	"github.com/crhym3/go-endpoints/endpoints"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"appengine/urlfetch"
)

type UserT struct {
	Name string
	Id int64
}

const TWITTER_ID = "PrTNrRxkWs6dw9XOr95A"

var TWITTER_SERVER = oauth.ServiceProvider{
	RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
	AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
	AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
}

var consumer = oauth.NewConsumer(TWITTER_ID, TWITTER_SECRET, TWITTER_SERVER)

func (s *Service) TwitterOauth(r *http.Request, req *NoRequest, resp *ResponseOauth) error {
	c := endpoints.NewContext(r)
	consumer.HttpClient=urlfetch.Client(c)
	requestToken, url, err := Oauth()
	resp.RequestToken=requestToken
	resp.Url=url
	return err
}

func (s *Service) TwitterOauthOob(r *http.Request, req *NoRequest, resp *ResponseOauth) error {
	c := endpoints.NewContext(r)
	consumer.HttpClient=urlfetch.Client(c)
	requestToken, url, err := Oauth()
	resp.RequestToken=requestToken
	resp.Url=url
	return err
}

func (s *Service) TwitterCallback(r *http.Request, req *RequestOauth, resp *Response) error {
	c := endpoints.NewContext(r)
	consumer.HttpClient=urlfetch.Client(c)
	requestToken := &oauth.RequestToken{Token:req.Oauth_token}
	b, err := TwitterUser(requestToken, req.Oauth_verifier)
	resp.Message=b.Name
	return err
}

func (s *Service) TwitterCallbackOob(r *http.Request, req *RequestOob, resp *Response) error {
	c := endpoints.NewContext(r)
	consumer.HttpClient=urlfetch.Client(c)
	b, err := TwitterUser(req.RequestToken, req.VerificationCode)
	resp.Message=b.Name
	return err
}

func Oauth() (requestToken *oauth.RequestToken, url string, err error) {
	//tokenUrl := "https://gcl-12.appspot.com/_ah/api/rest/v0/twitter/callback"
	tokenUrl := "http://localhost:8080/_ah/api/rest/v0/twitter/callback"
	requestToken, url, err = consumer.GetRequestTokenAndUrl(tokenUrl)
	if err != nil {log.Fatal(err)}
	return
}

func OauthOob() (requestToken *oauth.RequestToken, url string, err error) {
	requestToken, url, err = consumer.GetRequestTokenAndUrl("oob")
	if err != nil {log.Fatal(err)}
	return
}

func TwitterUser(requestToken *oauth.RequestToken, verificationCode string) (u *UserT, err error) {
	accessToken, err := consumer.AuthorizeToken(requestToken, verificationCode)
	if err != nil {log.Fatal(err)}
	response, err := consumer.Get("https://api.twitter.com/1.1/account/verify_credentials.json", nil, accessToken)
	if err != nil {log.Fatal(err)}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(b, &u)
	return
}

/*
response, err := consumer.Get("http://api.twitter.com/1.1/statuses/home_timeline.json", map[string]string{"count": "1"}, accessToken)

	if *postUpdate {
		status := fmt.Sprintf("Test post via the API using Go (http://golang.org/) at %s", time.Now().String())
		response, err = c.Post("http://api.twitter.com/1.1/statuses/update.json",map[string]string{"status": status,},accessToken)
		if err != nil {log.Fatal(err)}
		defer response.Body.Close()
	}
*/
