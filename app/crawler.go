package app

import (
	"encoding/json"
	"fmt"
	"github.com/fake-finder/fakefinder/external"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

func RequestCrawler(text string) (string, error) {
	client := external.NewClient()

	var bot external.CrawlerRequest
	bot.Description = strings.ToLower(text)

    fmt.Println("description",bot.Description)
	logrus.WithFields(logrus.Fields{
	}).Info("Do request.....")

	resp, err := client.DoRequest(bot)
	if err != nil {
		return "", err
	}

	logrus.WithFields(logrus.Fields{
	}).Info("request completed.....")

	var news external.CrawlerResponse
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))


	err = json.NewDecoder(resp.Body).Decode(&news)
	if err != nil {
		fmt.Println("convert err ", err)
		return "", err
	}
	var textResult string

	textResult = news.Description
	for _, i := range news.Text {
		textResult += i.Title
		textResult += i.Link.Encode()
		textResult += i.Similarity
		textResult += i.Date
		textResult += "\n"
	}
	fmt.Println("result:",textResult)


	return textResult, nil

}
