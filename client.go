package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	BaseURL      *url.URL
	UserAgent    string
	ClientID     string
	ClientSecret string

	httpClient *http.Client

	Beer    *BeerService
	Brewery *BreweryService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{httpClient: httpClient}
	c.BaseURL, _ = url.Parse("https://api.untappd.com")
	c.ClientID = os.Getenv("CLIENT_ID")
	c.ClientSecret = os.Getenv("CLIENT_SECRET")

	c.Beer = &BeerService{client: c}
	c.Brewery = &BreweryService{client: c}
	return c
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	q := u.Query()
	q.Add("client_id", c.ClientID)
	q.Add("client_secret", c.ClientSecret)
	u.RawQuery = q.Encode()

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
