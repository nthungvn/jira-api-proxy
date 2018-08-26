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
		HTTPStatusCode: http.StatusUnauthorized,
		ErrorMessages:  []string{"Login failed"},
	}
}

// ErrInvalidRequest ...
func ErrInvalidRequest(err error) render.Renderer {
	return &ResponseError{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		ErrorMessages:  []string{"Invalid request"},
	}
}

// ErrServerError ...
func ErrServerError(err error) render.Renderer {
	return &ResponseError{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		ErrorMessages:  []string{err.Error()},
	}
}
