package assert

import (
	"testing"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func errorEqual(err1, err2 error) bool {
	return amiderrors.Is(err1, err2)
}

func ErrorEqual(t testing.TB, exp error, act error, message string) {
	if !errorEqual(exp, act) {
		t.Fatalf(message+", expected %v actual is %v", exp, act)
	}
}

func ErrorNotEqual(t testing.TB, exp error, act error, message string) {
	if errorEqual(exp, act) {
		t.Fatalf(message+", actual %v equal %v", act, exp)
	}
}

func AnyOf(t testing.TB, message string, act error, exp ...error) {
	for _, e := range exp {
		if errorEqual(act, e) {
			return
		}
	}
	t.Fatalf(message+", %s not one of %s", act, exp)
}
