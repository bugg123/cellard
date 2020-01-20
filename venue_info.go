package main

import (
	"net/http"
	"strconv"
)

const venueInfoPath = "/v4/venue/info/"

func (v *VenueService) GetVenueInfo(venueID int) (*Venue, *http.Response, error) {
	url := venueInfoPath + strconv.Itoa(venueID)
	req, err := v.client.newRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	var t struct {
		Response struct {
			Venue Venue `json:"venue"`
		} `json:"response"`
	}
	resp, err := v.client.do(req, &t)
	if err != nil {
		return nil, resp, err
	}
	return &t.Response.Venue, resp, err
}
