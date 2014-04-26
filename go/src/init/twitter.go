package init

import (
	"log"
	"github.com/mrjones/oauth"
	"github.com/crhym3/go-endpoints/endpoints"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"appengine/urlfetch"
)

type RequestCallback struct {
	Code string `json:"code"`
	State string `json:"state"`
}

type RequestOauth struct {
	Oauth_token string `endpoints:"oauth_token"`
	Oauth_verifier string `endpoints:"oauth_verifier"`
}

type RequestOob struct {
	RequestToken *oauth.RequestToken `json:"requestToken"`
	VerificationCode string `json:"verificationCode"`
}

type ResponseOauth struct {
	RequestToken *oauth.RequestToken `json:"requestToken"`
	Url string `json:"url"`
}

const TWITTER_ID = "PrTNrRxkWs6dw9XOr95A"
const TWITTER_SECRET = ""

var TWITTER_SERVER = oauth.ServiceProvider{
	RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
	AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
	AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
}

var consumer = oauth.NewConsumer(TWITTER_ID, TWITTER_SECRET, TWITTER_SERVER)
//consumer.Debug(true)

func (s *Service) TwitterOauth(r *http.Request, req *NoRequest, resp *ResponseOauth) error {
	c := endpoints.NewContext(r)
	consumer.HttpClient=urlfetch.Client(c)
	requestToken, url, err := consumer.GetRequestTokenAndUrl("http://localhost:8080/_ah/api/rest/v0/twitter/callback")
	if err != nil {log.Fatal(err)}
	resp.RequestToken=requestToken
	resp.Url=url
	return err
}

func (s *Service) TwitterOauthOob(r *http.Request, req *NoRequest, resp *ResponseOauth) error {
	c := endpoints.NewContext(r)
	consumer.HttpClient=urlfetch.Client(c)
	requestToken, url, err := consumer.GetRequestTokenAndUrl("oob")
	if err != nil {log.Fatal(err)}
	resp.RequestToken=requestToken
	resp.Url=url
	return err
}

func (s *Service) TwitterCallback(r *http.Request, req *RequestOauth, resp *Token) error {
	c := endpoints.NewContext(r)
	consumer.HttpClient=urlfetch.Client(c)
	resp.Oauth_token = req.Oauth_token
	resp.Oauth_verifier = req.Oauth_verifier
	err := TwitterUser(resp)
	return err
}

func (s *Service) TwitterCallbackOob(r *http.Request, req *RequestOob, resp *Token) error {
	c := endpoints.NewContext(r)
	consumer.HttpClient=urlfetch.Client(c)
	resp.Oauth_token = req.RequestToken.Token
	resp.Oauth_verifier = req.VerificationCode
	err := TwitterUser(resp)
	return err
}

func TwitterUser(t *Token) (err error) {
	var u struct {
		Id int64
		Name string
	}
	requestToken := &oauth.RequestToken{Token:t.Oauth_token}
	accessToken, err := consumer.AuthorizeToken(requestToken, t.Oauth_verifier)
	if err != nil {log.Fatal(err)}
	response, err := consumer.Get("https://api.twitter.com/1.1/account/verify_credentials.json", nil, accessToken)
	if err != nil {log.Fatal(err)}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(b, &u)
	t.Id = u.Id
	t.Name = u.Name
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
