package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBreweryInfo(t *testing.T) {
	breweryID := 1

	//c := NewClient(nil)
	c, done := breweryInfoTestClient(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) {
		path := breweryInfoPath + strconv.Itoa(breweryID)
		if p := r.URL.Path; p != path {
			t.Fatalf("unexpected url path, got: %q want: %q", p, path)
		}
		infoJSON, err := ioutil.ReadFile("json/brewery_info/brewery_info.json")
		assert.NoErrorf(t, err, "couldn't read JSON file %v", err)
		w.Write(infoJSON)

	})
	defer done()
	got, err := c.Brewery.GetBreweryInfo(breweryID)
	want := Brewery{
		BreweryID:    1,
		BreweryName:  "(512) Brewing Company",
		BrewerySlug:  "512-brewing-company",
		BreweryLabel: "https://untappd.akamaized.net/site/brewery_logos/brewery-1_8ccec.jpeg",
	}

	assert.NoError(t, err)
	assert.Equal(t, want.BreweryID, got.BreweryID)
	assert.Equal(t, want.BreweryName, got.BreweryName)
	assert.Equal(t, want.BrewerySlug, got.BrewerySlug)
	assert.Equal(t, want.BreweryLabel, got.BreweryLabel)

}

func breweryInfoTestClient(t *testing.T, fn func(t *testing.T, w http.ResponseWriter, r *http.Request)) (*Client, func()) {
	return testClient(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) {

		if m := r.Method; m != http.MethodGet {
			t.Fatalf("expected GET but got unexpected http method: %q", m)
		}
		if p := r.URL.Path; !strings.HasPrefix(p, breweryInfoPath) {
		}

		if fn != nil {
			fn(t, w, r)
		}
	})

}
