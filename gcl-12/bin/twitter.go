package bin

import (
	"log"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"encoding/json"
)

func TwitterOauth() (requestToken *oauth.RequestToken, url string, err error) {
	requestToken, url, err = consumer.GetRequestTokenAndUrl("oob")
	if err != nil {log.Fatal(err)}
    return
}

func TwitterApi(requestToken *oauth.RequestToken, verificationCode string) (u *UserT, err error) {
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
