package main

import (
	"errors"
	"fmt"
)

const ConvertUSDRUB float64 = 75.03
const ConvertUSDEUR float64 = 0.85

func main() {

	for {
		fromCurrency, toCurrency, err := userScanCurrency()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//fmt.Println(currency)
		num, err := number()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//fmt.Println(num)
		res := Convert(float64(num), fromCurrency, toCurrency)
		fmt.Println(res)
		break
	}

}

func userScanCurrency() (string, string, error) {
	// Исходный список всех доступных валют
	currencies := []string{"USD", "EUR", "RUB"}

	// Первый выбор
	fromCurr, remaining, err := selectCurrency(currencies)
	if err != nil {
		return "", "", err
	}

	// Второй выбор – из оставшихся (первая уже удалена)
	toCurr, _, err := selectCurrency(remaining)
	if err != nil {
		return "", "", err
	}

	return fromCurr, toCurr, nil
}

func number() (int, error) {
	var num int
	fmt.Println("Введите количество валюты:")
	_, err := fmt.Scan(&num)
	if err != nil || num <= 0 {
		return 0, errors.New("некорректный ввод (требуется положительное число)")
	}
	return num, nil
}

func selectCurrency(available []string) (string, []string, error) {
	if len(available) == 0 {
		return "", nil, errors.New("нет доступных валют")
	}

	fmt.Println("Выберите валюту, введя номер:")
	for i, cur := range available {
		fmt.Printf("%d) %s\n", i+1, cur)
	}

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil || choice < 1 || choice > len(available) {
		return "", available, errors.New("некорректный ввод номера")
	}

	// Индекс в слайсе (на 1 меньше)
	idx := choice - 1
	selected := available[idx]

	// Удаляем выбранный элемент из слайса
	newAvailable := append(available[:idx], available[idx+1:]...)

	return selected, newAvailable, nil
}

func Convert(amount float64, fromCurrency string, toCurrency string) float64 {
	fmt.Printf("Конвертируем %.2f %s в %s\n", amount, fromCurrency, toCurrency)

	if fromCurrency == "RUB" || toCurrency == "RUB" {
		// amount/ConvertUSDRUB //руб в дол
		// amount*ConvertUSDRUB //дол в руб
		// amount *(1/ConvertUSDEUR * ConvertUSDRUB)//евр в руб
		// amount/ (1/ConvertUSDEUR * ConvertUSDRUB) //руб в евр
		// amount / ConvertUSDEUR // дол в евр
		// amount * ConvertUSDEUR // евр в дол

		res := amount * ConvertUSDRUB
		return res
	}

	return 42.0 // заглушка
}
