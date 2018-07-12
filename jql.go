package main

import (
	"net/http"

	"github.com/go-chi/render"
)

var autoCompleteData = APIDeclaration{
	Method: GET,
	URI:    "rest/api/2/jql/autocompletedata/suggestions",
}

// FieldSuggestion ...
type FieldSuggestion struct {
	FieldName  string `url:"fieldName,omitempty"`
	FieldValue string `url:"fieldValue,omitempty"`
}

// FieldSuggestionResult ...
type FieldSuggestionResult struct {
	Value       string `json:"value"`
	DisplayName string `json:"displayName"`
}

// FieldSuggestionResults ...
type FieldSuggestionResults struct {
	Results []FieldSuggestionResult `json:"results"`
}

// Render ...
func (f *FieldSuggestionResults) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}
