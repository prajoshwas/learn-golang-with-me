package operationmanager

import (
	"fmt"
)

type OperationManager struct {
}

func (om OperationManager) ReadLinesFromFile() ([]string, error) {
	var prices []string

	fmt.Println("Please Enter your prices")
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}
func (om OperationManager) WriteJSONToFile(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() OperationManager {
	return OperationManager{}
}
