package amiderrors

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error    string `json:"error"`
	Code     string `json:"code"`
	HttpCode int    `json:"http_code"`
	RawError error  `json:"raw"`
}

func NewErrorResponse(err string, httpcode int, code string) *ErrorResponse {
	return &ErrorResponse{Error: err, Code: code, HttpCode: httpcode}
}

func NewInternalErrorResponse(raw error) *ErrorResponse {
	if raw == nil {
		return nil
	}
	return NewErrorResponse(raw.Error(), http.StatusInternalServerError, "internal").WithRaw(raw)
}

func (e *ErrorResponse) WithRaw(raw error) *ErrorResponse {
	e.RawError = raw
	return e
}

func (e *ErrorResponse) SendWithFiber(c *fiber.Ctx) error {
	return c.Status(e.HttpCode).JSON(e)
}
