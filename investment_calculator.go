package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {

	var investmentAmount float64 = input("Investiment Amount: ")
	var expectedReturnRate float64 = input("Expected Return Rate: ")
	var years float64 = input("Number of Years: ")

	futureValue, futureRealValue := calculeFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("\n\n=====\n\n")

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value (Inflation Considered): %.2f\n", futureRealValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedFRV)
}

func input(label string) float64 {
	var value float64
	fmt.Print(label)
	fmt.Scan(&value)
	return value
}

func calculeFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}
