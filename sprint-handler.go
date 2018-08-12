package main

import (
	"net/http"

	"github.com/dghubble/sling"
	"github.com/go-chi/render"
	"github.com/oceanicdev/chi-param"
)

func sprintPickerHanlder(rester *sling.Sling, w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	excludeCompleted, _ := param.QueryBool(r, "excludeCompleted")
	query, _ := param.QueryString(r, "query")
	params := &SprintPickerRequest{
		ExcludeCompleted: excludeCompleted,
		Query:            query,
	}
	var sprintPickerSuggestion SprintPickerSuggestion

	res, err := rester.New().Get(sprintPicker.URI).QueryStruct(params).ReceiveSuccess(&sprintPickerSuggestion)
	if err == nil && res.StatusCode == http.StatusOK {
		return nil, render.Render(w, r, &sprintPickerSuggestion)
	}
	return res, err
}
