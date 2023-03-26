package amiderrors

import (
	"net/http"
	"strings"
)

func EmptyValues(args ...string) *ErrorResponse {
	return NewErrorResponse("Значение параметров "+strings.Join(args, ", ")+" не должны быть пустыми", http.StatusBadRequest, "empty_values")
}

func EmptyValue(name string) *ErrorResponse {
	return NewErrorResponse("Значение параметра "+name+" не должно быть пустым", http.StatusBadRequest, "empty_value")
}
