package main

import (
	"fmt"
)

func getUserInput() (value float64, currencyFrom string, currencyTo string) {
	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&value)
	fmt.Print("Введите исходную валюту: ")
	fmt.Scan(&currencyFrom)
	fmt.Print("Введите желаемую валюту: ")
	fmt.Scan(&currencyTo)
	return value, currencyFrom, currencyTo
}

func calculate(value float64, currencyFrom string, currencyTo string) float64 {
	return 0.0
}

func main() {
	var value, currencyFrom, currencyTo = getUserInput()
	fmt.Print(calculate(value, currencyFrom, currencyTo))
}
