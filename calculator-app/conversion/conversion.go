package conversion

import (
	"errors"
	"strconv"
)

func ConvertToFloat(lines []string) ([]float64, error) {
	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {

			return nil, errors.New("Failed to convert string to float")
		}
		prices[lineIndex] = floatPrice
	}
	return prices, nil
}
