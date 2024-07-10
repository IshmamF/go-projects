package main

import (
	"fmt"
	"go-projects/quizGame/readcsv"
)

func main() {
	record, err := read.ReadCSV()
	if err != nil {
		fmt.Println("Error reading file")
	} else {
		fmt.Println("File read successfully")
		fmt.Println(record)
	}
}