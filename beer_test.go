package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const searchString = "Bourbon County Brand Stout (2019)"

func TestAddBeerSearchQuery(t *testing.T) {
	c := NewClient(nil)
	t.Run("search simple query", func(t *testing.T) {
		beers, err := c.Beer.SearchBeerQuery(searchString, 10)
		assert.NoError(t, err)

		fmt.Println("Searched: Bourbon County Brand Stout (2019)")

		//Check default limit is returned
		assert.Len(t, *beers, 10)
		var bcbs Beer
		for _, beer := range *beers {
			fmt.Printf("Found beer: %s with ID: %f\n", beer.BeerName, beer.BID)
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
