package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type webHookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

var emptyString = ""

func ReceiveMessage() error {
	print("Lysten and serv")
	err := http.ListenAndServe(":8080", http.HandlerFunc(webHookHandler))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// This function is called whenever an update is recieved
func webHookHandler(rw http.ResponseWriter, req *http.Request) {

	// Create our web hook request body type instance
	body := &webHookReqBody{}
	print("webhook req body", body)
	// Decodes the incoming request into our cutom webhookreqbody type
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Printf("An error occured (webHookHandler)")
		log.Panic(err)
		return
	}

	// If the command /joke is recieved call the sendReply function
	text := strings.ToLower(body.Message.Text)
	switch {
	case strings.HasPrefix(text, "/start"):
		text := "Olá, Para consultar a veracidade de uma notícia digite \n /consultar: descrição da notícia"
		err := sendReply(body.Message.Chat.ID, text)
		if err != nil {
			println(err)
			return
		}
	case strings.HasPrefix(text, "/consultar"):
		text := strings.TrimPrefix(strings.ToLower(body.Message.Text), "/consultar")
		if text == emptyString {
			text = "para consultar a veracidade de uma notícia digite\n/consultar: descrição da notícia"
		}
		err := sendReply(body.Message.Chat.ID, text)
		if err != nil {
			println(err)
			return
		}
	default:
		text := "Não entendi sua mensagem,\npara consultar a veracidade de uma notícia digite\n/consultar: descrição da notícia"
		err := sendReply(body.Message.Chat.ID, text)
		if err != nil {
			println(err)
			return
		}
	}
}

func sendReply(chatID int64, text string) error {
	fmt.Println("sendReply called")

	// calls the joke fetcher fucntion and gets a random joke from the API
	// for text != "pare" {

	// 	err := http.ListenAndServe(":8080", http.HandlerFunc(webHookHandler))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return err
	// 	}
	// 	text = "pare"

	// }
	if text == "mateus" {
		time.Sleep(20 * time.Second)
	}
	//Creates an instance of our custom sendMessageReqBody Type
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   text,
	}

	// Convert our custom type into json format
	reqBytes, err := json.Marshal(reqBody)

	if err != nil {
		return err
	}

	// Make a request to send our message using the POST method to the telegram bot API
	resp, err := http.Post(
		"https://api.telegram.org/bot"+os.Getenv("TELEGRAM_BOT_KEY")+"/"+"sendMessage",
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
