package bin

type Request struct {
  Token string `json:"token"`
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

