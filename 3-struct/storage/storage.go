package storage

import (
	"encoding/json"
	"fmt"
	"main/bins"
	"os"
)

func ReadJson(name string) (*bins.Bin, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var file bins.Bin
	err = json.Unmarshal(b, &file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &file, nil
}
