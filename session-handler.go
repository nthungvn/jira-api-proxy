package main

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/sirupsen/logrus"
)

var auth Authentication

func loginHandler(w http.ResponseWriter, r *http.Request) error {
	u := &User{}
	if err := render.Bind(r, u); err != nil {
		return err
	}
	res, err := rest.New().Post(loginAPI.URI).BodyJSON(u).ReceiveSuccess(&auth)
	if err != nil {
		logrus.Info(res)
		return err
	}
	return render.Render(w, r, &auth)
}

func getCurrentLoginHandler(w http.ResponseWriter, r *http.Request) error {
	currentLogin := &CurrentSession{}
	res, err := rest.New().Get(currentUserAPI.URI).Set(COOKIE, auth.cookie()).ReceiveSuccess(currentLogin)
	if err != nil {
		logrus.Info(res)
		return err
	}
	return render.Render(w, r, currentLogin)
}
