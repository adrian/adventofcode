package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadFileIntoArrayOfStrings("day2_input")
	horizontal_position := 0
	depth := 0
	for _, command := range input {
		if command == "" {
			continue
		}
		direction, units := parseCommand(command)
		if direction == "forward" {
			horizontal_position = horizontal_position + units
		} else if direction == "down" {
			depth = depth + units
		} else if direction == "up" {
			depth = depth - units
		} else {
			fmt.Println("Invalid command:", direction)
		}
	}
	fmt.Println(horizontal_position * depth)
}

func parseCommand(command string) (string, int) {
	parts := strings.Split(command, " ")
	units, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return parts[0], units
}
