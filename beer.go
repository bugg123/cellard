package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const beerSearchPath = "v4/search/beer"
const beerInfoPath = "v4/beer/info/"

type Beer struct {
	BID             float64 `json:"bid"`
	BeerName        string  `json:"beer_name"`
	BeerLabel       string  `json:"beer_label"`
	BeerAbv         float64 `json:"beer_abv"`
	BeerSlug        string  `json:"beer_slug"`
	BeerIbu         float64 `json:"beer_ibu"`
	BeerDescription string  `json:"beer_description"`
	CreatedAt       string  `json:"created_at"`
	BeerStyle       string  `json:"beer_style"`
	InProduction    float64 `json:"in_production"`
	AuthRating      float64 `json:"auth_rating"`
	WishList        bool    `json:"wish_list"`
}

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

func GetBeerInfo(givenURL *url.URL, beerID int) {
	givenURL.Path = fmt.Sprintf(beerInfoPath+"%d", beerID)

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
