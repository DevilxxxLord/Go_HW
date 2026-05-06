package storage

import (
	"os"
)

func SaveStorage(fileJson []byte, name string) {
	os.WriteFile(name, fileJson, os.FileMode(os.O_CREATE))
}
