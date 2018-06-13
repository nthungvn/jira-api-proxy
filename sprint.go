package main

var sprintPicker = APIDeclaration{
	Method: POST,
	URI:    "greenhopper/1.0/sprint/picker",
}

// SprintPickerRequest ...
type SprintPickerRequest struct {
	ExcludeCompleted bool   `json:"excludeCompleted"`
	Query            string `json:"query"`
}

// SprintPickerResponse ...
type SprintPickerResponse struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	StateKey  string `json:"stateKey"`
	BoardName string `json:"boardName"`
	Date      string `json:"date"`
}

// Sprint status
const (
	FUTURE = "FUTURE"
	ACTIVE = "ACTIVE"
	CLOSED = "CLOSED"
)
