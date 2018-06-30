package main

var autoCompleteData = APIDeclaration{
	Method: GET,
	URI:    "api/2/jql/autocompletedata/suggestions",
}

// Suggestion ...
type Suggestion struct {
	FieldName  string `url:"fieldName,omitempty"`
	FieldValue string `url:"fieldValue,omitempty"`
}

// SprintSuggestion ...
type SprintSuggestion struct {
	Value       string `json:"value"`
	DisplayName string `json:"displayName"`
}

// SprintSuggestions ...
type SprintSuggestions struct {
	Results []SprintSuggestion `json:"results"`
}
