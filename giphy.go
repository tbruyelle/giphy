package giphy

import (
	"github.com/tbruyelle/apiclient"
	"net/http"
)

type Client struct {
	api *apiclient.API
	key string
}

func NewClient(APIKey string) *Client {
	c := &Client{
		api: apiclient.New("http://api.giphy.com/v1/"),
		key: APIKey,
	}
	return c
}

type Query struct {
	APIKey string `url:"api_key"`
}

type SearchQuery struct {
	Query
	Q string `url:"q"`
}

type SearchResponse struct {
	Data []Gif `json:"data"`
}

type Gif struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	URL    string `json:"url"`
	Images Images `json:"images"`
}

type Images struct {
	Origin Image `json:"original"`
}

type Image struct {
	URL  string `json:"url"`
	Mp4  string `json:"mp4"`
	Webp string `json:"webp"`
}

func (c *Client) Search(q string) (*SearchResponse, *http.Response, error) {
	s := &SearchQuery{Query: Query{APIKey: c.key}, Q: q}

	r, err := c.api.NewRequest("GET", "gifs/search", s, nil)
	if err != nil {
		return nil, nil, err
	}

	searchResp := new(SearchResponse)
	resp, err := c.api.Do(r, searchResp)
	if err != nil {
		return nil, resp, err
	}
	return searchResp, resp, nil
}
