package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DataKey struct {
	Key string
}

func ReadEnv() *DataKey {
	err := godotenv.Load()
	if err != nil {
		//fmt.Println(err)
		panic("err")
	}
	res := os.Getenv("KEY")
	if res == "" {
		panic("Не передан параметр KEY в переменные окружения!")
	}
	return &DataKey{
		Key: res,
	}
}
