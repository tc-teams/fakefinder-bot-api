package bot

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"os"
)

type Telegram struct {
	Bot  *tgbotapi.BotAPI
	Help tgbotapi.UpdateConfig
	msg  tgbotapi.MessageConfig
}

func (t *Telegram) ReceiveMessage() error {

	t.Bot.Debug = true

	logrus.WithFields(logrus.Fields{
		"UserName": t.Bot.Self.UserName,
	}).Info("Authorized on account")

	t.Help.Timeout = 60

	updates, err := t.Bot.GetUpdatesChan(t.Help)
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

		t.msg.ChatID = update.Message.Chat.ID
		t.msg.Text = fmt.Sprintf("Olá %s.",update.Message.From.FirstName)
		t.msg.ReplyToMessageID = update.Message.MessageID

		t.Bot.Send(t.msg)
	}

	return nil
}

func NewBot() (*Telegram, error) {
	Bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_KEY"))
	if err != nil {
		fmt.Println("key not found")
		return nil, err
	}
	return &Telegram{
		Bot:  Bot,
		Help: tgbotapi.NewUpdate(0),
		msg:  tgbotapi.MessageConfig{},
	}, nil

}
