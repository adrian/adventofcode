package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
)

func main() {
	inputAges := utils.ReadFileSingleLineOfInts("day6_input")

	ages := inputAges
	for i:=0; i<80; i++ {
		var newAges []int
		for j:=0; j<len(ages); j++ {
			if ages[j] == 0 {
				newAges = append(newAges, 6)
				newAges = append(newAges, 8)
			} else {
				newAges = append(newAges, ages[j]-1)
			}
		}
		ages = newAges
	}

	fmt.Println(len(ages))
}
