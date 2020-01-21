package main

import (
	"net/http/httptest"
	"testing"
)

func testAuthHandler(t *testing.T, redirectURL, oauthHost string, fn TokenHandlerFunc) (string, func()) {
	h, _, err := NewAuthHandler{
		"foo",
		"bar",
		redirectURL,
		fn,
		nil,
	}
	if err != nil {
		t.Fatal(err)
	}

	h.oAuthUrl.scheme = "http"
	h.oAuthUrl.Host = oauthHost

	srv := httptest.NewServer(h)
	return srv.URL, func() {
		srv.Close()
	}
}
