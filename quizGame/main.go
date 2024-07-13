package main

import (
	"flag"
	"fmt"
	"go-projects/quizGame/quiz"
	"go-projects/quizGame/read"
)

var (
	filepath = "problems.csv"
	score    int
)

func main() {
	record, err := read.ReadCSV(filepath)
	if err != nil {
		fmt.Println(err)
	}
	limit := flag.Int("limit", len(record), "Number of questions to be asked")
	timer := flag.Int("time", 30, "Time limit on answering questions in seconds")
	flag.Parse()
	score = quiz.Score(record, *limit, *timer)
}
