package external

import (
	"net/url"
)


type CrawlerRequest struct {
	Description string `json:"description"`
}

type CrawlerResponse struct {
	Description string       `json:"description"`
	Text        []TextResult `json:"text,omitempty"`
}
type TextResult struct {
	Date       string    `json:"date"`
	Title      string    `json:"title"`
	Similarity string    `json:"similarity,omitempty"`
	Link       url.Values  `json:"link"`
}
