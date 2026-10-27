package main

import (
	"example/calculator/filemanager"
	"example/calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.json", taxRate*100))
		//om := operationmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Printf("Job for tax rate %.2f completed successfully\n", taxRates[index])
		}
	}

	/* for _, errChan := range errorChans {
		<-errChan
	}

	for _, doneChan := range doneChans {
		<-doneChan
	} */
}
