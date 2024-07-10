package read

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV() (records [][]string, err error) {
	fd, err := os.Open("../problems.csv")
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