package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
)

func main() {
	input := utils.ReadFileIntoArrayOfInts("day1_input")
	num_increases := 0
	for i := 1; i < len(input)-2; i++ {
		if input[i]+input[i+1]+input[i+2] > input[i-1]+input[i]+input[i+1] {
			num_increases = num_increases + 1
		}
	}
	fmt.Println(num_increases)
}
