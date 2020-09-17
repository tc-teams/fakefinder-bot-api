package bot

import (
	"errors"
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

	updates, err := t.Bot.GetUpdatesChan(t.Help)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ErrorUpdate": err,
		}).Error("error to get updates")
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		logrus.WithFields(logrus.Fields{
			"UserID":    update.Message.From.ID,
			"UserName":  update.Message.From.UserName,
			"FirstName": update.Message.From.FirstName,
			"LastName":  update.Message.From.LastName,
			"Text":      update.Message.Text,
		}).Info("received message")
		t.Msg.ChatID = update.Message.Chat.ID
		t.Msg.Text = strings.ToUpper(update.Message.Text)

		t.Listener(updates)

	}
	return nil
}

func (t *Telegram) Listener(updates tgbotapi.UpdatesChannel) error {
	switch t.Msg.Text {
	case "OI", "OLA", "OLÁ", "OPA", "SALVE", "/START":
		t.Msg.Text = "Olá, Para consultar a veracidade de uma notícia digite Consultar ou aperte aqui \n /consultar"
	case "/CONSULTAR","CONSULTAR":
		t.Msg.Text = "Por favor, descreva a notícia em poucas palavras descrevendo os principais pontos."
		t.Bot.Send(t.Msg)
		for update := range updates {
			if update.Message == nil {
				continue
			}

			t.Msg.ChatID = update.Message.Chat.ID
			t.Msg.Text = strings.ToUpper(update.Message.Text)
			break
		}
		var text string
		text = t.Msg.Text
		t.Msg.Text = "Processando requisição"
		t.Bot.Send(t.Msg)

		resp, err := app.RequestCrawler(text)
		if err != nil {
			t.Msg.Text = "erro ao processar a requisição"
			return err
		}
		t.Msg.Text = resp
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
		keyNotFound := errors.New("key not found ")
		return nil, keyNotFound
	}

	logrus.WithFields(logrus.Fields{
		"UserName": Bot.Self.UserName,
	}).Info("Authorized on account")
	//Bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	return &Telegram{
		Bot:  Bot,
		Help: updateConfig,
		Msg:  tgbotapi.MessageConfig{},
	}, nil

}
