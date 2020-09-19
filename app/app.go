package app

import (
	"net/http"

	"github.com/fake-finder/fakefinder/app/bot"
	"github.com/fake-finder/fakefinder/handler"
)

func Run() handler.Router {
	mux := handler.NewRouteInstance()
	mux.AddRoute(&handler.Context{
		Name:    "Pet",
		Method:  http.MethodGet,
		Path:    "/hello/pet/project",
		Handler: bot.ReceiveMessage(),
	})
	return mux
}
