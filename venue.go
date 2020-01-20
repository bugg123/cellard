package main

type VenueService struct {
	client *Client
}

type Venue struct {
	VenueID   int      `json:"venue_id"`
	VenueName string   `json:"venue_name"`
	Location  Location `json:"location"`
}
