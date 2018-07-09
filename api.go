package main

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

// HTTP verbs
const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// APIDeclaration ...
type APIDeclaration struct {
	Method string
	URI    string
}

// Handler ...
type Handler func(w http.ResponseWriter, r *http.Request) (*http.Response, error)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if res, err := h(w, r); err != nil {
		logrus.Info(res)
		render.Render(w, r, ErrInvalidRequest(err))
	}
}
