package file

import (
	"errors"
	"os"
	"path/filepath"
)

func ReadFiles(name string) ([]byte, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	if filepath.Ext(name) == ".json" {
		err := errors.New("Расширение файла не должно быть .json")
		return nil, err
	}
	return b, nil
}
