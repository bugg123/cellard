package main

import (
	"bytes"
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
	AddBeerSearchQuery(gotURL, "Bourbon County Brand Stout")

	got, err := MakeQuery(gotURL)
	assert.NoError(t, err)

	body, err := ioutil.ReadAll(got.Body)
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, body, "", " "); err != nil {
		t.Fatal(err)
	}
	t.Log(dst)

}
