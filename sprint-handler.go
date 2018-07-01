package main

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/render"
	"github.com/oceanicdev/chi-param"
)

func sprintPickerHanlder(w http.ResponseWriter, r *http.Request) {
	excludeCompleted, _ := param.QueryBool(r, "excludeCompleted")
	query, _ := param.QueryString(r, "query")
	params := &SprintPickerRequest{
		ExcludeCompleted: excludeCompleted,
		Query:            query,
	}
	var sprintPickerSuggestion SprintPickerSuggestion

	res, err := rest.New().Get(sprintPicker.URI).QueryStruct(params).Set(COOKIE, auth.cookie()).ReceiveSuccess(&sprintPickerSuggestion)
	if err == nil {
		logrus.Info(res)
		render.Render(w, r, &sprintPickerSuggestion)
		return
	}
}
