package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	// type conversions and explicit type assignment
	/*
		var investmentAmount float64 = 1000
		var expectedReturnRate = 5.5
		var years float64 = 10

		var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	*/

	// using alternative variable declaration styles
	const inflationRate = 2.5
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	// fmt.Print("Enter investment amount: ")
	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Enter expected return rate: ")
	outputText("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Enter years: ")
	outputText("Enter years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	// creating formated strings
	formatedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formatedRFV := fmt.Sprintf("Future real value: %.2f\n", futureRealValue)

	// fmt.Printf(`Future value: %.2f\n
	// Future real value: %.2f\n`, futureValue, futureRealValue)
	// fmt.Printf("Type Future real value: %T\n", futureRealValue)

	fmt.Print(formatedFV, formatedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

// alternative return value syntax
func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	// fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// rfv := fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// can use only return
}
