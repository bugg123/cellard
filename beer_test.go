package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUntappdHost(t *testing.T) {
	got := GetUntappedHost()
	want, _ := url.Parse("https://api.untappd.com")

	assert.Equal(t, got, want)
}

func TestGetClientAuthString(t *testing.T) {

	got := GetClientAuthString()
	want := fmt.Sprintf("?client_id=%s&client_secret=%s", os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

	assert.Equal(t, got, want)
}

func TestSetClientAuthString(t *testing.T) {
	testMap := map[string]string{
		"client_id":     os.Getenv("CLIENT_ID"),
		"client_secret": os.Getenv("CLIENT_SECRET"),
	}

	untappdURL := GetUntappedHost()
	SetClientAuthString(untappdURL)

	for k, v := range untappdURL.Query() {
		assert.Contains(t, v, testMap[k])
	}
}

func TestAddBeerSearchQuery(t *testing.T) {

	t.Run("Test blank urls", func(t *testing.T) {
		got := url.URL{}
		AddBeerSearchQuery(&got, "testing123")
		want := url.URL{
			RawQuery: "q=testing123",
			Path:     beerSearchPath,
		}

		assert.Equal(t, want.Path, got.Path)
		assert.Equal(t, want.RawQuery, got.RawQuery)
		assert.Equal(t, got, want)
	})
	t.Run("test full urls", func(t *testing.T) {

		got := GetUntappedHost()
		AddBeerSearchQuery(got, "testing123")
		want := &url.URL{
			Scheme:   "https",
			Host:     "api.untappd.com",
			RawQuery: "q=testing123",
			Path:     beerSearchPath,
		}

		assert.Equal(t, got, want)
	})
}

func TestMakeQuery(t *testing.T) {

	gotURL := GetUntappedHost()
	SetClientAuthString(gotURL)
	GetBeerInfo(gotURL, 1)

	got, err := MakeQuery(gotURL)
	assert.NoError(t, err)

	want := Beer{
		BID:      1,
		BeerName: "Hocus Pocus",
		BeerSlug: "magic-hat-brewing-company-hocus-pocus",
	}

	var v struct {
		Response struct {
			Beer Beer `json:"beer"`
		} `json:"response"`
	}

	body, err := ioutil.ReadAll(got.Body)
	beer := &v
	if err := json.Unmarshal(body, beer); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want.BID, beer.Response.Beer.BID)
	assert.Equal(t, want.BeerName, beer.Response.Beer.BeerName)
	assert.Equal(t, want.BeerSlug, beer.Response.Beer.BeerSlug)

}
