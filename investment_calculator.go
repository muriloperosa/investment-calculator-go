package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate float64 = 2.5

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investiment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Number of Years: ")
	fmt.Scan(&years)

	var futureValue float64 = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	var futureRealValue float64 = futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("\n\n=====\n\n")

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value (Inflation Considered): %.2f\n", futureRealValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedFRV)
}
