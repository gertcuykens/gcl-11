package bin

import "github.com/mrjones/oauth"

type Request struct {
	Access_token string `json:"access_token"`
}

type RequestOauth struct {
	RequestToken *oauth.RequestToken `json:"requestToken"`
    VerificationCode string `json:"verificationCode"`
}

type Response struct {
	Message string `json:"message"`
}

type ResponseOauth struct {
	RequestToken *oauth.RequestToken `json:"requestToken"`
	Url string `json:"url"`
}

type Service struct {}

type User struct {
	Name string
	Id string
}

type NoRequest struct {}


