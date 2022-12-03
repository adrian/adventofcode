package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
)

func main() {
	input := utils.ReadFileIntoArrayOfStrings("day2_input")

	total_score := 0

	for i := 0; i < len(input); i++ {
		game := input[i]

		opponent_plays := string(game[0])
		expected_outcome := string(game[2])
		var i_play string
		var score int

		if (expected_outcome == "X") {
			// Play to lose
			if (opponent_plays == "A") {
				// Lose to rock
				i_play = "C"
			} else if (opponent_plays == "B") {
				// Lose to paper
				i_play = "A"
			} else if (opponent_plays == "C") {
				// Lose to scissors
				i_play = "B"
			}
		} else if (expected_outcome == "Y") {
			// Play to draw
			score = score + 3
			i_play = opponent_plays
		} else if (expected_outcome == "Z") {
			// Play to win
			score = score + 6
			if (opponent_plays == "A") {
				// Win to rock
				i_play = "B"
			} else if (opponent_plays == "B") {
				// Win to paper
				i_play = "C"
			} else if (opponent_plays == "C") {
				// Win to scissors
				i_play = "A"
			}
		}

		score = score + score_play(i_play)

		total_score = total_score + score
	}
	fmt.Println(total_score)
}

func score_play(play string) int {
	if (play == "A") {
		return 1
	} else if (play == "B") {
		return 2
	} else if (play == "C") {
		return 3
	} else {
		panic("Unexpected play " + play)
	}
}