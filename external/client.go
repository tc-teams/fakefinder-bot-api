package external

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	*http.Client
}

func (c *Client) DoRequest(r CrawlerRequest) (*http.Response,error) {

	reqBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
        "http://104.197.246.105:8080/search/news",
		bytes.NewBuffer(reqBytes),
	)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json; charset=utf-8")

	response, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	return response,nil
}

//NewClient return a new client instance
func NewClient() *Client {
	return &Client{
		&http.Client{
			Timeout: time.Minute,
		},
	}
}
