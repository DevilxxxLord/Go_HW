package main

import "fmt"

func main() {
	const USDRUB float64 = 75.03
	const USDEUR float64 = 0.85
	EURRUB := USDRUB / USDEUR
	fmt.Println(EURRUB)
}
