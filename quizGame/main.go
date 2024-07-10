package main

import (
	"fmt"
	"go-projects/quizGame/read"
	"go-projects/quizGame/quiz"
)

var (
	filepath = "problems.csv"
	score int
)


func main() {
	record, err := read.ReadCSV(filepath)
	if err != nil {
		fmt.Println(err)
	} 

	score = quiz.Score(record)
	fmt.Println("Your score is", score, "out of", len(record) - 1)
}