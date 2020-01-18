package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func testClient(t *testing.T, fn func(t *testing.T, w http.ResponseWriter, r *http.Request)) (*Client, func()) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if fn != nil {
			fn(t, w, r)
		}
	}))

	client := NewClient(nil)
	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	client.BaseURL = u
	return client, func() {
		server.Close()
	}

}
