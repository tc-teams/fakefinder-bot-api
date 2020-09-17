package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fake-finder/fakefinder/bot"
)

type Client struct {
	*http.Client
}

func (c *Client) Request(r bot.BotRequest) error {

	reqBytes, err := json.Marshal(r)
	if err != nil {
		print(err)
	}

	request, err := http.NewRequest(
		http.MethodPost,

		bytes.NewBuffer(reqBytes),
	)
	if err != nil {
		print(err)
	}

	request.Header.Set("Accept", "application/json; charset=utf-8")

	response, err := c.Client.Do(request)
	if err != nil {
		print(err)
	}
	fmt.Println(response)
	return nil
}

//NewClient return a new client instance
func NewClient() *Client {
	return &Client{
		&http.Client{
			Timeout: time.Minute,
		},
	}
}
