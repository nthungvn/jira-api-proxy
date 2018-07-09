package main

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/oceanicdev/chi-param"
)

func searchGetIssueHandler(w http.ResponseWriter, r *http.Request) (*http.Response, error) {
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
		return res, err
	}
	return res, render.Render(w, r, responseAPI)
}

func searchPostIssueHandler(w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	requestAPI := &SearchIssuePostRequest{}
	if err := render.Bind(r, requestAPI); err != nil {
		return nil, err
	}
	responseAPI := &SearchIssueReponse{}
	res, err := rest.New().Post(searchPost.URI).BodyJSON(requestAPI).Set(COOKIE, auth.cookie()).ReceiveSuccess(responseAPI)
	if err != nil {
		return nil, err
	}
	return res, render.Render(w, r, responseAPI)
}
