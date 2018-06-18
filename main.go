package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

var conf = AppConfiguration{}
var err = gonfig.GetConf(getConfFile(), &conf)
var rest = sling.New().Set("Content-Type", "application/json").Base(conf.BaseURL)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	logrus.SetFormatter(&logrus.JSONFormatter{})
	fmt.Printf("%+v\n", conf)

	r.Route("/rest", func(r chi.Router) {
		r.Route("/auth/1", func(r chi.Router) {
			r.Route("/session", func(r chi.Router) {
				r.Post("/", loginHandler)
				r.Get("/", getCurrentLoginHandler)
			})
		})
		r.Route("/api/2", func(r chi.Router) {

		})
	})

	http.ListenAndServe(":8470", r)
}
