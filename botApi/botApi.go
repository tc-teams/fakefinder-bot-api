package botApi

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"os"
)

type BotApi struct {
	Bot  *tgbotapi.BotAPI
	Help tgbotapi.UpdateConfig
	msg  tgbotapi.MessageConfig
}

func (b *BotApi) ReceiveMessage() error {

	b.Bot.Debug = true

	logrus.WithFields(logrus.Fields{
		"UserName": b.Bot.Self.UserName,
	}).Info("Authorized on account")

	b.Help.Timeout = 60

	updates, err := b.Bot.GetUpdatesChan(b.Help)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ErrorUpdate": err,
		}).Error("error to get updates")
		return err
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		logrus.WithFields(logrus.Fields{
			"UserID":    update.Message.From.ID,
			"UserName":  update.Message.From.UserName,
			"FirstName": update.Message.From.FirstName,
			"LastName":  update.Message.From.LastName,
			"Text":      update.Message.Text,
		}).Info("received message")

		b.msg.ChatID = update.Message.Chat.ID
		b.msg.Text = update.Message.Text
		b.msg.ReplyToMessageID = update.Message.MessageID

		b.Bot.Send(b.msg)
	}

	return nil
}

func NewBot() (*BotApi, error) {
	Bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_KEY"))
	if err != nil {
		fmt.Println("key not found")
		return nil, err
	}
	return &BotApi{
		Bot:  Bot,
		Help: tgbotapi.NewUpdate(0),
		msg:  tgbotapi.MessageConfig{},
	}, nil

}
