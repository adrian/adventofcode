package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"strconv"
	"os"
	"sort"
)

func main() {
	input := utils.ReadFileIntoArrayOfStrings("day1_input")
	var all_elf_calories []int

	elf_calories := 0
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			all_elf_calories = append(all_elf_calories, elf_calories)
			elf_calories = 0
			continue
		}
		c, err := strconv.Atoi(input[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		elf_calories = elf_calories + c
	}
	sort.Ints(all_elf_calories)
	l := len(all_elf_calories)
	fmt.Println(all_elf_calories[l-1] + all_elf_calories[l-2] + all_elf_calories[l-3])
}
