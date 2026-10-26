package main

import (
	"example/calculator/filemanager"
	"example/calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices2.txt", fmt.Sprintf("result_%v.json", taxRate*100))
		//om := operationmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}

	}
}
