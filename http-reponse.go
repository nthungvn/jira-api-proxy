package main

import (
	"net/http"

	"github.com/go-chi/render"
)

// ResponseError ...
type ResponseError struct {
	Err            error             `json:"-"`
	HTTPStatusCode int               `json:"-"`
	ErrorMessages  []string          `json:"errorMessages"`
	Errors         map[string]string `json:"errors,omitempty"`
}

// Render ...
func (e *ResponseError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrUnauthorized ...
func ErrUnauthorized(err error) render.Renderer {
	return &ResponseError{
		Err:            err,
		HTTPStatusCode: 401,
		ErrorMessages:  []string{"Login failed"},
	}
}

// ErrInvalidRequest ...
func ErrInvalidRequest(err error) render.Renderer {
	return &ResponseError{
		Err:            err,
		HTTPStatusCode: 400,
		ErrorMessages:  []string{"Invalid request"},
	}
}
