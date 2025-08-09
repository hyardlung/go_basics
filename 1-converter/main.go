package main

import "fmt"

func main() {
	const USD_TO_EUR = 0.86
	const USD_TO_RUB = 79.74
	eurToRub := USD_TO_RUB / USD_TO_EUR

	fmt.Print(eurToRub)
}
