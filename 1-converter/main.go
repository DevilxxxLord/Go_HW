package main

import (
	"errors"
	"fmt"
)

const ConvertUSDRUB float64 = 75.03
const ConvertUSDEUR float64 = 0.85

type mapStrFloat = map[string]float64
type mapInMap = map[string]map[string]float64

func main() {

	for {
		fromCurrency, toCurrency, err := userScanCurrency()
		if err != nil {
			fmt.Println(err)
			continue
		}
		num, err := number()
		if err != nil {
			fmt.Println(err)
			continue
		}

		mRUB := mapStrFloat{"USD": float64(num) / ConvertUSDRUB, "EUR": float64(num) / (1 / ConvertUSDEUR * ConvertUSDRUB)}
		mUSD := mapStrFloat{"RUB": float64(num) * ConvertUSDRUB, "EUR": float64(num) / ConvertUSDEUR}
		mEUR := mapStrFloat{"RUB": float64(num) * (1 / ConvertUSDEUR * ConvertUSDRUB), "USD": float64(num) * ConvertUSDEUR}
		m := mapInMap{"RUB": mRUB, "USD": mUSD, "EUR": mEUR}
		mMap := &m
		res := resConvert(mMap, fromCurrency, toCurrency)
		fmt.Printf("Результат конвертации %s в %s: %.2f", fromCurrency, toCurrency, res)
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

func resConvert(m *mapInMap, fromCurrency string, toCurrency string) float64 {
	// mRUB := mapStrFloat{"USD": amount / ConvertUSDRUB, "EUR": amount / (1 / ConvertUSDEUR * ConvertUSDRUB)}
	// mUSD := mapStrFloat{"RUB": amount * ConvertUSDRUB, "EUR": amount / ConvertUSDEUR}
	// mEUR := mapStrFloat{"RUB": amount * (1 / ConvertUSDEUR * ConvertUSDRUB), "USD": amount * ConvertUSDEUR}

	// m := mapInMap{"RUB": mRUB, "USD": mUSD, "EUR": mEUR}
	res := (*m)[fromCurrency][toCurrency]
	return res
}
