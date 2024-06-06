package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

const inflationRate float64 = 2.5

func main() {

	var fileName string = getFilename()

	var investmentAmount float64 = input("Investiment Amount: ")
	var expectedReturnRate float64 = input("Expected Return Rate: ")
	var years float64 = input("Number of Years: ")

	futureValue, futureRealValue := calculeFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("\n\n=====\n\n")

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value (Inflation Considered): %.2f\n", futureRealValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedFRV)

	storeResults(fileName, investmentAmount, expectedReturnRate, years, futureValue, futureRealValue)
}

func getFilename() string {
	var value string = "resume"
	fmt.Print("Enter a name to your simulation: ")
	fmt.Scan(&value)
	return value
}

func input(label string) float64 {
	var value float64
	fmt.Print(label)
	fmt.Scan(&value)

	if value <= 0 {
		panic("ERROR: The value provided should be greater than zero.")
	}

	return value
}

func calculeFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}

func storeResults(fileName string, investmentAmount float64, expectedReturnRate float64, years float64, futureValue float64, futureRealValue float64) {
	resultsTxt := fmt.Sprintf("Investment amount: %.2f\nExpected Return Rate: %.2f\nNumber of Years: %.0f\n\n========================\n\nFuture Value: %.2f\nFuture Real Value (Inflation Considered): %.2f", investmentAmount, expectedReturnRate, years, futureValue, futureRealValue)
	os.WriteFile(fileName+"-"+time.Now().Format("2006-01-02 15:04:05")+".txt", []byte(resultsTxt), 0644)
}
