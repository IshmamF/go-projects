package read

import (
	"testing"
	"fmt"
)

func TestReadCSV(t *testing.T) {
	record, err := ReadCSV("../problems.csv")
	expected := []string{"5+5", "10"}
	fmt.Println(record)

	if err != nil {
		t.Errorf("Error reading file")
	}

	firstElement := record[1]
	
	
	if firstElement[0] != expected[0] || firstElement[1] != expected[1] {
		t.Errorf("Expected %v but got %v", expected, record)
	}
}