package quiz

import (
	"fmt"
)

func Score(record [][]string) (score int) {
	score = 0

	for i := 1; i < len(record); i++ {
		var userAmswer string
		answer := record[i][1]
		question := record[i][0]
		fmt.Println(question)
		fmt.Scan(&userAmswer)
		if userAmswer == answer {
			score++
		}
	}
	return
}