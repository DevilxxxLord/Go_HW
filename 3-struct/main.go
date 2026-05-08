package main

import (
	"encoding/json"
	"fmt"
	"main/bins"
	"main/file"
	"main/storage"
)

// type Account struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

func main() {
	fileName := "account.json"
	id := "1111"
	priv := true
	name := "qweqeqw"

	b := &bins.Bin{}
	b.CreatedBinList(id, priv, name)

	bins.BinList = append(bins.BinList, *b) //Нужно это или нет, пока не понятно

	fileByte, err := ToBit(*b)
	if err != nil {
		fmt.Println(err)
		return
	}
	storage.SaveStorage(fileByte, fileName)

	str, err := storage.ReadJson(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла json")
		return
	}
	fmt.Println(*str)

	fileByteNoJson, jsOrNot, err := file.ReadFiles("account.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	if jsOrNot {
		fmt.Println("Файл с разрешением .json")
	}
	fmt.Println(fileByteNoJson)
}

func ToBit(acc bins.Bin) ([]byte, error) {
	fileJson, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return fileJson, nil
}
