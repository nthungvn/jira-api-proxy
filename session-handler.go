package main

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/sirupsen/logrus"
)

var auth Authentication

func loginHandler(w http.ResponseWriter, r *http.Request) {
	u := &User{}
	if err := render.Bind(r, u); err != nil {
		render.Render(w, r, ErrorUnauthorized(err))
		return
	}
	_, err := rest.New().Post(loginAPI.URI).BodyJSON(u).ReceiveSuccess(&auth)
	if err == nil {
		logrus.Info(auth)
		render.Status(r, http.StatusOK)
		render.Render(w, r, &auth)
	}
}

func getCurrentLoginHandler(w http.ResponseWriter, r *http.Request) {
	var currentLogin CurrentSession
	_, err = rest.New().Get(currentUserAPI.URI).Set(COOKIE, auth.cookie()).ReceiveSuccess(&currentLogin)
	if err == nil {
		logrus.Info(currentLogin)
		render.Render(w, r, &currentLogin)
		return
	}
}
