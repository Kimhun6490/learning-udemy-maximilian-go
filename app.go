package main

import (
	"fmt"

	"example.com/app/filemanager"
	"example.com/app/prices"
)

func main() {
	taxRates := []float64{0.07, 0.08, 0.06, 0.09, 0.05}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("tax_included_prices_%.2f.json", taxRate))
		// cmdm := cmdmanager.NewCommandManager()

		job := prices.NewTaxIncludedPriceJob(fm, taxRate)

		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		go job.Calculate(doneChans[i], errorChans[i])
	}

	for i := range taxRates {
		select {
		case <-doneChans[i]:
			fmt.Printf("Tax rate %.2f calculation completed successfully.\n", taxRates[i])
		case err := <-errorChans[i]:
			fmt.Printf("Error occurred for tax rate %.2f: %v\n", taxRates[i], err)
		}
	}
}
