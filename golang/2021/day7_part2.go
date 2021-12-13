package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"math"
)

func fuelForPostion(targetPosition int, startingPositions []int) int {
	var totalFuel int
	for _, position := range startingPositions {
		fuelForPostion := utils.SumNumbersUpTo(int(math.Abs(float64(targetPosition-position)) + 1))
		totalFuel = totalFuel + fuelForPostion
	}
	return totalFuel
}

func main() {
	inputPositions := utils.ReadFileSingleLineOfInts("day7_input")

	minPosition := utils.Min(inputPositions)
	maxPosition := utils.Max(inputPositions)

	minFuel := math.MaxInt

	for i := minPosition; i <= maxPosition; i++ {
		fuel := fuelForPostion(i, inputPositions)
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	fmt.Println(minFuel)
}
