package main

import (
	"net/http"

	"github.com/dghubble/sling"
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

// Cookie names
const (
	JSESSIONID = "JSESSIONID"
)

// APIDeclaration ...
type APIDeclaration struct {
	Method string
	URI    string
}

// Handler ...
type Handler func(rester *sling.Sling, w http.ResponseWriter, r *http.Request) (*http.Response, error)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, _ := r.Cookie(JSESSIONID)
	username, password, _ := r.BasicAuth()
	rester := rest.New().Set(COOKIE, session.String())

	if len(username) > 0 && len(password) > 0 {
		rester.SetBasicAuth(username, password)
	}

	if res, err := h(rester, w, r); err != nil {
		logrus.Info(res)
		render.Render(w, r, ErrInvalidRequest(err))
	}
}
