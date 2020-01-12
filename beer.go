package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const beerSearchPath = "v4/search/beer"

func GetUntappedHost() *url.URL {
	apiURL, err := url.Parse("https://api.untappd.com")
	if err != nil {
		log.Fatal(err)
	}
	return apiURL
}

func GetClientAuthString() string {
	return fmt.Sprintf("?client_id=%s&client_secret=%s", os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
}

func SetClientAuthString(givenURL *url.URL) {
	q := givenURL.Query()
	q.Add("client_id", os.Getenv("CLIENT_ID"))
	q.Add("client_secret", os.Getenv("CLIENT_SECRET"))

	givenURL.RawQuery = q.Encode()
}

func AddBeerSearchQuery(givenURL *url.URL, search string) {
	givenURL.Path = beerSearchPath

	q := givenURL.Query()
	q.Add("q", search)
	givenURL.RawQuery = q.Encode()

}

func MakeQuery(givenURL *url.URL) (*http.Response, error) {
	resp, err := http.Get(givenURL.String())
	return resp, err
}
