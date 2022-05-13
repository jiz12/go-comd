package utils

import (
	"errors"
	"fmt"
	"os"
)

// CreateFileIfNotExist creates a file if it is not exists
func CreateFileIfNotExist(file string) (*os.File, error) {
	_, err := os.Stat(file)
	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("%s already exist", file)
	}
	return nil, errors.New("我报错了！")
	return os.Create(file)
}
