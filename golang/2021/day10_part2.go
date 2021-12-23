package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"log"
	"sort"
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

func closingFor(opening rune) rune {
	if opening == '(' {
		return ')'
	} else if opening == '{' {
		return '}'
	} else if opening == '[' {
		return ']'
	}
	return '>'
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

func calcScore(completionString []rune) int {
	score := 0
	for _, c := range completionString {
		score = score * 5
		if c == ')' {
			score = score + 1
		} else if c == ']' {
			score = score + 2
		} else if c == '}' {
			score = score + 3
		} else if c == '>' {
			score = score + 4
		}
	}
	return score
}

func main() {
	input := utils.ReadFileIntoArrayOfStrings("day10_input")

	stack := stack{}
	completionStrings := [][]rune{}
	scores := []int{}

	for _, line := range input {
		for _, c := range line {
			if opening(c) {
				stack.push(c)
			} else if closing(c) {
				if !pair(stack.pop(), c) {
					stack.contents = []rune{}
					break
				}
			} else {
				log.Fatalf("Invalid character: %c\n", c)
			}
		}
		if len(stack.contents) == 0 {
			continue
		}

		closingChars := []rune{}
		for len(stack.contents) > 0 {
			closingChar := closingFor(stack.pop())
			closingChars = append(closingChars, closingChar)
		}

		completionStrings = append(completionStrings, closingChars)
	}

	for _, completionString := range completionStrings {
		score := calcScore(completionString)
		scores = append(scores, score)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(scores)))
	fmt.Println(scores[(len(scores)-1)/2])
}
