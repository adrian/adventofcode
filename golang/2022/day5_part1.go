package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"regexp"
)

func main() {
	input := utils.ReadFileIntoArrayOfStrings("day5_input")

	// Find the blank line that separates stacks from instructions
	var blank_line int
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			blank_line = i
		}
	}

	// Create stacks
	stack_numbers := input[blank_line-1]
	stacks := []*utils.Stack{}
	for i := 0; i < len(stack_numbers); i++ {
		if stack_numbers[i] == 32 { // ignore spaces
			continue
		}
		stack := utils.Stack{}

		for j := blank_line - 2; j >= 0; j-- {
			item := string(input[j][i])
			if item == " " {
				break
			}
			stack.Push(item)
		}

		stacks = append(stacks, &stack)
	}

	// Carry out instructions
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for i := blank_line + 1; i < len(input); i++ {
		m := re.FindAllStringSubmatch(input[i], -1)
		from_stack := stacks[utils.Atoi(m[0][2])-1]
		to_stack := stacks[utils.Atoi(m[0][3])-1]
		num_to_move := utils.Atoi(m[0][1])
		for j := 0; j < num_to_move; j++ {
			s := from_stack.Pop()
			to_stack.Push(s)
		}
	}

	// Print top item on each stack
	for i := 0; i < len(stacks); i++ {
		fmt.Print((*stacks[i]).Peek())
	}
	fmt.Println()
}
