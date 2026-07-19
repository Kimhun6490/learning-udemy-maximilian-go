package main

import (
	"fmt"

	"example.com/app/cmdmanager"
	"example.com/app/prices"
)

func main() {
	taxRates := []float64{0.07, 0.08, 0.06, 0.09, 0.05}

	results := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		//fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("tax_included_prices_%.2f.json", taxRate))
		cmdm := cmdmanager.NewCommandManager()

		job := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := job.Calculate()
		if err != nil {
			fmt.Println("Error calculating tax-included prices:", err)
			continue
		}
	}

	fmt.Println("Tax Included Prices:")
	for taxRate, prices := range results {
		fmt.Printf("Tax Rate: %.2f, Prices: %v\n", taxRate, prices)
	}
}
