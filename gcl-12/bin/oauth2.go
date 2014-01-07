package bin

import (
	"appengine"
	"net/url"
	"net/http"
	"mime"
	"io/ioutil"
	"encoding/json"
	"time"
	"strings"
)

type Transport struct {
	*Token
    ClientId string
    ClientSecret string
    RedirectURL string
    AuthURL string
    TokenURL string
	Context appengine.Context
	Client *http.Client
	Scopes []string
	TokenCache *Cache
	ApprovalPrompt string
	AccessType string
}

type OAuthError struct {
	prefix string
	msg    string
}

func (oe OAuthError) Error() string {
	return "OAuthError: " + oe.prefix + ": " + oe.msg
}

func (t *Transport) Refresh() error {
	tok, expiry, err := appengine.AccessToken(t.Context, t.Scopes...)
	if err != nil {return err}
	t.Token = &Token{
		AccessToken: tok,
		Expiry: expiry,
	}
	if t.TokenCache != nil {t.TokenCache.PutToken(t.Token)}
	return nil
}

func (t *Transport) FetchToken() error {
	if t.Token == nil && t.TokenCache != nil {t.Token, _ = t.TokenCache.Token()}
	if t.Token == nil || t.Expired() {if err := t.Refresh(); err != nil {return err}}
	return nil
}

func (t *Transport) Exchange(code string) (*Token, error) {
	tok := t.Token
	if tok == nil && t.TokenCache != nil {tok, _ = t.TokenCache.Token()}
	if tok == nil {tok = new(Token)}
	err := t.UpdateToken(tok, url.Values{
		"grant_type":   {"authorization_code"},
		"redirect_uri": {t.RedirectURL},
		"scope":        {strings.Join(t.Scopes, " ")},
		"code":         {code},
	})
	if err != nil {return nil, err}
	t.Token = tok
	if t.TokenCache != nil {return tok, t.TokenCache.PutToken(tok)}
	return tok, nil
}

func (t *Transport) UpdateToken(tok *Token, v url.Values) error {
	v.Set("client_id", t.ClientId)
	v.Set("client_secret", t.ClientSecret)
	r, err := t.Client.PostForm(t.TokenURL, v)
	if err != nil {return err}
	defer r.Body.Close()
	if r.StatusCode != 200 {return OAuthError{"updateToken", r.Status}}
	var b struct {
		Access    string        `json:"access_token"`
		Refresh   string        `json:"refresh_token"`
		ExpiresIn time.Duration `json:"expires_in"`
		Id        string        `json:"id_token"`
	}
	content, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	switch content {
		case "application/x-www-form-urlencoded", "text/plain":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {return err}
			vals, err := url.ParseQuery(string(body))
			if err != nil {return err}
			b.Access = vals.Get("access_token")
			b.Refresh = vals.Get("refresh_token")
			b.ExpiresIn, _ = time.ParseDuration(vals.Get("expires_in") + "s")
			b.Id = vals.Get("id_token")
		default:
			if err = json.NewDecoder(r.Body).Decode(&b); err != nil {return err}
			b.ExpiresIn *= time.Second
	}
	tok.AccessToken = b.Access
	if len(b.Refresh) > 0 {tok.RefreshToken = b.Refresh}
	if b.ExpiresIn == 0 {tok.Expiry = time.Time{}} else { tok.Expiry = time.Now().Add(b.ExpiresIn) }
	if b.Id != "" {if tok.Extra == nil {tok.Extra = make(map[string]string)}; tok.Extra["id_token"] = b.Id}
	return nil
}

func (t *Transport) AuthCodeURL(state string) string {
	u, err := url.Parse(t.AuthURL)
	if err != nil {panic("AuthURL malformed: " + err.Error())}
	q := url.Values{
		"response_type":   {"code"},
		"client_id":       {t.ClientId},
		"redirect_uri":    {t.RedirectURL},
		"scope":           {strings.Join(t.Scopes, " ")},
		"state":           {state},
		"access_type":     {t.AccessType},
		"approval_prompt": {t.ApprovalPrompt},
	}.Encode()
	if u.RawQuery == "" {u.RawQuery = q} else {u.RawQuery += "&" + q}
	return u.String()
}

/*
//Transport http.RoundTripper
//r, err := (&http.Client{Transport: t.transport()}).PostForm(t.TokenURL, v)

func cloneRequest(r *http.Request) *http.Request {
	r2 := new(http.Request)
	*r2 = *r
	r2.Header = make(http.Header)
	for k, s := range r.Header {r2.Header[k] = s}
	return r2
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := t.FetchToken(); err != nil {return nil, err}
	newReq := cloneRequest(req)
	newReq.Header.Set("Authorization", "Bearer "+t.AccessToken)
	return t.Transport.RoundTrip(newReq)
}
*/
