package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
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
