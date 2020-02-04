package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	untappdOAuthAuthenticate = "https://untappd.com/oauth/authenticate/?client_id=%s&response_type=code&redirect_url=%s"
	untappdOAuthAuthorize    = "https://untappd.com/oauth/authorize/?client_id=%s&client_secret=%s&response_type=code&redirect_url=%s"
)

var (
	ErrNoClientID     = errors.New("no client ID")
	ErrNoClientSecret = errors.New("no client secret")
	ErrNoAccessToken  = errors.New("no access token")
)

type AuthService struct {
	client *Client
}

type AuthHandler struct {
	clientID     string
	clientSecret string
	redirectURL  *url.URL
	oAuthURL     *url.URL
	handler      TokenHandlerFunc
	client       *http.Client
}

type TokenHandlerFunc func(token string, w http.ResponseWriter, r *http.Request)

var defaultTokenFn = func(token string, w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(token)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewAuthHandler(clientID, clientSecret, redirectURL string, fn TokenHandlerFunc, client *http.Client) (*AuthHandler, *url.URL, error) {
	if clientID == "" {
		return nil, nil, ErrNoClientID
	}
	if clientSecret == "" {
		return nil, nil, ErrNoClientSecret
	}
	ru, err := url.Parse(redirectURL)
	if err != nil {
		return nil, nil, err
	}

	cu, err := url.Parse(fmt.Sprintf(
		untappdOAuthAuthenticate,
		clientID,
		ru.String(),
	))
	if err != nil {
		return nil, nil, err
	}

	ou, err := url.Parse(fmt.Sprintf(
		untappdOAuthAuthorize,
		clientID,
		clientSecret,
		ru.String(),
	))
	if err != nil {
		return nil, nil, err
	}

	if fn == nil {
		fn = defaultTokenFn
	}

	if client == nil {
		client = http.DefaultClient
	}

	return &AuthHandler{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURL:  ru,
		oAuthURL:     ou,
		handler:      fn,
		client:       client,
	}, cu, nil
}

func (a *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "no 'code' GET parameter", http.StatusBadRequest)
		return
	}

	res, err := a.client.Get(a.oAuthURL.String() + "&code=" + code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	if c := res.StatusCode; c > 299 || c < 200 {
		http.Error(w, fmt.Sprintf("authentication server error: HTTP %03d", c), http.StatusBadGateway)
		return
	}

	if !strings.Contains(res.Header.Get("Content-Type"), jsonContentType) {
		http.Error(w, "authentication server sent non-JSON conten", http.StatusBadGateway)
	}

	var v struct {
		Response struct {
			AccessToken string `json:"access_token"`
		} `json:"response"`
	}

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	a.handler(v.Response.AccessToken, w, r)
}
