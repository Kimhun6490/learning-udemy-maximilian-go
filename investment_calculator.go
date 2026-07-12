package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected annual return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	formattedFutureValue := fmt.Sprintf("Future Value of Investment: $%.2f\n", futureValue)
	formattedFutureRealValue := fmt.Sprintf("Future Real Value of Investment (Adjusted for Inflation): $%.2f\n", futureRealValue)

	fmt.Print(formattedFutureValue)
	fmt.Print(formattedFutureRealValue)
}
