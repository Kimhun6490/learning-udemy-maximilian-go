package prices

import (
	"fmt"

	"example.com/app/conversion"
	"example.com/app/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLinesFromFile()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Error converting prices:", err)
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Calculate() error {
	err := job.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)

		key := fmt.Sprintf("%.2f", price)
		value := fmt.Sprintf("%.2f", taxIncludedPrice)

		result[key] = value
	}

	job.TaxIncludedPrices = result

	fileName := fmt.Sprintf("tax_included_prices_%.2f.json", job.TaxRate)
	err = job.IOManager.WriteLinesToFile(job)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Printf("Tax-included prices saved to %s\n", fileName)
	return nil
}

func NewTaxIncludedPriceJob(ioManager iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:           taxRate,
		InputPrices:       []float64{10, 20},
		TaxIncludedPrices: make(map[string]string),
		IOManager:         ioManager,
	}
}
