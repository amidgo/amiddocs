package reqfields

import (
	"github.com/amidgo/amiddocs/internal/errorutils/reqerror"
)

type DocumentType string

const (
	STUDY_DOCUMENT DocumentType = "STUDY_DOCUMENT"
)

func (dt DocumentType) Validate() error {
	for _, d := range []DocumentType{STUDY_DOCUMENT} {
		if d == dt {
			return nil
		}
	}
	return reqerror.WRONG_DOCUMENT_TYPE
}
