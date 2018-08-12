package main

import (
	"net/http"

	"github.com/dghubble/sling"
	"github.com/go-chi/render"
)

func loginHandler(rester *sling.Sling, w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	auth := &Authentication{}
	u := &User{}
	if err := render.Bind(r, u); err != nil {
		return nil, err
	}
	res, err := rester.New().Post(loginAPI.URI).BodyJSON(u).ReceiveSuccess(auth)
	if err == nil && res.StatusCode == http.StatusOK {
		return nil, render.Render(w, r, auth)
	}
	return res, err
}

func getCurrentLoginHandler(rester *sling.Sling, w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	currentLogin := &CurrentSession{}
	res, err := rester.New().Get(currentUserAPI.URI).ReceiveSuccess(currentLogin)
	if err != nil {
		return res, err
	}
	return res, render.Render(w, r, currentLogin)
}
