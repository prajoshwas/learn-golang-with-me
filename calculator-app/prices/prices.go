package prices

import (
	"errors"
	"example/calculator/conversion"
	"example/calculator/iomanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	Inputprices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOFileManager     iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) Process(doneChan chan<- bool, errorChan chan error) {
	err := job.LoadData()
	errorChan <- errors.New("Error!")
	if err != nil {
		errorChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.Inputprices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result
	job.IOFileManager.WriteJSONToFile(job)

	doneChan <- true
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOFileManager.ReadLinesFromFile()

	if err != nil {
		fmt.Println(err)
		return err
	}

	inputPrices, err := conversion.ConvertToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return err
	}

	job.Inputprices = inputPrices
	return nil
}

// constructors
func NewTaxIncludedPriceJob(io iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:           taxRate,
		Inputprices:       []float64{10, 20, 30},
		TaxIncludedPrices: make(map[string]string),
		IOFileManager:     io,
	}
}
