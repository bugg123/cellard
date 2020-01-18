package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const searchString = "Bourbon County Brand Stout (2019)"

func TestAddBeerSearchQuery(t *testing.T) {
	c := NewClient(nil)
	t.Run("search simple query", func(t *testing.T) {
		beers, err := c.Beer.SearchBeerQuery(searchString, 10)
		assert.NoError(t, err)

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
		assert.NoError(t, err)
		sevenBeers, err := c.Beer.SearchBeerQuery(searchString, 7)
		assert.NoError(t, err)
		assert.Len(t, *fiveBeers, 5)
		assert.Len(t, *sevenBeers, 7)
	})

}

func TestGetBeerInfo(t *testing.T) {

	c := NewClient(nil)
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
