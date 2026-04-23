package main

import "fmt"

func main() {
	const ConvertUSDRUB float64 = 75.03
	const ConvertUSDEUR float64 = 0.85
	EURRUB := ConvertUSDRUB / ConvertUSDEUR
	fmt.Println(EURRUB)
}

func userScan() string {
	var str string
	fmt.Scan(&str)
	return str
}

func Convert(amount float64, fromCurrency string, toCurrency string) float64 {

	return 0
}
