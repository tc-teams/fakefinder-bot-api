package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/fake-finder/fakefinder/crawler"
)

var (
	start     = string("/start")
	consultar = string("/consultar")
	botToken  = "1228162506:AAFQr_ipJ3dsaOieJFupA5Pw4BlhuRFoOyE"
)

func TelegramWebHookHandler(w http.ResponseWriter, r *http.Request) {

	body := &webHookReqBody{}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		return
	}

	text := strings.ToLower(body.Message.Text)
	switch {
	case strings.HasPrefix(text, start):
		text := startCommand
		err := TelegramReply(body.Message.Chat.ID, text)
		if err != nil {
			return
		}
	case strings.HasPrefix(text, consultar):
		text := strings.TrimPrefix(strings.ToLower(body.Message.Text), consultar)
		if text == emptyString {
			text = consultCommand
			err := TelegramReply(body.Message.Chat.ID, text)
			if err != nil {
				return
			}
			break

		}
		result, err := crawler.RequestCrawler(text)
		if err != nil {
			return
		}

		err = TelegramReply(body.Message.Chat.ID, result)
		if err != nil {
			return
		}
	default:
		text := defaltCommand
		err := TelegramReply(body.Message.Chat.ID, text)
		if err != nil {
			return
		}
	}
}

func TelegramReply(chatID int64, text string) error {

	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   text,
	}
	reqBytes, err := json.Marshal(reqBody)

	if err != nil {
		return err
	}

	resp, err := http.Post(
		"https://api.telegram.org/bot"+botToken+"/"+"sendMessage",
		"application/json",
		bytes.NewBuffer(reqBytes),
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + resp.Status)
	}

	return err
}
