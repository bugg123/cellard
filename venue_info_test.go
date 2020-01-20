package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVenueInfo(t *testing.T) {
	venueID := 1222

	t.Run("venue info with valid id", func(t *testing.T) {
		c, done := venueInfoTestClient(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) {
			path := venueInfoPath + strconv.Itoa(venueID)
			if p := r.URL.Path; p != path {
				t.Fatalf("unexpected url path, got: %q want: %q", p, path)
			}
			infoJSON, err := ioutil.ReadFile("json/venue_info/venue_info.json")
			assert.NoErrorf(t, err, "couldn't read JSON file %v", err)
			w.Write(infoJSON)
		})
		defer done()

		got, resp, err := c.Venue.GetVenueInfo(venueID)

		want := Venue{
			VenueID:   1222,
			VenueName: "Untappd HQ - East",
		}

		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusOK)
		assert.Equal(t, got.VenueID, want.VenueID)
		assert.Equal(t, got.VenueName, want.VenueName)
	})

	t.Run("venue info with invalid id", func(t *testing.T) {

		c, done := venueInfoTestClient(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})
		defer done()
		venue, resp, err := c.Venue.GetVenueInfo(78947589739847)
		assert.Error(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusNotFound)
		assert.Nil(t, venue)
	})

}

func venueInfoTestClient(t *testing.T, fn func(t *testing.T, w http.ResponseWriter, r *http.Request)) (*Client, func()) {
	return testClient(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("expected GET but got unexpected http method: %q", r.Method)
		}

		if p := r.URL.Path; !strings.HasPrefix(p, venueInfoPath) {
			t.Fatalf("expected %q to have prefix %q", p, venueInfoPath)
		}

		if fn != nil {
			fn(t, w, r)
		}
	})

}
