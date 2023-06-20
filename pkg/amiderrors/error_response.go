package amiderrors

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// error which return transport layer to user
type ErrorResponse struct {
	Err  string `json:"error"`
	Code string `json:"code"`
}

func NewErrorResponse(err, code string) *ErrorResponse {
	return &ErrorResponse{Err: err, Code: code}
}

// error interface implementation
func (e *ErrorResponse) Error() string {
	return e.Err
}

// ComparableError interface implementation
func (e *ErrorResponse) Equal(raw error) bool {
	err, ok := raw.(*ErrorResponse)
	if !ok {
		return false
	}
	return e.Code == err.Code
}

// cause for additional error info
// users dont see causes
// causes use in logs
type Cause struct {
	Action         string
	Method         string
	MethodProvider string
}

func NewCause(action, method, methodProvider string) *Cause {
	return &Cause{Action: action, Method: method, MethodProvider: methodProvider}
}

// Stringer implementation
// return info of cause: action, method, method provider
func (c *Cause) String() string {
	return fmt.Sprintf(
		"Action is %s, Method is %s, MethodProvider is %s",
		c.Action, c.Method, c.MethodProvider,
	)
}

// main error struct of the project
// users dont see exceptions, transport layer map it to ErrorResponse
type Exception struct {
	Err      error
	Code     string
	Type     string
	HttpCode int
	Causes   []*Cause
}

func NewException(httpcode int, etype string, code string) *Exception {
	return &Exception{Code: code, Type: etype, HttpCode: httpcode, Causes: make([]*Cause, 0)}
}

// Stringer implementation,
// return info of cause: err, code, httpcode for rest transport and list of causes if not empty
func (e *Exception) String() string {
	sCauses := make([]string, 0, len(e.Causes))
	for _, cause := range e.Causes {
		sCauses = append(sCauses, cause.String())
	}
	s := "empty"
	if len(sCauses) > 0 {
		s = "\n" + strings.Join(sCauses, "\n")
	}
	return fmt.Sprintf(
		"Err is %s, Code %s, HttpCode %d, Causes: %s",
		e.Err,
		MakeCode(e.Type, e.Code),
		e.HttpCode,
		s,
	)
}

// CError method
// like String but require error config for error mapping
func (e *Exception) CError(c *Config) string {
	sCauses := make([]string, 0, len(e.Causes))
	for _, cause := range e.Causes {
		sCauses = append(sCauses, cause.String())
	}
	s := "empty"
	if len(sCauses) > 0 {
		s = "\n" + strings.Join(sCauses, "\n")
	}
	return fmt.Sprintf(
		"Err is %s, Code %s, HttpCode %d, Internal %v Causes: %s",
		c.ExceptionToResponse(e).Error(),
		MakeCode(e.Type, e.Code),
		e.HttpCode,
		e.Err,
		s,
	)
}

// CError with Fprint
func (e *Exception) Fprint(c *Config, w io.Writer) {
	fmt.Fprint(
		w,
		e.CError(c),
	)
}

// error interface implementation
// return same the String() method
func (e *Exception) Error() string {
	return e.String()
}

// builder pattern implementation
// set Err of Exception and return copy
func (e Exception) WithErr(err error) *Exception {
	e.Err = err
	return &e
}

// builder pattern implementation
// add Cause to Cause List of Exception and return copy
func (e Exception) WithCause(cause *Cause) *Exception {
	e.Causes = append(e.Causes, cause)
	return &e
}

// wrap error with cause
// if type of err is Excpetion then add cause to list and return error
// default return new internal exception with err and cause
func Wrap(err error, cause *Cause) error {
	switch err := err.(type) {
	case *Exception:
		return err.WithCause(cause)
	default:
		return NewException(http.StatusInternalServerError, "common", INTERNAL).WithErr(err).WithCause(cause)
	}
}

// ComparableError interface implementation
// Compare all fields
func (e *Exception) Equal(raw error) bool {
	err, ok := raw.(*Exception)
	if !ok {
		return false
	}
	if !errors.Is(e.Err, err.Err) {
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
