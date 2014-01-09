package bin

import "github.com/mrjones/oauth"

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

type NoRequest struct {}
