package bin

import (
	"net/url"
	"mime"
	"io/ioutil"
	"encoding/json"
	"time"
	"strings"
)

type Transport struct {
	Token *Token
	TokenCache *Cache
    ClientId string
    ClientSecret string
    RedirectURL string
    AuthURL string
    TokenURL string
	Scopes []string
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

func (t *Transport) Exchange(code string) (err error) {
	if t.Token == nil && t.TokenCache != nil {t.Token, _ = t.TokenCache.Token()}
	if t.Token == nil {t.Token = new(Token)}
	err = t.UpdateToken(url.Values{
		"grant_type":   {"authorization_code"},
		"redirect_uri": {t.RedirectURL},
		"scope":        {strings.Join(t.Scopes, " ")},
		"code":         {code},
	})
	if err != nil {return err}
	if t.TokenCache != nil {return t.TokenCache.PutToken(t.Token)}
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

func (t *Transport) Refresh() error {
	if t.Token == nil {return OAuthError{"Refresh", "no existing Token"}}
	if t.Token.Refresh == "" {return OAuthError{"Refresh", "Token expired; no Refresh Token"}}

	err := t.UpdateToken(url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {t.Token.Refresh},
	})

	if err != nil {return err}
	if t.TokenCache != nil {return t.TokenCache.PutToken(t.Token)}
	return nil
}

func (t *Transport) UpdateToken(v url.Values) error {
	v.Set("client_id", t.ClientId)
	v.Set("client_secret", t.ClientSecret)
	r, err := t.Token.Client.PostForm(t.TokenURL, v)
	if err != nil {return err}
	defer r.Body.Close()
	if r.StatusCode != 200 {return OAuthError{"updateToken", r.Status}}
	var b Token
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
	t.Token.Access = b.Access
	if len(b.Refresh) > 0 {t.Token.Refresh = b.Refresh}
	if b.ExpiresIn == 0 {t.Token.Expiry = time.Time{}} else { t.Token.Expiry = time.Now().Add(b.ExpiresIn) }
	if b.Id != "" {t.Token.Id = b.Id}
	return nil
}

/*
func (t *Transport) FetchToken() error {
	if t.Token == nil && t.TokenCache != nil {t.Token, _ = t.TokenCache.Token()}
	if t.Token == nil || t.Token.Expired() {if err := t.Refresh(); err != nil {return err}}
	return nil
}

func (t *Transport) Refresh() error {
	tok, expiry, err := appengine.AccessToken(t.Token.Context, t.Scopes...)
	if err != nil {return err}
	t.Token = &Token{
		Access: tok,
		Expiry: expiry,
	}
	if t.TokenCache != nil {t.TokenCache.PutToken(t.Token)}
	return nil
}

Transport http.RoundTripper
r, err := (&http.Client{Transport: t.transport()}).PostForm(t.TokenURL, v)

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

if err := transport.FetchToken(); err != nil {return err}

transport.Transport=&urlfetch.Transport{
	Context: context,
	Deadline: 0,
	AllowInvalidServerCertificate: false,
}
*/
