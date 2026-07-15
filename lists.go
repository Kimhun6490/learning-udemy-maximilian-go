package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	// var prices []float64 = []float64{19.99, 29.99, 39.99}
	prices := [4]float64{19.99, 29.99, 39.99, 49.99}
	fmt.Println(prices)
	fmt.Println(prices[3])

	// featuredPrices := prices[1:3]
	featuredPrices := prices[:3]
	highlightedPrices := featuredPrices[1:]
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices))
	fmt.Println(cap(highlightedPrices))

	var productNames []string = []string{"Laptop", "Smartphone", "Tablet"}
	productNames = append(productNames, "Smartwatch", "Headphones")
	fmt.Println(productNames)

	discountedPrices := []float64{9.99, 14.99, 19.99}
	anotherPrices := append(prices[1:], discountedPrices...)
	fmt.Println(anotherPrices)
}
