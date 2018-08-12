package main

import (
	"net/http"

	"github.com/dghubble/sling"
	"github.com/go-chi/render"
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

	res, err := h(rester, w, r)

	if res == nil && err == nil {
		// The handler already handle
	} else if res != nil && err == nil {
		h.handleErrorCode(w, res, err)
	} else {
		render.Render(w, r, ErrServerError(err))
	}
}

func (h Handler) handleErrorCode(w http.ResponseWriter, res *http.Response, err error) {
	if res.StatusCode == http.StatusUnauthorized {
		render.Render(w, res.Request, ErrUnauthorized(err))
	} else {
		render.Render(w, res.Request, &ResponseError{
			HTTPStatusCode: res.StatusCode,
		})
	}
}
