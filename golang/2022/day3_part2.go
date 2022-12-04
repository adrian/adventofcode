package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
)

func main() {
	input := utils.ReadFileIntoArrayOfStrings("day3_input")

	sum_priorities := 0

	for i := 0; i < len(input); i = i + 3 {
		rucksack1_contents := input[i]
		rucksack2_contents := input[i + 1]
		rucksack3_contents := input[i + 2]

		for j := 0; j < len(rucksack1_contents); j++ {
			if utils.Contains(rucksack1_contents[j], rucksack2_contents) &&
			   utils.Contains(rucksack1_contents[j], rucksack3_contents) {
				sum_priorities = sum_priorities + ascii_to_priority(int(rucksack1_contents[j]))
				break
			}
		}
	}

	fmt.Println(sum_priorities)
}

func ascii_to_priority(item int) int {
	if item > 96 {
		return item - 96
	}
	return item - 38
}
