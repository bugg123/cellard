package main

import (
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
		w.Write(hocusPocusJSON)
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

var hocusPocusJSON = []byte(`
{
	"meta": {
		"code": 200,
		"response_time": {
			"time": 0,
			"measure": "seconds"
		}
	},
	"response":{
		"beer":{
			"bid":1,
			"beer_name":"Hocus Pocus",
			"beer_label":"https:\/\/untappd.akamaized.net\/site\/beer_logos\/beer-1_d4bd9_sm.jpeg",
			"beer_label_hd":"https:\/\/untappd.akamaized.net\/site\/beer_logos_hd\/beer-1_55f47_hd.jpeg",
			"beer_abv":4.5,
			"beer_ibu":13,
			"beer_description":"Our take on a classic summer ale. A toast to weeds, rays, and summer haze. A light, crisp ale for mowing lawns, hitting lazy fly balls, and communing with nature, Hocus Pocus is offered up as a summer sacrifice to cloudless days.\r\n\r\nIts malty sweetness finishes tart and crisp and is best appreciated with a wedge of orange.",
			"beer_style":"Wheat Beer - American Pale Wheat",
			"is_in_production":1,
			"beer_slug":"magic-hat-brewing-company-hocus-pocus",
			"is_homebrew":0,
			"created_at":"Sat, 21 Aug 2010 14:26:35 +0000",
			"rating_count":16935,
			"rating_score":3.27881,
			"stats":{
				"total_count":23313,
				"monthly_count":14,
				"total_user_count":19541,
				"user_count":0
			}
		}
	}
}`)
