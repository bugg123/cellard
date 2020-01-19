package main

import (
	"net/http"
	"strconv"
)

const beerInfoPath = "/v4/beer/info/"

func (b *BeerService) GetBeerInfo(beerID int) (*Beer, error) {
	req, err := b.client.newRequest(http.MethodGet, beerInfoPath+strconv.Itoa(beerID), nil)
	if err != nil {
		return nil, err
	}

	var v struct {
		Response struct {
			Beer Beer `json:"beer"`
		} `json:"response"`
	}
	_, err = b.client.do(req, &v)
	if err != nil {
		return nil, err
	}

	return &v.Response.Beer, nil
}
