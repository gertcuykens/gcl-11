package bin

import (
	"github.com/crhym3/go-endpoints/endpoints"
	"github.com/mrjones/oauth"
)

const WEB_CLIENT_ID string = "1034966141188-b4cup6jccsjqpdc14c9218fhb488e515.apps.googleusercontent.com"
const ANDROID_CLIENT_ID_d string = ""
const ANDROID_CLIENT_ID_r string = ""
const SCOPE string = "https://www.googleapis.com/auth/userinfo.email"

var CLIENTIDS = []string{WEB_CLIENT_ID, ANDROID_CLIENT_ID_d, ANDROID_CLIENT_ID_r, endpoints.ApiExplorerClientId}
var AUDIENCES = []string{WEB_CLIENT_ID}
var SCOPES = []string{SCOPE}

const TWITTER_ID = "PrTNrRxkWs6dw9XOr95A"
const TWITTER_SECRET = ""
var   TWITTER_SERVER = oauth.ServiceProvider{ RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
					                          AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
                                              AccessTokenUrl:    "https://api.twitter.com/oauth/access_token", }

var consumer = oauth.NewConsumer(TWITTER_ID, TWITTER_SECRET, TWITTER_SERVER)
