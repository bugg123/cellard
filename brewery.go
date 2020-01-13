package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	breweryInfoPath = "/v4/brewery/info/"
)

type Brewery struct {
	BreweryID           float64 `json:"brewery_id"`
	BreweryName         string  `json:"brewery_name"`
	BrewerySlug         string  `json:"brewery_slug"`
	BreweryLabel        string  `json:"brewery_label"`
	CountryName         string  `json:"country_name"`
	BreweryInProduction string  `json:"brewery_in_production"`
	IsIndependent       string  `json:"is_independent"`
	ClaimedStatus       struct {
		IsClaimed    bool    `json:"is_claimed"`
		ClaimedSlug  string  `json:"claimed_slug"`
		FollowStatus bool    `json:"follow_status"`
		UID          float64 `json:"uid"`
		MuteStatus   string  `json:"mute_status"`
	} `json:"claimed_status"`
	BeerCount float64 `json:"beer_count"`
	Contact   struct {
		Twitter   string `json:"twitter"`
		Facebook  string `json:"facebook"`
		Instagram string `json:"instagram"`
		URL       string `json:"url"`
	} `json:"contact"`
	BreweryType   string  `json:"brewery_type"`
	BreweryTypeID float64 `json:"brewery_type_id"`
	Location      struct {
		BreweryAddress string  `json:"brewery_address"`
		BreweryCity    string  `json:"brewery_city"`
		BreweryState   string  `json:"brewery_state"`
		BreweryLAT     float64 `json:"brewery_lat"`
		BreweryLNG     float64 `json:"brewery_lng"`
	} `json:"location"`
}

func GetBreweryInfo(breweryID int) Brewery {
	givenURL := GetUntappedHost()
	SetClientAuthString(givenURL)
	givenURL.Path = fmt.Sprintf(breweryInfoPath+"%d", breweryID)
	resp, _ := http.Get(givenURL.String())
	body, _ := ioutil.ReadAll(resp.Body)

	var v struct {
		Response struct {
			Brewery Brewery `json:"brewery"`
		} `json:"response"`
	}
	brewery := &v
	json.Unmarshal(body, brewery)

	return brewery.Response.Brewery
}
