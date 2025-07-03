package main

import "fmt"

func main() {
	const USD_TO_EUR float64 = 0.85
	const USD_TO_RUB float64 = 79.03
	var EUR_TO_RUB float64 = USD_TO_RUB / USD_TO_EUR
	fmt.Println(EUR_TO_RUB)
}
