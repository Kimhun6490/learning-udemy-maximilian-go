package main

import (
	"fmt"
	"math"
)

const INFLATION_RATE = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected annual return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// formattedFutureValue := fmt.Sprintf("Future Value of Investment: $%.2f\n", futureValue)
	// formattedFutureRealValue := fmt.Sprintf("Future Real Value of Investment (Adjusted for Inflation): $%.2f\n", futureRealValue)

	// fmt.Print(formattedFutureValue)
	// fmt.Print(formattedFutureRealValue)

	fmt.Printf(`Future Value of Investment: $%.2f

Future Real Value of Investment (Adjusted for Inflation): $%.2f
`, futureValue, futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rf := fv / math.Pow((1+INFLATION_RATE/100), years)

	return fv, rf
}
