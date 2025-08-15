package main

import (
	"fmt"
	"strconv"
	"strings"
	"slices"
)

var avg, sum, med = "AVG", "SUM", "MED"
var operation string

func main() {
	fmt.Println("___ Calculator ___")
	fmt.Print("Enter operation: ")
	
	for {
		operation = getOperation()
		err := validateInputOperation(operation)
		if err == nil {
			break
		}
		fmt.Print(err)
	}
	
	fmt.Println("Enter numbers (separated by commas): ")
	numbers := getNumbers()
	convertedSlice := convertStrToSlice(numbers)

	switch operation {
	case "SUM":
		sum := calculateSum(convertedSlice)
		fmt.Printf("Sum of your numbers is: %.2f", sum)
	case "AVG":
		avg := calculateAvg(convertedSlice)
		fmt.Printf("Average of your numbers is: %.2f", avg)
	case "MED":
		med := calculateMed(convertedSlice)
		fmt.Printf("Mediane of your numbers is: %.1f", med)
	}
}

func getOperation() string {
	var op string
	fmt.Scan(&op)
	return strings.ToUpper(op)
}

func getNumbers() string {
	var numbers string
	fmt.Scan(&numbers)
	return numbers
}

func validateInputOperation(op string) error {
	switch op {
	case avg, sum, med:
		return nil
	default:
		return fmt.Errorf("invalid operation. Only %s, %s or %s: ", avg, sum, med)
	}
}

func convertStrToSlice(str string) []float64 {
	subStrs := strings.Split(str, ",")
	result := make([]float64, 0)
	for _, char := range subStrs {
		num, _ := strconv.ParseFloat(char, 64)
		result = append(result, num)
	}
	return result
}

func calculateSum(slice []float64) float64 {
	var sum float64
	for _, val := range slice {
		sum += val
	}
	return sum
}

func calculateAvg(slice []float64) float64 {
	sum := calculateSum(slice)
	sliceLen := float64(len(slice))
	return sum / sliceLen
}

func calculateMed(slice []float64) float64 {
	slices.Sort(slice)
	n := len(slice)
	if n % 2 == 1 {
		return slice[n / 2]
	}
	return (slice[n / 2 - 1] + slice[n / 2]) / 2
}