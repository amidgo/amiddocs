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

func (t Date) T() time.Time {
	return time.Time(t)
}

func (t Date) String() string {
	return t.T().Format(time.DateOnly)
}

// returns string in dd.mm.yyyy format
func (t Date) Human() string {
	time := t.T()
	return fmt.Sprintf("%02d.%02d.%d", time.Day(), time.Month(), time.Year())
}

func (t *Date) Scan(src any) error {
	switch src := src.(type) {
	case nil:
		*t = Date{}
		return nil
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
	return []byte(`"` + t.T().Format(time.DateOnly) + `"`), nil
}
