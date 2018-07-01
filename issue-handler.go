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

	requestAPI := &SearchIssueRequest{
		Jql:        jql,
		Fields:     fields,
		MaxResults: maxResults,
	}
	reponseAPI := &SearchIssueReponse{}
	res, err := rest.New().Get(searchGet.URI).QueryStruct(requestAPI).Set(COOKIE, auth.cookie()).ReceiveSuccess(reponseAPI)
	if err == nil {
		render.Render(w, r, reponseAPI)
		return
	}
	logrus.Info(res)
}

func searchPostIssueHandler(w http.ResponseWriter, r *http.Request) {

}
