package handlers

import (
	"context"

	"github.com/fake-finder/fakefinder/server"
	mux "github.com/gorilla/mux"
)

type Router struct {
	Mux *mux.Router
	Srv server.Client
	Ctx context.Context
}

//AddRoute, references a new handler
func (r *Router) AddRoute(c *Context) {
	r.Mux.HandleFunc(c.Path, c.Handler)

}

func (r *Router) Server() {
	r.Srv.Start(r.Ctx, r.Mux)
}

//NewRouterInstance created a new instance of the router
func NewRouterInstance() *Router {
	return &Router{
		Mux: mux.NewRouter().StrictSlash(true),
	}

}
