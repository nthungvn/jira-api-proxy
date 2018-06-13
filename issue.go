package main

var searchGet = APIDeclaration{
	Method: GET,
	URI:    "api/2/search",
}

var searchPost = APIDeclaration{
	Method: POST,
	URI:    "api/2/search",
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
	Self   string `json:"self"`
	ID     string `json:"id"`
	Key    string `json:"key"`
	Fields Field  `json:"fields"`
}

// Field ...
type Field struct {
	IssueType IssueType     `json:"issuetype"`
	Parent    SimpleIssue   `json:"parent"`
	Priority  Priority      `json:"priority"`
	Assignee  Person        `json:"assignee"`
	Status    Status        `json:"status"`
	Summary   string        `json:"summary"`
	SubTasks  []SimpleIssue `json:"subtasks"`
	Reporter  Person        `json:"reporter"`
	Point     int           `json:"customfield_10002"`
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
