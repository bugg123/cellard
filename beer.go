package main

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
