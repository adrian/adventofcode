package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"strconv"
	"os"
)

func main() {
	input := utils.ReadFileIntoArrayOfStrings("day1_input")
	max_calories := 0
	total_calories := 0

	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			if total_calories > max_calories {
				max_calories = total_calories
			}
			total_calories = 0
			continue
		}
		calories, err := strconv.Atoi(input[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		total_calories = total_calories + calories
	}
	fmt.Println(max_calories)
}
