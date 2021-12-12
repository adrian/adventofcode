package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
)

func sumItems(a []int) int {
	var sum int
	for i:=0; i<len(a); i++ {
		sum = sum + a[i]
	}
	return sum
}

func main() {
	inputAges := utils.ReadFileSingleLineOfInts("day6_input")

	ageCounts := make([]int, 9)
	for _, age := range inputAges {
		ageCounts[age] = ageCounts[age] + 1
	}

	for i:=0; i<256; i++ {
		newAgeCounts := make([]int, 9)
		for j:=0; j<len(ageCounts); j++ {
			newAgeCounts[8] = ageCounts[0]
			newAgeCounts[7] = ageCounts[8]
			newAgeCounts[6] = ageCounts[7]+ ageCounts[0]
			newAgeCounts[5] = ageCounts[6]
			newAgeCounts[4] = ageCounts[5]
			newAgeCounts[3] = ageCounts[4]
			newAgeCounts[2] = ageCounts[3]
			newAgeCounts[1] = ageCounts[2]
			newAgeCounts[0] = ageCounts[1]
		}
		ageCounts = newAgeCounts
	}

	fmt.Println(sumItems(ageCounts))
}
