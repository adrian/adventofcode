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
		i_play := normalize_my_play(string(game[2]))

		score := score_play(i_play)

		if (i_play == "B" && opponent_plays == "A") {
			// paper beats rock
			score = score + 6
		} else if (i_play == "A" && opponent_plays == "C") {
			// rock beats scissors
			score = score + 6
		} else if (i_play == "C" && opponent_plays == "B") {
			score = score + 6
			// scissors beats paper
		} else if (i_play == opponent_plays) {
			// draw
			score = score + 3
		}

		total_score = total_score + score
	}
	fmt.Println(total_score)
}

func normalize_my_play(play string) string {
	if (play == "X") {
		return "A"
	} else if (play == "Y") {
		return "B"
	} else if (play == "Z") {
		return "C"
	} else {
		panic("Unexpected play " + play)
	}
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