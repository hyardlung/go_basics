package main

import "fmt"

func main() {
	fmt.Println("___ Currency Converter ___")
	money := getUserInput()

	const USD_TO_EUR = 0.86
	const USD_TO_RUB = 79.74
	eurToRub := USD_TO_RUB / USD_TO_EUR

	fmt.Print(eurToRub)
}

func getUserInput() (amount float64) {
	fmt.Print("	Enter amount in USD: ")
	fmt.Scanln(&amount)
	return amount
}

func calculateAmount(amount float64, sourceCurrency, targetCurrency string) float64 {
	// ...
}
