package main

import (
	"net/http"

	"github.com/go-chi/render"
)

var auth Authentication

func loginHandler(w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	u := &User{}
	if err := render.Bind(r, u); err != nil {
		return nil, err
	}
	res, err := rest.New().Post(loginAPI.URI).BodyJSON(u).ReceiveSuccess(&auth)
	if err != nil {
		return res, err
	}
	return res, render.Render(w, r, &auth)
}

func getCurrentLoginHandler(w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	currentLogin := &CurrentSession{}
	res, err := rest.New().Get(currentUserAPI.URI).Set(COOKIE, auth.cookie()).ReceiveSuccess(currentLogin)
	if err != nil {
		return res, err
	}
	return res, render.Render(w, r, currentLogin)
}
