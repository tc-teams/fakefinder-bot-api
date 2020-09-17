package app

import (
	"github.com/fake-finder/fakefinder/bot"
	"github.com/fake-finder/fakefinder/external"
)

func RequestCrawler() {
	client := external.NewClient()

	var bot bot.BotRequest

	client.Request()
}
