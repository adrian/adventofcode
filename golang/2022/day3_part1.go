package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
)

func main() {
	input := utils.ReadFileIntoArrayOfStrings("day3_input")

	sum_priorities := 0

	for i := 0; i < len(input); i++ {
		rucksack_contents := input[i]
		compartment1 := rucksack_contents[:len(rucksack_contents)/2]
		compartment2 := rucksack_contents[len(rucksack_contents)/2:]
		for j := 0; j < len(compartment1); j++ {
			if utils.Contains(compartment1[j], compartment2) {
				sum_priorities = sum_priorities + ascii_to_priority(int(compartment1[j]))
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
