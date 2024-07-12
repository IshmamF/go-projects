package quiz

import (
	"fmt"
)

func Score(record [][]string, limit int) (score int) {
	score = 0

	for i := 0; i < limit; i++ {
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