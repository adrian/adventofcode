package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"log"
)

type stack struct {
	contents []rune
}

func (s *stack) push(c rune) {
	s.contents = append(s.contents, c)
}

func (s *stack) pop() rune {
	lastChar := s.contents[len(s.contents)-1]
	s.contents = s.contents[0 : len(s.contents)-1]
	return lastChar
}

func opening(c rune) bool {
	return c == '(' || c == '{' || c == '[' || c == '<'
}

func closing(c rune) bool {
	return c == ')' || c == '}' || c == ']' || c == '>'
}

func pair(opening rune, closing rune) bool {
	if opening == '(' && closing == ')' {
		return true
	}
	if opening == '{' && closing == '}' {
		return true
	}
	if opening == '[' && closing == ']' {
		return true
	}
	if opening == '<' && closing == '>' {
		return true
	}
	return false
}

func main() {
	input := utils.ReadFileIntoArrayOfStrings("day10_input")

	stack := stack{}
	var illegalChars []rune

	for _, line := range input {
		for _, c := range line {
			if opening(c) {
				stack.push(c)
			} else if closing(c) {
				if !pair(stack.pop(), c) {
					illegalChars = append(illegalChars, c)
					break
				}
			} else {
				log.Fatalf("Invalid character: %c\n", c)
			}
		}
	}

	var score int
	for _, c := range illegalChars {
		if c == ')' {
			score = score + 3
		} else if c == ']' {
			score = score + 57
		} else if c == '}' {
			score = score + 1197
		} else if c == '>' {
			score = score + 25137
		}
	}

	fmt.Println(score)
}
