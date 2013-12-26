package bin

import "github.com/crhym3/go-endpoints/endpoints"

const FACEBOOK_ID string = "..."
const WEB_CLIENT_ID string = "...apps.googleusercontent.com"
const ANDROID_CLIENT_ID_d string = ""
const ANDROID_CLIENT_ID_r string = ""
const SCOPE string = "https://www.googleapis.com/auth/userinfo.email"

var CLIENTIDS = []string{WEB_CLIENT_ID, ANDROID_CLIENT_ID_d, ANDROID_CLIENT_ID_r, endpoints.ApiExplorerClientId}
var AUDIENCES = []string{WEB_CLIENT_ID}
var SCOPES = []string{SCOPE}

