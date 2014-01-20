//CREATE NEW FILE secret.go IN bin
package bin

import (
	"code.google.com/p/goauth2/oauth"
	"time"
)

const WEB_CLIENT_ID string = "...apps.googleusercontent.com"
const ANDROID_CLIENT_ID string = "...apps.googleusercontent.com"
const SERVER_CLIENT_ID string = "...apps.googleusercontent.com"

const SCOPE1 string = "https://www.googleapis.com/auth/userinfo.email"
const SCOPE2 string = "https://www.googleapis.com/auth/devstorage.full_control"
const SCOPE3 string = "https://www.googleapis.com/auth/androidpublisher"

const R = "..."

var CONFIG = &oauth.Config{
	ClientId: SERVER_CLIENT_ID,
	ClientSecret: "...",
	Scope:        SCOPE1+" "+SCOPE2+" "+SCOPE3,
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://accounts.google.com/o/oauth2/token",
}

var TOKEN = &oauth.Token{
	AccessToken: "",
	RefreshToken: "...",
	Expiry: time.Now(),
	Extra: nil,
}

var TRANS = &oauth.Transport{
	Token: TOKEN,
	Config: CONFIG,
}
