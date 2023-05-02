package amidtime

import (
	"fmt"
	"strconv"
	"time"
)

type DateTime time.Time

func (t DateTime) Unix() int64 {
	return t.Time().Unix()
}

func (t DateTime) Time() time.Time {
	return time.Time(t)
}

func (t DateTime) String() string {
	return t.Time().Format(time.DateTime)
}

func (t *DateTime) Scan(src any) error {
	if src == nil {
		*t = DateTime{}
		return nil
	}
	switch src := src.(type) {
	case int64:
		tm := time.Unix(src, 0)
		*t = DateTime(tm)
	case time.Time:
		*t = DateTime(src)
		return nil
	}

	return fmt.Errorf("cannot scan %T", src)
}

func (t *DateTime) UnmarshalJSON(b []byte) error {
	value, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	tm := time.Unix(value, 0)
	*t = DateTime(tm)
	return nil
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprint(t.Unix())), nil
}
