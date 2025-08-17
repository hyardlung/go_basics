package main

import (
	"errors"
	"fmt"
	"strconv"
)

type stringsFloat = map[string]float64
var rates = stringsFloat {
	"USD_TO_EUR": 0.85,
	"USD_TO_RUB": 79.72,
	"EUR_TO_USD": 1.17,
	"RUB_TO_USD": 0.013,
	"EUR_TO_RUB": 93.32,
	"RUB_TO_EUR": 0.011,
}

func main() {
	fmt.Println("___ Currency Converter ___")
	fmt.Println("Enter source currency (USD, EUR, RUB): ")

	var sourceCurrency string
	for {
		tempSource, err := getSourceCurrency()
		if err != nil {
			fmt.Println(err)
			continue
		}
		sourceCurrency = tempSource
		break
	}

	switch sourceCurrency {
	case "USD":
		fmt.Println("Enter target currency (EUR or RUB): ")
	case "EUR":
		fmt.Println("Enter target currency (USD or RUB): ")
	case "RUB":
		fmt.Println("Enter target currency (USD or EUR): ")
	}

	var targetCurrency string
	for {
		tempTarget, err := getTargetCurrency(sourceCurrency)
		if err != nil {
			fmt.Println(err)
			continue
		}
		targetCurrency = tempTarget
		break
	}

	fmt.Println("Enter amount:")
	var amount float64
	for {
		tempAmount, err := getAmount()
		if err != nil {
			fmt.Println(err)
			continue
		}
		amount = tempAmount
		break
	}

	result := convertCurrencies(amount, sourceCurrency, targetCurrency)

	fmt.Println("____________________________________________")
	fmt.Printf("%v %s at the current exchange rate is %.2f %s.", amount, sourceCurrency, result, targetCurrency)
}

func getSourceCurrency() (source string, err error) {
	
	fmt.Scan(&source)
	if checkAcceptableCurrency(source) {
		return source, nil
	} else {
		return "", errors.New("only USD, EUR or RUB, try again")
	}
}

func getTargetCurrency(source string) (target string, err error) {
	fmt.Scan(&target)
	if !checkAcceptableCurrency(target) {
		return "", errors.New("incorrect value, try again")
	} else if source == target {
		return "", errors.New("can't be the same currency, try again")
	} else {
		return target, nil
	}
}

func getAmount() (float64, error)  {
	var value string
	fmt.Scan(&value)

	amount, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, errors.New("incorrect amount value, try again")
	}
	return amount, nil
}

func checkAcceptableCurrency(curr string) bool {
	if curr == "USD" || curr == "EUR" || curr == "RUB" {
		return true
	}
	return false
}

func convertCurrencies(amount float64, from, to string) float64 {
	rate := from + "_TO_" + to
	result := amount * rates[rate]
	return result
}
