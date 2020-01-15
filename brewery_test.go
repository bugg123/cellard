package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBreweryInfo(t *testing.T) {

	c := NewClient(nil)
	got, err := c.Brewery.GetBreweryInfo(1)
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
