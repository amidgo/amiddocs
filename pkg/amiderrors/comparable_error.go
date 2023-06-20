package amiderrors

import "errors"

type ComparableError interface {
	Equal(err error) bool
}

func Is(err1, err2 error) bool {
	if err1 == nil {
		return err2 == nil
	}
	if errors.Is(err1, err2) {
		return true
	}
	switch err1 := err1.(type) {
	case ComparableError:
		return err1.Equal(err2)
	default:
		return false
	}
}
