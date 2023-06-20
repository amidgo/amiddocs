package amidstr

import (
	"fmt"
	"strings"
)

func UnmarshalNullString(b []byte) (string, error) {
	v := string(b)
	v = strings.Trim(v, `"`)
	switch v {
	case "null":
		return "", nil
	default:
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
	switch src := src.(type) {
	case nil:
		return "", nil
	case string:
		return src, nil
	default:
		return "", fmt.Errorf("cannot scan %T into *Email", src)
	}
}
