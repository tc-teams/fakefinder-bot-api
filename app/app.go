package app

import (
	"net/http"

	"github.com/fake-finder/fakefinder/app/bot"
	"github.com/fake-finder/fakefinder/handlers"
)

func Run() *handlers.Router {
	mux := handlers.NewRouter()
	mux.AddRoute(&handlers.Context{
		Name:    "bot api ",
		Method:  http.MethodPost,
		Path:    "/telegram/bot/api",
		Handler: bot.TelegramWebHookHandler,
	})
	return mux
}
