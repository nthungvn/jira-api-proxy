package main

var loginAPI = APIDeclaration{
	Method: POST,
	URI:    "rest/auth/1/session",
}

var currentUserAPI = APIDeclaration{
	Method: GET,
	URI:    "rest/auth/1/session",
}

// Session ...
type Session struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// LoginInfo ...
type LoginInfo struct {
	FailedLoginCount    int    `json:"failedLoginCount"`
	LoginCount          int    `json:"loginCount"`
	LastFailedLoginTime string `json:"lastFailedLoginTime"`
	PreviousLoginTime   string `json:"previousLoginTime"`
}

// Authentication ...
type Authentication struct {
	Session   Session   `json:"session"`
	LoginInfo LoginInfo `json:"loginInfo"`
}

// CurrentSession ...
type CurrentSession struct {
	Name      string    `json:"name"`
	LoginInfo LoginInfo `json:"loginInfo"`
}

// User ...
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a Authentication) cookie() string {
	return a.Session.Name + "=" + a.Session.Value
}
