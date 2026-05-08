package storage

import (
	"fmt"
	"os"
)

func SaveStorage(fileJson []byte, name string) {
	err := os.WriteFile(name, fileJson, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
