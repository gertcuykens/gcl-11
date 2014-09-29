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
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Authorization string `json:"-"`
	Server string `json:"-"`
	//Context endpoints.Context `json:"-"`
}

const PRIVATE_KEY string = "00000000"

func check(e error) {
	if e != nil {panic(e)}
}

func main() {
	index := http.FileServer(http.Dir("oauth"))
	http.Handle("/", index)
	http.HandleFunc("/google", connect)
	http.HandleFunc("/facebook", connect)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/**************************************************************/

func (u *User) get() (err error) {
	//client := urlfetch.Client(u.Context)
	var req *http.Request;
	client := &http.Client{}
	switch u.Server{
	case "/google":
		req, err = http.NewRequest("GET", "https://www.googleapis.com/userinfo/v2/me", nil)
	case "/facebook":
		req, err = http.NewRequest("GET", "https://graph.facebook.com/me?fields=id,name,email", nil)
	default:
		return
	}
	req.Header = map[string][]string{"Authorization": {u.Authorization}}
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
		//Context: c,
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

type Data struct {
	Name string `json:"name"`
	Perms []string `json:"perms"`
}

type Accounts struct {
	Data []Data `json:"data"`
	Authorization string `json:"-"`
	//Context endpoints.Context `json:"-"`
}

func (a *Accounts) set() error {
	//client := urlfetch.Client(a.Context)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://graph.facebook.com/me/accounts", nil) //&access_token=
	if err != nil {return nil}
	req.Header = map[string][]string{"Authorization": {a.Authorization}}
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

/*****************************************************************************/

