package main

import (
	"net/http"
	"net/url"
	"strconv"
)

const beerSearchPath = "/v4/search/beer"
const beerInfoPath = "/v4/beer/info/"

type BeerService struct {
	client *Client
}

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

func (b *BeerService) SearchBeerQuery(query string, limit int) (*[]Beer, error) {

	req, err := b.client.newRequest(http.MethodGet, beerSearchPath, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("q", query)
	q.Add("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()

	var v struct {
		Response struct {
			Beers struct {
				Items []struct {
					Beer Beer `json:"beer"`
				} `json:"items"`
			} `json:"beers"`
		}
	}

	_, err = b.client.do(req, &v)
	if err != nil {
		return nil, err
	}
	var beers []Beer
	for _, item := range v.Response.Beers.Items {
		beers = append(beers, item.Beer)
	}

	return &beers, nil
}

func AddBeerSearchQuery(givenURL *url.URL, search string) {
	givenURL.Path = beerSearchPath

	q := givenURL.Query()
	q.Add("q", search)
	givenURL.RawQuery = q.Encode()

}
