package amidtime

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	WRONG_TIME_FORMAT = errors.New("wrong time format")
)

type Date time.Time

func (t Date) Time() time.Time {
	return time.Time(t)
}

func (t Date) String() string {
	return t.Time().Format(time.DateOnly)
}

func (t *Date) Scan(src any) error {
	if src == nil {
		*t = Date{}
		return nil
	}
	switch src := src.(type) {
	case string:
		tm, err := time.Parse(time.DateOnly, src)
		if err != nil {
			return WRONG_TIME_FORMAT
		}
		*t = Date(tm)
		return nil
	case time.Time:
		*t = Date(src)
		return nil
	}

	return fmt.Errorf("cannot scan %T", src)
}

func (t *Date) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	tm, err := time.Parse(time.DateOnly, value)
	if err != nil {
		return WRONG_TIME_FORMAT
	}
	*t = Date(tm)
	return nil
}

func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format(time.DateOnly) + `"`), nil
}
