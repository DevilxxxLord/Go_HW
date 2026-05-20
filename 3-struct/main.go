package main

import (
	"encoding/json"
	"fmt"
	"main/adapted"
	"main/bins"
)

func main() {
	var all adapted.All = &adapted.Adapted{}
	var BinList = []bins.Bin{}
	fileName := "account.json"
	id := "3333"
	priv := true
	name := "qweqeqw"

	b := &bins.Bin{}
	b.CreatedBinList(id, priv, name)

	BinList = append(BinList, *b)

	fileByte, err := ToBit(BinList)
	if err != nil {
		fmt.Println(err)
		return
	}
	all.SaveStorage(fileByte, fileName)

	str, err := all.ReadJson(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла json")
		return
	}
	fmt.Println(str)

	fileByteNoJson, jsOrNot, err := all.ReadFiles("account.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	if jsOrNot {
		fmt.Println("Файл с разрешением .json")
	}
	fmt.Println(fileByteNoJson)
}

func ToBit(acc []bins.Bin) ([]byte, error) {
	fileJson, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return fileJson, nil
}
