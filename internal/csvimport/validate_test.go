package csvimport_test

import (
	"testing"

	"github.com/amidgo/amiddocs/internal/csvimport"
	"github.com/amidgo/amiddocs/pkg/assert"
)

func TestValidateFileNameIsCsv(t *testing.T) {
	emptyFileName := ""
	assert.ErrorEqual(t, nil, csvimport.ValidateFileNameIsCsv(emptyFileName), "empty file name test failed")
	lowLenFileName := "..."
	assert.ErrorEqual(t, nil, csvimport.ValidateFileNameIsCsv(lowLenFileName), "low length file name test failed")
	onlyCsvFileName := ".csv"
	assert.ErrorEqual(t, nil, csvimport.ValidateFileNameIsCsv(onlyCsvFileName), "only csv file name test failed")
	rightFileName := "example.csv"
	assert.ErrorEqual(t, nil, csvimport.ValidateFileNameIsCsv(rightFileName), "right input file name test failed")
	wrongFileName := "example.sql"
	assert.ErrorEqual(t, csvimport.ErrWrongFileFormat, csvimport.ValidateFileNameIsCsv(wrongFileName), "wrong input file name test failed")
}
