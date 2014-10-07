package main

import (
	"net/http"
	"log"
	"io"
	"encoding/json"
	"io/ioutil"
	"crypto/sha1"
	"fmt"
	"encoding/base64"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"bytes"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Authorization string `json:"-"`
	Server string `json:"-"`
}

const PRIVATE_KEY string = "00000000"

var (
	privateKey []byte
	publicKey []byte
)

//openssl genrsa -out demo.rsa 1024
//openssl rsa -in demo.rsa -pubout > demo.rsa.pub
func main() {
	privateKey, _ = ioutil.ReadFile("demo.rsa")
	publicKey, _ = ioutil.ReadFile("demo.rsa.pub")
	index := http.FileServer(http.Dir("oauth"))
	http.Handle("/", index)
	http.HandleFunc("/jwt", authJwt)
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/google", connect)
	http.HandleFunc("/facebook", connect)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func check(e error) {
	if e != nil {panic(e)}
}

/**************************************************************/

func (u *User) get() (err error) {
	var req *http.Request;
	switch u.Server{
	case "/google":
		req, err = http.NewRequest("GET", "https://www.googleapis.com/userinfo/v2/me", nil)
	case "/facebook":
		req, err = http.NewRequest("GET", "https://graph.facebook.com/me?fields=id,name,email", nil)
	default:
		return
	}
	req.Header = map[string][]string{"Authorization": {u.Authorization}}
	client := &http.Client{}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	log.Printf("USER ============\n%s\n============",string(b))
	err = json.Unmarshal(b,u)
	return
}

func connect(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL ============\n%s\n============",r.URL.Path)
	var t = r.Header.Get("Authorization")
	log.Printf("TOKEN ============\n%s\n============",t[7:])
	var u = &User{
		Authorization: t,
		Server: r.URL.Path,
	}
	u.get()
	log.Printf("USER ============\n%v\n============",u)

	h := sha1.New()
	data:= fmt.Sprintf("{oauth:\"Google\", id:\"%s\", scope:\"%s\"}", u.Id, Scope(t))
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", base64.URLEncoding.EncodeToString(h.Sum([]byte(data+PRIVATE_KEY))))
	io.WriteString(w, data)
}

/*****************************************************************************/
// Create a Token that will be signed with RSA 256.
//{
//   "typ":"JWT",
//   "alg":"RS256"
//}

func authJwt(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims["ID"] = "This is my super fake ID"
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, _ := token.SignedString(privateKey)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+tokenString)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"token\": %s}", tokenString)
}

func auth(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		var b bytes.Buffer
		b.Write(publicKey)
		return b, nil
	})
	if err == nil && token.Valid {
		//OK
	} else {
		//NOT OK
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"token\": %v}", token)
}

/*****************************************************************************/

type Data struct {
	Name string `json:"name"`
	Perms []string `json:"perms"`
}

type Accounts struct {
	Data []Data `json:"data"`
	Authorization string `json:"-"`
}

func (a *Accounts) set() error {
	req, err := http.NewRequest("GET", "https://graph.facebook.com/me/accounts", nil) //&access_token=
	if err != nil {return nil}
	req.Header = map[string][]string{"Authorization": {a.Authorization}}
	client := &http.Client{}
	buf, err := client.Do(req)
	defer buf.Body.Close()
	b, err := ioutil.ReadAll(buf.Body)
	err = json.Unmarshal(b,a)
	return err
}

func (a *Accounts) Editor() bool {
	for i,x := range a.Data {
		if x.Name=="Gcl-11" {
			for _,y := range a.Data[i].Perms {
				if y == "CREATE_CONTENT" {
					return true
				}
			}
		}
	}
	return false
}

func Scope(t string) string{
	return "{user:true}"
	/*
	var a = &Accounts{
		Authorization: t,
		//Context: c,
	}
	a.set()
	scope := fmt.Sprintf("{user:\"%t\"}", a.Editor())
	return scope
	*/
}
