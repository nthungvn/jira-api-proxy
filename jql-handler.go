package main

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/render"

	"github.com/oceanicdev/chi-param"
)

func getFieldAutoCompleteDataHandler(w http.ResponseWriter, r *http.Request) {
	fieldName, _ := param.QueryString(r, "fieldName")
	fieldValue, _ := param.QueryString(r, "fieldValue")
	fieldSuggestion := &FieldSuggestion{
		FieldName:  fieldName,
		FieldValue: fieldValue,
	}

	fieldSuggestionResults := &FieldSuggestionResults{}
	res, err := rest.New().Get(autoCompleteData.URI).QueryStruct(fieldSuggestion).ReceiveSuccess(fieldSuggestionResults)
	if err == nil {
		logrus.Info(res)
		render.Render(w, r, fieldSuggestionResults)
		return
	}
}
