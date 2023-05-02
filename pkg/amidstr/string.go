package amidstr

import (
	"errors"
	"fmt"
	"strings"
)

func UnmarshalNullString(b []byte) (string, error) {
	v := string(b)
	switch v {
	case "null":
		return "", nil
	default:
		v = strings.Trim(v, `"`)
		return v, nil
	}
}

func MarshalNullString(s string) ([]byte, error) {
	if len(s) == 0 {
		return []byte("null"), nil
	}
	return []byte(`"` + s + `"`), nil
}

func ScanNullString(src interface{}) (string, error) {
	switch src.(type) {
	case nil:
		return "", nil
	case string:
		return src.(string), nil
	default:
		return "", errors.New(fmt.Sprintf("cannot scan %T into *Email", src))
	}
}
