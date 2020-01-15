package main

import (
	"net/http"
	"strconv"
)

const (
	breweryInfoPath = "/v4/brewery/info/"
)

type BreweryService struct {
	client *Client
}

type Brewery struct {
	BreweryID           int           `json:"brewery_id"`
	BreweryName         string        `json:"brewery_name"`
	BrewerySlug         string        `json:"brewery_slug"`
	BreweryLabel        string        `json:"brewery_label"`
	CountryName         string        `json:"country_name"`
	BreweryInProduction int           `json:"brewery_in_production"`
	IsIndependent       int           `json:"is_independent"`
	ClaimedStatus       ClaimedStatus `json:"claimed_status"`
	BeerCount           int           `json:"beer_count"`
	Contact             Contact       `json:"contact"`
	BreweryType         string        `json:"brewery_type"`
	BreweryTypeID       int           `json:"brewery_type_id"`
	Location            Location      `json:"location"`
}

type ClaimedStatus struct {
	IsClaimed    bool   `json:"is_claimed"`
	ClaimedSlug  string `json:"claimed_slug"`
	FollowStatus bool   `json:"follow_status"`
	UID          int    `json:"uid"`
	MuteStatus   string `json:"mute_status"`
}

type Contact struct {
	Twitter   string `json:"twitter"`
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
	URL       string `json:"url"`
}

type Location struct {
	BreweryAddress string  `json:"brewery_address"`
	BreweryCity    string  `json:"brewery_city"`
	BreweryState   string  `json:"brewery_state"`
	BreweryLAT     float64 `json:"brewery_lat"`
	BreweryLNG     float64 `json:"brewery_lng"`
}

func (b *BreweryService) GetBreweryInfo(breweryID int) (*Brewery, error) {
	req, err := b.client.newRequest(http.MethodGet, breweryInfoPath+strconv.Itoa(breweryID), nil)
	if err != nil {
		return nil, err
	}

	var v struct {
		Response struct {
			Brewery Brewery `json:"brewery"`
		} `json:"response"`
	}
	_, err = b.client.do(req, &v)
	if err != nil {
		return nil, err
	}

	return &v.Response.Brewery, nil
}
