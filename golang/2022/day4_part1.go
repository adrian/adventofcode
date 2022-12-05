package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"strings"
)

type Assignment struct {
	start int
	end   int
}

func main() {
	input := utils.ReadFileIntoArrayOfStrings("day4_input")

	var nested_assignments int

	for i := 0; i < len(input); i++ {
		asignments := strings.Split(input[i], ",")

		asignment1_range := strings.Split(asignments[0], "-")
		assignment1 := Assignment{utils.Atoi(asignment1_range[0]), utils.Atoi(asignment1_range[1])}

		asignment2_range := strings.Split(asignments[1], "-")
		assignment2 := Assignment{utils.Atoi(asignment2_range[0]), utils.Atoi(asignment2_range[1])}

		if assignment1.start >= assignment2.start && assignment1.end <= assignment2.end {
			// assignment 2 is within assignment 1
			nested_assignments = nested_assignments + 1
		} else if assignment2.start >= assignment1.start && assignment2.end <= assignment1.end {
			// assignment 1 is within assignment 2
			nested_assignments = nested_assignments + 1
		}
	}
	fmt.Println(nested_assignments)
}
