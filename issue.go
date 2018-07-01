package main

import (
	"net/http"

	"github.com/go-chi/render"
)

var searchGet = APIDeclaration{
	Method: GET,
	URI:    "rest/api/2/search",
}

var searchPost = APIDeclaration{
	Method: POST,
	URI:    "rest/api/2/search",
}

// SimpleIssue ...
type SimpleIssue struct {
	ID     string      `json:"id"`
	Key    string      `json:"key"`
	Self   string      `json:"self"`
	Fields SimpleField `json:"fields"`
}

// SimpleField ...
type SimpleField struct {
	Summary   string    `json:"summary"`
	Status    Status    `json:"status"`
	Priority  Priority  `json:"priority"`
	IssueType IssueType `json:"issuetype"`
}

// Issue ...
type Issue struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Self   string `json:"self"`
	Fields Field  `json:"fields"`
}

// Field ...
type Field struct {
	IssueType IssueType   `json:"issuetype"`
	Parent    SimpleIssue `json:"parent"`
	Priority  Priority    `json:"priority"`
	Assignee  Person      `json:"assignee"`
	Status    Status      `json:"status"`
	Summary   string      `json:"summary"`
	SubTasks  []Issue     `json:"subtasks"`
	Reporter  Person      `json:"reporter"`
	Point     float32     `json:"customfield_10002"`
}

// Priority ...
type Priority struct {
	IconURL string `json:"iconUrl"`
	Name    string `json:"name"`
}

// Status ...
type Status struct {
	IconURL string `json:"iconUrl"`
	Name    string `json:"name"`
}

// Person ...
type Person struct {
	Name         string     `json:"name"`
	Key          string     `json:"key"`
	AvatarUrls   AvatarUrls `json:"avatarUrls"`
	EmailAddress string     `json:"emailAddress"`
	DisplayName  string     `json:"displayName"`
}

// AvatarUrls ...
type AvatarUrls struct {
	ExtraSmall string `json:"16x16"`
	Small      string `json:"24x24"`
	Medium     string `json:"32x32"`
	Large      string `json:"48x48"`
}

// IssueType ...
type IssueType struct {
	IconURL string `json:"iconUrl"`
	Name    string `json:"name"`
	SubTask bool   `json:"subtask"`
}

// SearchIssueRequest ...
type SearchIssueRequest struct {
	Jql        string `url:"jql,omitempty"`
	Fields     string `url:"fields,omitempty"`
	MaxResults int    `url:"maxResults,omitempty"`
}

// SearchIssueReponse ...
type SearchIssueReponse struct {
	Expand          string   `json:"expand,omitempty"`
	StartAt         int      `json:"startAt,omitempty"`
	MaxResults      int      `json:"maxResults,omitempty"`
	Total           int      `json:"total,omitempty"`
	Issues          []Issue  `json:"issues,omitempty"`
	WarningMessages []string `json:"warningMessages,omitempty"`
}

// Render ...
func (si *SearchIssueReponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}
