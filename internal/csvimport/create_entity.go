package csvimport

import (
	"context"
	"encoding/csv"
	"os"

	amidcsv "github.com/amidgo/amiddocs/pkg/csv"
)

func CreateEntityFromCsv[T any](fileName string, createEntityFunc func(ctx context.Context, depList []*T) error) error {
	fileName, err := ValidateFileName(fileName)
	if err != nil {
		return err
	}
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	csvReader := csv.NewReader(file)
	list, err := amidcsv.ParseCsvToList[T](csvReader)
	if err != nil {
		return err
	}
	err = createEntityFunc(context.Background(), list)
	if err != nil {
		return err
	}
	return nil
}
