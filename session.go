package main

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

var loginAPI = APIDeclaration{
	Method: POST,
	URI:    "rest/auth/1/session",
}

var currentUserAPI = APIDeclaration{
	Method: GET,
	URI:    "rest/auth/1/session",
}

// Session ...
type Session struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// LoginInfo ...
type LoginInfo struct {
	FailedLoginCount    int    `json:"failedLoginCount"`
	LoginCount          int    `json:"loginCount"`
	LastFailedLoginTime string `json:"lastFailedLoginTime"`
	PreviousLoginTime   string `json:"previousLoginTime"`
}

// Authentication ...
type Authentication struct {
	Session   Session   `json:"session"`
	LoginInfo LoginInfo `json:"loginInfo"`
}

// Render ...
func (a *Authentication) Render(w http.ResponseWriter, r *http.Request) error {
	logrus.Info(a)
	render.Status(r, http.StatusOK)
	return nil
}

// CurrentSession ...
type CurrentSession struct {
	Name      string    `json:"name"`
	LoginInfo LoginInfo `json:"loginInfo"`
}

// Render ...
func (s *CurrentSession) Render(w http.ResponseWriter, r *http.Request) error {
	logrus.Info(s)
	render.Status(r, http.StatusOK)
	return nil
}

// User ...
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Bind interface for managing request payloads.
func (u *User) Bind(r *http.Request) error {
	if len(u.Username) == 0 || len(u.Password) == 0 {
		return errors.New("The username or password is invalid")
	}
	return nil
}
