package external

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type Client struct {
	*http.Client
}

func (c *Client) DoRequest(r CrawlerRequest) (*http.Response, error) {

	reqBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBytes).Encode(r)
	if err != nil {

	}

	request, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("CRAWLER_URL"),
		reqBytes,
	)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json; charset=utf-8")

	response, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//NewClient return a new client instance
func NewClient() *Client {
	return &Client{
		&http.Client{
			Timeout: time.Minute,
		},
	}
}
