package main

import (
	"net/http"

	"github.com/go-chi/render"
)

var sprintPicker = APIDeclaration{
	Method: GET,
	URI:    "rest/greenhopper/1.0/sprint/picker",
}

// SprintPickerRequest ...
type SprintPickerRequest struct {
	ExcludeCompleted bool   `url:"excludeCompleted,omitempty"`
	Query            string `url:"query,omitempty"`
}

// SprintPickerResponse ...
type SprintPickerResponse struct {
	Name      string `json:"name"`
	ID        int    `json:"id"`
	StateKey  string `json:"stateKey"`
	BoardName string `json:"boardName"`
	Date      string `json:"date"`
}

// SprintPickerSuggestion ...
type SprintPickerSuggestion struct {
	Suggestions []SprintPickerResponse `json:"suggestions"`
}

// Render ...
func (p *SprintPickerSuggestion) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}
