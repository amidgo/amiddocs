package csverror

import (
	"fmt"
	"strings"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const CSV_TYPE = "csv"
const INDEX_CODE = "index"

type ErrorWithIndex struct {
	Err   error
	Index int
}

func (e *ErrorWithIndex) Error() string {
	return fmt.Sprintf("Error in %d index (%d row), description: %v", e.Index, e.Index+2, e.Err)
}

func NewErrorWithIndex(err error, index int) *ErrorWithIndex {
	return &ErrorWithIndex{Err: err, Index: index}
}

func (e *ErrorWithIndex) ErrorResponse(c *amiderrors.Config) *amiderrors.ErrorResponse {
	csv, ok := c.Errors[CSV_TYPE]
	if !ok {
		return c.Internal()
	}
	msg, ok := csv[INDEX_CODE]
	if len(msg) == 0 || !ok {
		return c.Internal()
	}
	replacer := strings.NewReplacer("${row}", fmt.Sprint(e.Index+2), "${error}", e.Err.Error())
	msg = replacer.Replace(msg)
	return amiderrors.NewErrorResponse(msg, amiderrors.MakeCode(CSV_TYPE, INDEX_CODE))
}
