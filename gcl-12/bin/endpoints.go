package bin

type Request struct {
	Access_token string `json:"access_token"`
}

type Response struct {
	Message string `json:"message"`
}

type Service struct {
}

type User struct {
	Name string
	Id string
}

