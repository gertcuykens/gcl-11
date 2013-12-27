package bin

import (
    "fmt"
    "net/http"
    "appengine"
    "appengine/user"
    "facebook"
)

func welcome(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        url, _ := user.LoginURL(c, "/")
        fmt.Fprintf(w, `<a href="%s">Sign in google</a>`, url)
        return
    }
    url, _ := user.LogoutURL(c, "/")
    fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}

func welcome2(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u, err := user.CurrentOAuth(c, "https://www.googleapis.com/auth/userinfo.email")
    if err != nil {
        http.Error(w, "OAuth Authorization header required", http.StatusUnauthorized)
        return
    }
    /*if !u.Admin {
        http.Error(w, "Admin login only", http.StatusUnauthorized)
        return
    }*/
    fmt.Fprintf(w, `Welcome, %s!`, u)
}

func welcome3(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "text/html; charset=utf-8")
    c := appengine.NewContext(r)
    g := user.Current(c)
    if g != nil {
        url, _ := user.LogoutURL(c, "/")
        fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)</br>`, g, url)
    }

    t:= facebook.AccessToken{Token:"token", Expiry:123}
    f:= facebook.GetMe(t)
    fmt.Fprintf(w, `Welcome, %s! from facebook</br>`, f)
}

/*
func init() {
    http.HandleFunc("/_ah/login_required", openIdHandler)
}

func openIdHandler(w http.ResponseWriter, r *http.Request) {
    // ...
}
*/
