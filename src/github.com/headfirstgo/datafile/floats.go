package datafile

import (
	"bufio"
	"os"
	"strconv"
)
// GetFloats reads the string line from specific file and transfer them into a float array
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		item, err := strconv.ParseFloat(scanner.Text(), 64)
		numbers = append(numbers, item)
		if err != nil {
			return nil, err
		}
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}