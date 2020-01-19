package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const searchString = "Bourbon County Brand Stout (2019)"

func TestBeerSearchQuery(t *testing.T) {
	c, done := beerSearchTestClient(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) {
		if q := r.URL.Query().Get("q"); q != searchString {
			t.Fatalf("search query doesn't match, got: %q want: %q", q, searchString)
		}

		var (
			searchJSON []byte
			err        error
		)
		if r.URL.Query().Get("limit") == "5" {
			searchJSON, err = ioutil.ReadFile("json/beer_search/beer_search_5.json")
		} else {
			searchJSON, err = ioutil.ReadFile("json/beer_search/beer_search.json")
		}
		assert.NoErrorf(t, err, "couldn't read JSON file %v", err)
		w.Write(searchJSON)
	})
	defer done()
	// c := NewClient(nil)
	t.Run("search simple query", func(t *testing.T) {
		beers, err := c.Beer.SearchBeerQuery(searchString, 10)
		assert.NotNil(t, beers)
		assert.NoError(t, err)
		if t.Failed() {
			t.FailNow()
		}

		//Check default limit is returned
		assert.Len(t, *beers, 10)
		var bcbs Beer
		for _, beer := range *beers {
			if int(beer.BID) == 3507187 {
				bcbs = beer
			}
		}
		//Should be able to get the beer most related to search
		assert.NotNil(t, bcbs)
	})
	t.Run("check limit sizes", func(t *testing.T) {
		fiveBeers, err := c.Beer.SearchBeerQuery(searchString, 5)
		assert.NotNil(t, fiveBeers)
		assert.NoError(t, err)
		if t.Failed() {
			t.FailNow()
		}
		assert.Len(t, *fiveBeers, 5)
	})
}

func beerSearchTestClient(t *testing.T, fn func(t *testing.T, w http.ResponseWriter, r *http.Request)) (*Client, func()) {
	return testClient(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) {

		if m := r.Method; m != http.MethodGet {
			t.Fatalf("expected GET but got unexpected http method: %q", m)
		}

		prefix := beerSearchPath
		if p := r.URL.Path; !strings.HasPrefix(p, prefix) {
			t.Fatalf("expected %q to have prefix %q", p, prefix)
		}
		if fn != nil {
			fn(t, w, r)
		}
	})
}
