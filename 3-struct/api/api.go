package api

import (
	"fmt"
	"main/config"
)

func ConnectApi() {
	config.ReadEnv()
	key := config.ReadEnv().Key
	fmt.Println(key)
}
