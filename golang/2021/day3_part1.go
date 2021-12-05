package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
)

func main() {
	input := utils.ReadFileIntoArrayOfStrings("day3_input")

	frequency_counts := make([]int, len(input[0]))
	for _, binary_num := range input {
		for i := 0; i < len(binary_num); i++ {
			frequency_counts[i] = frequency_counts[i] + int(binary_num[i]-48)
		}
	}

	var gamma_rate_binary_string string
	var epsilon_rate_binary_string string
	for i := 0; i < len(frequency_counts); i++ {
		if frequency_counts[i] > len(input)/2 {
			gamma_rate_binary_string = gamma_rate_binary_string + "1"
			epsilon_rate_binary_string = epsilon_rate_binary_string + "0"
		} else {
			gamma_rate_binary_string = gamma_rate_binary_string + "0"
			epsilon_rate_binary_string = epsilon_rate_binary_string + "1"
		}
	}

	gamma_rate := utils.ConvertBinaryStringToDecimal(gamma_rate_binary_string)
	epsilon_rate := utils.ConvertBinaryStringToDecimal(epsilon_rate_binary_string)

	fmt.Println(gamma_rate * epsilon_rate)
}
