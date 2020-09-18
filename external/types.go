package external

type CrawlerRequest struct {
	Description string `json:"description"`
}

type CrawlerResponse struct {
	Description string       `json:"description"`
	Text        []TextResult `json:"text,omitempty"`
}
type TextResult struct {
	Date       string `json:"date"`
	Title      string `json:"title"`
	Similarity string `json:"similarity,omitempty"`
	Link       string `json:"link"`
}
