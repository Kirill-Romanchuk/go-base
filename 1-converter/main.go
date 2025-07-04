package main

import "fmt"

func getUserInput() (USD_TO_EUR float64, USD_TO_RUB float64) {
	fmt.Print("Введите курс доллора к евро: ")
	fmt.Scan(USD_TO_EUR)
	fmt.Print("Введите курс доллора к рублю: ")
	fmt.Scan(USD_TO_RUB)
	return USD_TO_EUR, USD_TO_RUB
}

func main() {
	var USD_TO_EUR, USD_TO_RUB = getUserInput()
	var EUR_TO_RUB float64 = USD_TO_RUB / USD_TO_EUR
	fmt.Println(EUR_TO_RUB)
}
