package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func testAuthHandler(t *testing.T, redirectURL, oauthHost string, fn TokenHandlerFunc) (string, func()) {
	h, _, err := NewAuthHandler(
		"foo",
		"bar",
		redirectURL,
		fn,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	h.oAuthURL.Scheme = "http"
	h.oAuthURL.Host = oauthHost

	srv := httptest.NewServer(h)
	return srv.URL, func() {
		srv.Close()
	}
}

func testOAuthHandler(t *testing.T, fn func(t *testing.T, w http.ResponseWriter, r *http.Request)) (string, func()) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)

		if fn != nil {
			fn(t, w, r)
		}
	}))

	u, err := url.Parse(srv.URL)
	if err != nil {
		log.Fatal(err)
	}

	return u.Host, func() {
		srv.Close()
	}
}

func testOAuthBadGateway(t *testing.T, fn func(t *testing.T, w http.ResponseWriter, r *http.Request)) {
	oauthHost, done := testOAuthHandler(t, fn)
	defer done()

	url, done := testAuthHandler(t, "http://fun.com", oauthHost, nil)
	defer done()

	res, err := http.Get(url + "?code=")
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res.StatusCode, http.StatusBadGateway; got != want {
		t.Fatalf("unexpected HTTP status code: got %d, want %d", got, want)
	}
}
