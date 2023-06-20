package csvimport

import (
	"errors"
	"os"
	"path"
)

var (
	ErrWrongFileFormat = errors.New("filename should end with .csv")
	ErrFileNotExist    = errors.New("file not found")
)

func ValidateFileName(fileName string) (string, error) {
	err := ValidateFileNameIsCsv(fileName)
	if err != nil {
		return "", err
	}
	fileName, err = ValidateFileExist(fileName)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

// check file name ends with '.csv', if empty return nil
func ValidateFileNameIsCsv(fileName string) error {
	f := []rune(fileName)
	l := len(f)
	if l == 0 {
		return nil
	}
	if l < 4 {
		return ErrWrongFileFormat
	}
	if string(f[l-4:]) != ".csv" {
		return ErrWrongFileFormat
	}
	return nil
}

// check file exists and return abs path to file
func ValidateFileExist(fileName string) (string, error) {
	if !path.IsAbs(fileName) {
		fileName = path.Join(os.Getenv("PWD"), fileName)
	}
	if !fileExist(fileName) {
		return "", ErrFileNotExist
	}
	return fileName, nil
}

func fileExist(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
