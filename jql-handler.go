package main

import (
	"net/http"

	"github.com/dghubble/sling"
	"github.com/go-chi/render"
	"github.com/oceanicdev/chi-param"
)

func getFieldAutoCompleteDataHandler(rester *sling.Sling, w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	fieldName, _ := param.QueryString(r, "fieldName")
	fieldValue, _ := param.QueryString(r, "fieldValue")
	fieldSuggestion := &FieldSuggestion{
		FieldName:  fieldName,
		FieldValue: fieldValue,
	}

	fieldSuggestionResults := &FieldSuggestionResults{}
	res, err := rester.New().Get(autoCompleteData.URI).QueryStruct(fieldSuggestion).ReceiveSuccess(fieldSuggestionResults)
	if err != nil {
		return res, err
	}
	return res, render.Render(w, r, fieldSuggestionResults)
}
