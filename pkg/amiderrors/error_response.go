package amiderrors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Code string

type ErrorResponse struct {
	Err      string   `json:"error"`
	Code     Code     `json:"code"`
	HttpCode int      `json:"-"`
	RawError error    `json:"-"`
	Causes   []*Cause `json:"-"`
}

type Cause struct {
	Action         string
	Method         string
	MethodProvider string
}

func (c *Cause) String() string {
	return fmt.Sprintf(
		"Action is %s, Method is %s, MethodProvider is %s",
		c.Action, c.Method, c.MethodProvider,
	)
}

func NewCause(action, method, methodProvider string) *Cause {
	return &Cause{Action: action, Method: method, MethodProvider: methodProvider}
}

func NewErrorResponse(err string, httpcode int, code Code) *ErrorResponse {
	return &ErrorResponse{Err: err, Code: code, HttpCode: httpcode, Causes: make([]*Cause, 0)}
}

func NewInternalErrorResponse(raw error, cause *Cause) *ErrorResponse {
	if raw == nil {
		return nil
	}
	switch raw.(type) {
	case *ErrorResponse:
		return raw.(*ErrorResponse).WithCause(cause)
	default:
		return NewErrorResponse(raw.Error(), http.StatusInternalServerError, "internal").WithRaw(raw).WithCause(cause)
	}
}

func (e *ErrorResponse) WithRaw(raw error) *ErrorResponse {
	e.RawError = raw
	return e
}

func (e *ErrorResponse) WithCause(cause *Cause) *ErrorResponse {
	e.Causes = append(e.Causes, cause)
	return e
}

func (e *ErrorResponse) SendWithFiber(c *fiber.Ctx) error {
	return c.Status(e.HttpCode).JSON(e)
}

func (e *ErrorResponse) Error() string {
	return e.Err
}

func Wrap(err error, cause *Cause) error {
	switch err.(type) {
	case *ErrorResponse:
		return err.(*ErrorResponse).WithCause(cause)
	default:
		return NewInternalErrorResponse(err, cause)
	}
}

func (e *ErrorResponse) Equal(raw error) bool {
	err, ok := raw.(*ErrorResponse)
	if !ok {
		return false
	}
	if e.Err != err.Err {
		return false
	}
	if e.Code != err.Code {
		return false
	}
	if e.HttpCode != err.HttpCode {
		return false
	}
	return true
}

func Is(err1, err2 error) bool {
	if err1 == nil {
		return err2 == nil
	}
	if errors.Is(err1, err2) {
		return true
	}
	switch err1.(type) {
	case *ErrorResponse:
		return err1.(*ErrorResponse).Equal(err2)
	default:
		return false
	}
}
