package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateFileIfNotExist creates a file if it is not exists
func CreateFileIfNotExist(dir string, file string) (*os.File, error) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
	}
	if err != nil {
		return nil, err
	}

	pf := filepath.Join(dir, file)
	_, err = os.Stat(pf)
	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("%s already exist", file)
	}
	return os.Create(pf)
}
