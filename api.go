package main

// HTTP verbs
const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// APIDeclaration ...
type APIDeclaration struct {
	Method string
	URI    string
}
