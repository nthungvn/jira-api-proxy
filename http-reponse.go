package main

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse ...
type ErrorResponse struct {
	Err            error             `json:"-"`
	HTTPStatusCode int               `json:"-"`
	ErrorMessages  []string          `json:"errorMessages"`
	Errors         map[string]string `json:"errors,omitempty"`
}

// Render ...
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrorUnauthorized ...
func ErrorUnauthorized(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 401,
		ErrorMessages:  []string{"Login failed"},
	}
}

// ErrorInvalidRequest ...
func ErrorInvalidRequest(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 400,
		ErrorMessages:  []string{"Invalid request"},
	}
}
