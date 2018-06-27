package main

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse ...
type ErrorResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	AppCode        int64  `json:"code,omitempty"`
	ErrorText      string `json:"error,omitempty"`
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
		StatusText:     "Unauthorized",
		ErrorText:      err.Error(),
	}
}

// ErrorInvalidRequest ...
func ErrorInvalidRequest(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request",
		ErrorText:      err.Error(),
	}
}
