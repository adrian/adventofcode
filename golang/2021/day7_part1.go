package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"math"
)

func fuelForPostion(targetPosition int, startingPositions []int) uint64 {
	var fuel uint64
	for _, position := range startingPositions {
		fuel = fuel + uint64(math.Abs(float64(targetPosition-position)))
	}
	return fuel
}

func main() {
	inputPositions := utils.ReadFileSingleLineOfInts("day7_input")

	minPosition := utils.Min(inputPositions)
	maxPosition := utils.Max(inputPositions)

	var minFuel uint64
	minFuel = math.MaxUint64

	for i := minPosition; i <= maxPosition; i++ {
		fuel := fuelForPostion(i, inputPositions)
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	fmt.Println(minFuel)
}
