package main

import (
	"fmt"

	"github.com/tkanos/gonfig"

	"github.com/dghubble/sling"
)

var conf = AppConfiguration{}
var err = gonfig.GetConf(getConfFile(), &conf)
var rest = sling.New().Set("Content-Type", "application/json").Base(conf.BaseURL)

func main() {
	fmt.Printf("%+v\n", conf)
	u := User{
		Username: "nthung",
		Password: "R50kZLs@6",
	}
	var auth Authentication
	_, err := rest.New().Post(loginAPI.URI).BodyJSON(u).ReceiveSuccess(&auth)
	if err == nil {
		fmt.Printf("%+v\n", auth)
	}

	var currentLogin CurrentSession
	_, err = rest.New().Get(currentUserAPI.URI).Set(COOKIE, auth.cookie()).ReceiveSuccess(&currentLogin)
	if err == nil {
		fmt.Printf("%+v\n", currentLogin)
	}
}
