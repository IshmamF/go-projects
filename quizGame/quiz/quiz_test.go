package quiz

import (
	"testing"
)

func TestQuiz(t *testing.T) {
	record := [][]string{{"5+5", "10"}}
	score := Score(record, 5, 2)
	expected := 0
	if score != expected {
		t.Errorf("Expected %v but got %v", expected, score)
	}
}