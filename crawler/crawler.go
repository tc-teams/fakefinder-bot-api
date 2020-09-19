package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fake-finder/fakefinder/external"
	"github.com/sirupsen/logrus"
	"strings"
)

var empty = string("")

//RequestCrawler
func RequestCrawler(text string) (string, error) {
	client := external.NewClient()

	var bot external.CrawlerRequest
	bot.Description = strings.ToLower(text)

	fmt.Println("description", bot.Description)
	logrus.WithFields(logrus.Fields{}).Info("Do request.....")

	resp, err := client.DoRequest(bot)
	if err != nil {
		return empty, err
	}

	logrus.WithFields(logrus.Fields{}).Info("request completed.....")

	var news external.CrawlerResponse

	err = json.NewDecoder(resp.Body).Decode(&news)
	if err != nil {
		fmt.Println("convert err ", err)
		return empty, err
	}

	return CrawlerMakeResponse(news)

}

func CrawlerMakeResponse(news external.CrawlerResponse) (string, error) {

	var textResult string

	textResult = news.Description
	for _, i := range news.Text {
		textResult += i.Title
		textResult += i.Link
		textResult += i.Similarity
		textResult += i.Date
		textResult += "\n"
	}

	if textResult == empty {
		errEmpty := errors.New("string not found")
		return empty, errEmpty

	}
	logrus.WithFields(logrus.Fields{
		"textResult": textResult,
	}).Info("result")

	return textResult, nil

}
