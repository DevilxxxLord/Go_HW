package file

import (
	"os"
	"path/filepath"
)

func ReadFiles(name string) ([]byte, bool, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		return nil, false, err
	}
	if filepath.Ext(name) == ".json" {
		return b, true, err
	}
	return b, false, nil
}
