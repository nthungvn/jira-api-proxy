package main

import (
	"net/http"

	"github.com/google/uuid"

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
		if u.Remember {
			if err := saveLogin(w, r, u); err != nil {
				return nil, err
			}
		}
		sessionID := &http.Cookie{
			Name:  auth.Session.Name,
			Value: auth.Session.Value,
			Path:  "/",
		}
		http.SetCookie(w, sessionID)
		return nil, render.Render(w, r, auth)
	}
	return res, err
}

func saveLogin(w http.ResponseWriter, r *http.Request, u *User) error {
	if session, err := sessionStore.Get(r, JREMEMBERME); err == nil {
		if uuid, err := uuid.NewRandom(); err == nil {
			session.Values[uuid.String()] = u
			remeberCookie := &http.Cookie{
				Name:  JREMEMBERME,
				Value: uuid.String(),
			}
			http.SetCookie(w, remeberCookie)
			return session.Save(r, w)
		}
	}
	return nil
}

func getCurrentLoginHandler(rester *sling.Sling, w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	currentLogin := &CurrentSession{}
	res, err := rester.New().Get(currentUserAPI.URI).ReceiveSuccess(currentLogin)
	if err == nil && res.StatusCode == http.StatusOK {
		return nil, render.Render(w, r, currentLogin)
	}
	return res, err
}
