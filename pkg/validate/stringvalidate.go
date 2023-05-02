package validate

import "github.com/amidgo/amiddocs/pkg/amiderrors"

func StringValidate(s string, name string, min int, max int) error {
	l := len([]rune(s))
	if l < min || l > max {
		return amiderrors.WrongLength(name, min, max)
	}
	return nil
}
