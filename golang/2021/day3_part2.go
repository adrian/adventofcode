package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
)

func main() {
	input := utils.ReadFileIntoArrayOfStrings("day3_input")

	oxygenGeneratorRatings := input
	i := 0
	for len(oxygenGeneratorRatings) > 1 {
		m := mostCommon(oxygenGeneratorRatings, i)
		oxygenGeneratorRatings = filter(oxygenGeneratorRatings, i, m)
		i++
	}

	co2ScrubberRatings := input
	i = 0
	for len(co2ScrubberRatings) > 1 {
		m := leastCommon(co2ScrubberRatings, i)
		co2ScrubberRatings = filter(co2ScrubberRatings, i, m)
		i++
	}

	oxygenGeneratorRating := utils.ConvertBinaryStringToDecimal(oxygenGeneratorRatings[0])
	co2ScrubberRating := utils.ConvertBinaryStringToDecimal(co2ScrubberRatings[0])

	fmt.Println(oxygenGeneratorRating * co2ScrubberRating)
}

func mostCommon(numbers []string, position int) int {
	ones := 0
	for _, binary_num := range numbers {
		if binary_num == "" {
			continue
		}
		ones = ones + int(binary_num[position]-48)
	}
	if ones >= len(numbers)-ones {
		return 1
	}
	return 0
}

func leastCommon(numbers []string, position int) int {
	ones := 0
	for _, binary_num := range numbers {
		if binary_num == "" {
			continue
		}
		ones = ones + int(binary_num[position]-48)
	}
	if ones >= len(numbers)-ones {
		return 0
	}
	return 1
}

func filter(numbers []string, position int, matchNumber int) []string {
	var filtered_numbers []string
	for _, number := range numbers {
		if number == "" {
			continue
		}
		if int(number[position])-48 == matchNumber {
			filtered_numbers = append(filtered_numbers, number)
		}
	}
	return filtered_numbers
}
