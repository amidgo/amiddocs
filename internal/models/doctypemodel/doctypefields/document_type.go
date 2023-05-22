package doctypefields

import (
	"github.com/amidgo/amiddocs/internal/errorutils/reqerror"
)

type DocumentType string

const (
	STUDY_DOCUMENT_BUDGET    DocumentType = "STUDY_DOCUMENT_BUDGET"
	STUDY_DOCUMENT_NO_BUDGET DocumentType = "STUDY_DOCUMENT_NO_BUDGET"
)

func (dt DocumentType) Validate() error {
	for _, d := range []DocumentType{STUDY_DOCUMENT_BUDGET, STUDY_DOCUMENT_NO_BUDGET} {
		if d == dt {
			return nil
		}
	}
	return reqerror.WRONG_DOCUMENT_TYPE
}
