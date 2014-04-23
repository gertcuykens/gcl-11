package bin

import "github.com/crhym3/go-endpoints/endpoints"

type Service struct {}

func init() {
	service := &Service{}
	scope := []string{"https://www.googleapis.com/auth/userinfo.email"}
	api, err := endpoints.RegisterService(service, "rest", "v0", "API", true)
	if err != nil {panic(err.Error())}
	info1 := api.MethodByName("GoogleUserService").Info()
	info1.Name, info1.HttpMethod, info1.Path, info1.Desc, info1.Scopes = "google.user", "POST", "google/user", "Oauth2 google user.", scope
	info2 := api.MethodByName("GoogleRevokeService").Info()
	info2.Name, info2.HttpMethod, info2.Path, info2.Desc = "google.revoke", "POST", "google/revoke", "Oauth2 google revoke."
	info3 := api.MethodByName("GooglePurchasesService").Info()
	info3.Name, info3.HttpMethod, info3.Path, info3.Desc, info3.Scopes= "google.purchases", "POST", "google/purchases", "Oauth2 google purchases.", scope
	info4 := api.MethodByName("GoogleStorageService").Info()
	info4.Name, info4.HttpMethod, info4.Path, info4.Desc, info4.Scopes = "google.storage", "GET", "google/storage", "Oauth2 google storage.", scope
	endpoints.HandleHttp()
}
