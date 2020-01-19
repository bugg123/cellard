package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeerInfo(t *testing.T) {
	beerID := 1

	c, done := beerInfoTestClient(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) {
		path := beerInfoPath + strconv.Itoa(beerID)
		if p := r.URL.Path; p != path {
			t.Fatalf("unexpected url path, got: %q want: %q", p, path)
		}
		infoJSON, err := ioutil.ReadFile("json/beer_info/beer_info.json")
		assert.NoErrorf(t, err, "couldn't read JSON file %v", err)
		w.Write(infoJSON)
	})
	defer done()

	//c := NewClient(nil)
	got, err := c.Beer.GetBeerInfo(1)

	want := Beer{
		BID:      1,
		BeerName: "Hocus Pocus",
		BeerSlug: "magic-hat-brewing-company-hocus-pocus",
	}

	assert.NoError(t, err)
	assert.Equal(t, want.BID, got.BID)
	assert.Equal(t, want.BeerName, got.BeerName)
	assert.Equal(t, want.BeerSlug, got.BeerSlug)

}

func beerInfoTestClient(t *testing.T, fn func(t *testing.T, w http.ResponseWriter, r *http.Request)) (*Client, func()) {
	return testClient(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) {

		if m := r.Method; m != http.MethodGet {
			t.Fatalf("expected GET but got unexpected http method: %q", m)
		}

		prefix := "/v4/beer/info/"
		if p := r.URL.Path; !strings.HasPrefix(p, prefix) {
			t.Fatalf("expected %q to have prefix %q", p, prefix)
		}

		if fn != nil {
			fn(t, w, r)
		}
	})
}
