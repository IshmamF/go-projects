package read

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV(filepath string) (records [][]string, err error) {
	fd, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file")
		return nil, err
	} else {
		fmt.Println("File opened successfully")
	}
	csvReader := csv.NewReader(fd)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file")
		return nil, err
	} else {
		fmt.Println("File read successfully")
	}
	defer fd.Close()
	return data, nil
}