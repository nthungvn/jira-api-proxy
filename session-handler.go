package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

var auth Authentication

func loginHandler(w http.ResponseWriter, r *http.Request) {
	u := User{
		Username: "nthung",
		Password: "R50kZLs@6",
	}
	_, err := rest.New().Post(loginAPI.URI).BodyJSON(u).ReceiveSuccess(&auth)
	if err == nil {
		logrus.Info(auth)
	}

}

func getCurrentLoginHandler(w http.ResponseWriter, r *http.Request) {
	var currentLogin CurrentSession
	_, err = rest.New().Get(currentUserAPI.URI).Set(COOKIE, auth.cookie()).ReceiveSuccess(&currentLogin)
	if err == nil {
		fmt.Printf("%+v\n", currentLogin)
	}
	json.NewEncoder(w).Encode(currentLogin)
}
