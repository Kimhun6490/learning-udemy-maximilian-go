package main

import (
	"fmt"
	"math"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue := getUserInput("Enter the revenue: ")
	expenses := getUserInput("Enter the expenses: ")
	taxRate := getUserInput("Enter the tax rate (in %): ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Profit After Tax: %.2f\n", profit)
	fmt.Printf("Profit Ratio: %.2f\n", math.Abs(ratio))
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var input float64
	fmt.Print(infoText)
	fmt.Scan(&input)
	return input
}
