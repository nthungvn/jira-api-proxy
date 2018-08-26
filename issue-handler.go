package main

import (
	"net/http"

	"github.com/dghubble/sling"
	"github.com/go-chi/render"
	"github.com/oceanicdev/chi-param"
)

func searchGetIssueHandler(rester *sling.Sling, w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	jql, _ := param.QueryString(r, "jql")
	fields, _ := param.QueryString(r, "fields")
	maxResults, _ := param.QueryInt(r, "maxResults")

	requestAPI := &SearchIssueGetRequest{
		Jql:        jql,
		Fields:     fields,
		MaxResults: maxResults,
	}
	responseAPI := &SearchIssueReponse{}
	res, err := rester.New().Get(searchGet.URI).QueryStruct(requestAPI).ReceiveSuccess(responseAPI)
	if err == nil && res.StatusCode == http.StatusOK {
		return nil, render.Render(w, r, responseAPI)
	}
	return res, err
}

func searchPostIssueHandler(rester *sling.Sling, w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	requestAPI := &SearchIssuePostRequest{}
	if err := render.Bind(r, requestAPI); err != nil {
		return nil, err
	}
	responseAPI := &SearchIssueReponse{}
	res, err := rester.New().Post(searchPost.URI).BodyJSON(requestAPI).ReceiveSuccess(responseAPI)
	if err == nil && res.StatusCode == http.StatusOK {
		return nil, render.Render(w, r, responseAPI)
	}
	return res, err
}
