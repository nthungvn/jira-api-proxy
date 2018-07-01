package main

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

var conf = AppConfiguration{}
var _ = gonfig.GetConf(getConfFile(), &conf)
var rest = sling.New().Set("Content-Type", "application/json").Base(conf.BaseURL)

func main() {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info(conf)

	r.Route("/rest", func(r chi.Router) {
		r.Route("/auth/1", func(r chi.Router) {
			r.Route("/session", func(r chi.Router) {
				r.Post("/", loginHandler)
				r.Get("/", getCurrentLoginHandler)
			})
		})
		r.Route("/api/2", func(r chi.Router) {
			r.Route("/jql", func(r chi.Router) {
				r.Route("/autocompletedata", func(r chi.Router) {
					r.Get("/suggestions", getFieldAutoCompleteDataHandler)
				})
			})
			r.Route("/search", func(r chi.Router) {
				r.Get("/", searchGetIssueHandler)
				r.Post("/", searchPostIssueHandler)
			})
		})
		r.Route("/greenhopper/1.0", func(r chi.Router) {
			r.Route("/sprint", func(r chi.Router) {
				r.Get("/picker", sprintPickerHanlder)
			})
		})
	})

	http.ListenAndServe(":8470", r)
}
