package bot

import (
	"fmt"
	"os"
	"strings"

	"github.com/fake-finder/fakefinder/app"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type Telegram struct {
	Bot  *tgbotapi.BotAPI
	Help tgbotapi.UpdateConfig
	Msg  tgbotapi.MessageConfig
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

		// logrus.WithFields(logrus.Fields{
		// 	"UserID":    update.Message.From.ID,
		// 	"UserName":  update.Message.From.UserName,
		// 	"FirstName": update.Message.From.FirstName,
		// 	"LastName":  update.Message.From.LastName,
		// 	"Text":      update.Message.Text,
		// }).Info("received message")
		t.Msg.ChatID = update.Message.Chat.ID
		t.Msg.Text = strings.ToUpper(update.Message.Text)

		t.Listener(updates)

	}
	return nil
}

func (t *Telegram) Listener(updates tgbotapi.UpdatesChannel) error {
	switch t.Msg.Text {
	case "OI", "OLA", "OLÁ", "OPA", "SALVE", "/START":
		t.Msg.Text = "Olá, Para consultar a veracidade de uma notícia digite Consultar ou aperte aqui /consultar"
	case "/CONSULTAR":
		t.Msg.Text = "Por favor, descreva a notícia em poucas palavras descrevendo os principais pontos."
		t.Bot.Send(t.Msg)
		for update := range updates {
			println("Entrou na poha do for")
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}

			t.Msg.ChatID = update.Message.Chat.ID
			t.Msg.Text = strings.ToUpper(update.Message.Text)
			break
		}

		app.RequestCrawler(t.Msg.Text)

	case "/OUTRO":
		t.Msg.Text = "I'm ok."
	default:
		t.Msg.Text = "I don't know that command"
	}
	t.Bot.Send(t.Msg)
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
		Msg:  tgbotapi.MessageConfig{},
	}, nil

}
