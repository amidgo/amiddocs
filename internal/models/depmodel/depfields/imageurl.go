package depfields

import "github.com/amidgo/amiddocs/pkg/amidstr"

type ImageUrl string

func (i *ImageUrl) UnmarshalJSON(b []byte) error {
	s, err := amidstr.UnmarshalNullString(b)
	*i = ImageUrl(s)
	return err
}

func (i ImageUrl) MarshalJSON() ([]byte, error) {
	return amidstr.MarshalNullString(string(i))
}

func (i *ImageUrl) Scan(src any) error {
	s, err := amidstr.ScanNullString(src)
	*i = ImageUrl(s)
	return err
}
