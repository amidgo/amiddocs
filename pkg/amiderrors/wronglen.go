package amiderrors

import (
	"fmt"
	"net/http"
)

func WrongLength(name string, min int, max int) *ErrorResponse {
	return NewErrorResponse("Длина параметра "+name+" должна находится в диапазоне от "+fmt.Sprint(min)+" до "+fmt.Sprint(max), http.StatusBadRequest, "wrong_len")
}
