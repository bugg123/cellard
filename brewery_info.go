package main

import (
	"net/http"
	"strconv"
)

const (
	breweryInfoPath = "/v4/brewery/info/"
)

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
