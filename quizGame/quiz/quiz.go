package quiz

import (
	"fmt"
	"time"
)

func Score(record [][]string, limit int, seconds int) (score int) {
	score = 0
	timer := time.NewTimer(time.Duration(seconds) * time.Second)

	for i := 0; i < limit; i++ {
		question := record[i][0]
		fmt.Println(question)
		answerChan := make(chan string)
		go func() {
			var userAmswer string
			fmt.Scan(&userAmswer)
			answerChan <- userAmswer
		}()

		select {
		case <- timer.C:
			fmt.Println("\nTIMER RAN OUT YOU SLOW POKE")
			fmt.Println("Your score is", score, "out of", limit)
			return
		case userAmswer := <-answerChan:
			answer := record[i][1]
			if userAmswer == answer {
				score++
			}
		}
	}
	return
}