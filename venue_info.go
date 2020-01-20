package main

import (
	"net/http"
	"strconv"
)

const venueInfoPath = "/v4/venue/info/"

func (v *VenueService) GetVenueInfo(venueID int) (*Venue, error) {
	url := venueInfoPath + strconv.Itoa(venueID)
	req, err := v.client.newRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var t struct {
		Response struct {
			Venue Venue `json:"venue"`
		} `json:"response"`
	}
	_, err = v.client.do(req, &t)
	if err != nil {
		return nil, err
	}
	return &t.Response.Venue, err
}
