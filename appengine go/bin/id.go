package bin

import ("github.com/crhym3/go-endpoints/endpoints")

const WEB_CLIENT_ID string = "522156758812-09f5qv0e4gqjdjqfocerqcud5m5jutau.apps.googleusercontent.com"
const ANDROID_CLIENT_ID_d string = "522156758812-speqt3cnr7ggje0r3hhjtjg14iigru1f.apps.googleusercontent.com"
const ANDROID_CLIENT_ID_r string = "522156758812-29jkcaiofrismobslc4ioop1dvfhhgoi.apps.googleusercontent.com"
const SCOPE string = "https://www.googleapis.com/auth/userinfo.email"

var CLIENTIDS = []string{WEB_CLIENT_ID, ANDROID_CLIENT_ID_d, ANDROID_CLIENT_ID_r, endpoints.ApiExplorerClientId}
var AUDIENCES = []string{WEB_CLIENT_ID}
var SCOPES = []string{SCOPE}

