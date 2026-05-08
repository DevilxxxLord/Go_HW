package storage

import (
	"os"
)

func SaveStorage(fileJson []byte, name string) {
	os.WriteFile(name, fileJson, 0644)
}
