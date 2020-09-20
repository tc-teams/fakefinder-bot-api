package handlers

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"

	"github.com/fake-finder/fakefinder/server"
	mux "github.com/gorilla/mux"
)

type Router struct {
	Mux *mux.Router
	Srv *server.ServerTls
	Ctx context.Context
}

//AddRoute, references a new handlers
func (r *Router) AddRoute(c *Context) {
	r.Mux.HandleFunc(c.Path, c.Handler)

}

func (r *Router) Server() error {

	context, cancel := context.WithCancel(r.Ctx)
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		logrus.Info("signal caught. shutting down...")
		cancel()
		r.Srv.Shutdown(context)

	}()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		r.Srv.Start(r.Ctx, r.Mux)
	}()

	wg.Wait()

	return nil
}

//NewRouter create a new instancie of http client
func NewRouter() *Router {
	return &Router{
		Mux: mux.NewRouter().StrictSlash(true),
		Srv: server.NewServerTLS(),
		Ctx: context.Background(),

	}

}
