package main

import "fmt"

func main() {
	const ConvertUSDRUB float64 = 75.03
	const ConvertUSDEUR float64 = 0.85
	EURRUB := ConvertUSDRUB / ConvertUSDEUR
	fmt.Println(EURRUB)
}
