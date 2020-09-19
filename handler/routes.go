package handlers

import (
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type Context struct {
	Name    string
	Method  string
	Path    string
	Handler Handler
}
