package main

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/oceanicdev/chi-param"
	"github.com/sirupsen/logrus"
)

func searchGetIssueHandler(w http.ResponseWriter, r *http.Request) {
	jql, _ := param.QueryString(r, "jql")
	fields, _ := param.QueryString(r, "fields")
	maxResults, _ := param.QueryInt(r, "maxResults")

	requestAPI := &SearchIssueGetRequest{
		Jql:        jql,
		Fields:     fields,
		MaxResults: maxResults,
	}
	responseAPI := &SearchIssueReponse{}
	res, err := rest.New().Get(searchGet.URI).QueryStruct(requestAPI).Set(COOKIE, auth.cookie()).ReceiveSuccess(responseAPI)
	if err != nil {
		logrus.Info(res)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Render(w, r, responseAPI)
}

func searchPostIssueHandler(w http.ResponseWriter, r *http.Request) {
	requestAPI := &SearchIssuePostRequest{}
	if err := render.Bind(r, requestAPI); err != nil {
		logrus.Error(err.Error())
	}
	responseAPI := &SearchIssueReponse{}
	res, err := rest.New().Post(searchPost.URI).BodyJSON(requestAPI).Set(COOKIE, auth.cookie()).ReceiveSuccess(responseAPI)
	if err != nil {
		logrus.Info(res)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Render(w, r, responseAPI)
}
