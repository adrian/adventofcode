package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"strings"
)

func main() {
	text := utils.ReadFileIntoArrayOfStrings("day8_input")

	count := 0

	for _, text := range text {
		if text == "" {
			continue
		}
		outputValues := strings.Split(strings.Split(text, "|")[1], " ")

		for _, value := range outputValues {
			if len(value) == 2 || len(value) == 4 || len(value) == 3 || len(value) == 7 {
				count = count + 1
			}
		}
	}

	fmt.Println(count)
}
