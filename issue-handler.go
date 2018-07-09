package main

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/oceanicdev/chi-param"
	"github.com/sirupsen/logrus"
)

func searchGetIssueHandler(w http.ResponseWriter, r *http.Request) error {
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
		return err
	}
	return render.Render(w, r, responseAPI)
}

func searchPostIssueHandler(w http.ResponseWriter, r *http.Request) error {
	requestAPI := &SearchIssuePostRequest{}
	if err := render.Bind(r, requestAPI); err != nil {
		logrus.Error(err.Error())
		return err
	}
	responseAPI := &SearchIssueReponse{}
	res, err := rest.New().Post(searchPost.URI).BodyJSON(requestAPI).Set(COOKIE, auth.cookie()).ReceiveSuccess(responseAPI)
	if err != nil {
		logrus.Info(res)
		return err
	}
	return render.Render(w, r, responseAPI)
}
