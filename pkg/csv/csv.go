package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

// parse csv to list fun with validate check
func ParseCsvToListValidate[T any](reader csv.Reader, validateFunc func(T) error) ([]*T, error) {
	// map to store indexes of column name
	fields := make(map[int]string)
	// map which already store csv tag name to struct field name
	csvFields := csvFields[T]()
	// result values
	values := make([]*T, 0)
	// read row by row
Loop:
	for i := 0; true; i++ {
		// read elems
		elems, err := reader.Read()
		switch err {
		case nil:
			// if it is first row, we parse names to fields map
			if i == 1 {
				fields, err = parseFieldNames(elems)
				if err != nil {
					return nil, err
				}
				continue Loop
			}
			v, err := ParseSingleRow[T](elems, fields, csvFields)
			if err != nil {
				return nil, fmt.Errorf("error in %d row, description %v", i, err)
			}
			err = validateFunc(*v)
			if err != nil {
				return nil, fmt.Errorf("error in %d row, description %v", i, err)
			}
			values = append(values, v)
			continue Loop
		case io.EOF:
			break Loop
		default:
			return nil, err
		}
	}
	return values, nil
}

func ParseCsvToList[T any](reader *csv.Reader) ([]*T, error) {
	// map to store indexes of column name
	fields := make(map[int]string)
	// map which already store csv tag name to struct field name
	csvFields := csvFields[T]()
	// result values
	values := make([]*T, 0)
	// read row by row
Loop:
	for i := 1; true; i++ {
		// read elems
		elems, err := reader.Read()
		switch err {
		case nil:
			// if it is first row, we parse names to fields map
			if i == 1 {
				fields, err = parseFieldNames(elems)
				if err != nil {
					return nil, err
				}
				continue Loop
			}
			v, err := ParseSingleRow[T](elems, fields, csvFields)
			if err != nil {
				return nil, fmt.Errorf("error in %d row, description %v", i, err)
			}
			values = append(values, v)
			continue Loop
		case io.EOF:
			break Loop
		default:
			return nil, err
		}
	}
	return values, nil
}

// parse elems and returns map of elem position to elem
func parseFieldNames(elems []string) (map[int]string, error) {
	fields := make(map[int]string)
	if len(elems) == 0 {
		return nil, errors.New("zero length first line")
	}
	for in, el := range elems {
		// trimCsv trim spaces and '"' symbol
		name := trimCsv(el)
		fields[in] = name
	}
	return fields, nil
}

/*
parse single row of csv data

	fields is map of field position index to fields name
	csvFields is map of struct field name to own csv tag
*/
func ParseSingleRow[T any](elems []string, fields map[int]string, csvFields map[string]string) (*T, error) {
	// if count of row values not equal count of fields function return wrong input error
	if len(elems) != len(fields) {
		return nil, fmt.Errorf(
			"wrong input, number of rows is not equal title fields expected %d actual is %d", len(fields), len(elems),
		)
	}
	// create new variable
	v := new(T)
	// map row elems to struct
	err := parseElems(elems, fields, csvFields, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

type Parser interface {
	Parse(s string, rf *reflect.Value)
}

/*
parseElems godoc

	function parse csv row as elems to struct v
	elems it is csv row
	fields it is map[<index of csv field name>]<csv field name>
	v it is a destination
*/
func parseElems[T any](elems []string, fields map[int]string, csvFields map[string]string, v *T) error {
	// iterate by row elems
	for in, el := range elems {
		csvFieldName, ok := fields[in]
		if !ok {
			continue
		}
		structFieldName, ok := csvFields[csvFieldName]
		if !ok {
			continue
		}
		field := reflect.ValueOf(v).Elem().FieldByName(structFieldName)
		el := trimCsv(el)
		switch f := field.Interface().(type) {
		case string:
			field.SetString(el)
		case int8:
			i, _ := strconv.ParseInt(el, 10, 8)
			field.SetInt(i)
		case int16:
			i, _ := strconv.ParseInt(el, 10, 32)
			field.SetInt(i)
		case int32:
			i, _ := strconv.ParseInt(el, 10, 32)
			field.SetInt(i)
		case int64:
			i, _ := strconv.ParseInt(el, 10, 64)
			field.SetInt(i)
		case uint8:
			i, _ := strconv.ParseUint(el, 10, 8)
			field.SetUint(i)
		case uint16:
			i, _ := strconv.ParseUint(el, 10, 16)
			field.SetUint(i)
		case uint32:
			i, _ := strconv.ParseUint(el, 10, 32)
			field.SetUint(i)
		case uint64:
			i, _ := strconv.ParseUint(el, 10, 64)
			field.SetUint(i)
		case float32:
			i, _ := strconv.ParseFloat(el, 32)
			field.SetFloat(i)
		case float64:
			i, _ := strconv.ParseFloat(el, 64)
			field.SetFloat(i)
		case Parser:
			f.Parse(el, &field)
		default:
			return errors.New("unsupported field type")
		}
	}
	return nil
}

// in first trim '"' in second trim space
func trimCsv(raw string) string {
	name := strings.TrimSpace(raw)
	name = strings.Trim(name, `"`)
	return name
}

// return map [<name of csv tag>]<name of struct field>
func csvFields[T any]() map[string]string {
	fields := make(map[string]string)
	v := *new(T)
	vtype := reflect.TypeOf(v)
	fieldsCount := vtype.NumField()
	for i := 0; i < fieldsCount; i++ {
		field := vtype.Field(i)
		tag, ok := field.Tag.Lookup("csv")
		name := field.Name
		if tag == "-" {
			continue
		}
		if !ok {
			fields[name] = name
		}
		fields[tag] = name
	}
	return fields
}
