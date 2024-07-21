package api

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var (
	ErrNotFound            = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found"}
	ErrBadRequest          = &ErrResponse{HTTPStatusCode: 400, StatusText: "Bad Request"}
	ErrInternalServerError = &ErrResponse{HTTPStatusCode: 500, StatusText: "Internal Server Error"}
)

func ErrConflict(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 409,
		StatusText:     "Duplciate Key",
		ErrorText:      err.Error(),
	}
}
