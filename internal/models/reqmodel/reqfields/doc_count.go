package reqfields

import (
	"github.com/amidgo/amiddocs/internal/errorutils/reqerror"
)

type DocumentCount uint8

const (
	MIN_DOCUMENT_COUNT uint8 = 1
	MAX_DOCUMENT_COUNT uint8 = 3
)

func (dc DocumentCount) Validate() error {
	if dc < DocumentCount(MIN_DOCUMENT_COUNT) || dc > DocumentCount(MAX_DOCUMENT_COUNT) {
		return reqerror.WRONG_DOCUMENT_COUNT
	}
	return nil
}
