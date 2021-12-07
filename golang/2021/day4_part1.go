package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"strings"
)

type Cell struct {
	number int
	marked bool
}

type Board struct {
	numbers [][]*Cell
}

func (b *Board) AddRow(text string) {
	numbers := strings.Fields(text)
	row := make([]*Cell, len(numbers))
	for i, number := range numbers {
		row[i] = &Cell{utils.Atoi(number), false}
	}
	b.numbers = append(b.numbers, row)
}

func (b *Board) Mark(numberDrawn int) {
outside:
	for _, row := range b.numbers {
		for _, cell := range row {
			if cell.number == numberDrawn {
				cell.marked = true
				break outside
			}
		}
	}
}

func (b Board) HasWin() bool {
	for _, row := range b.numbers {
		allRowMarked := true
		for _, cell := range row {
			if cell.marked == false {
				allRowMarked = false
				break
			}
		}
		if allRowMarked {
			return true
		}
	}

	for columnNumber := 0; columnNumber < len(b.numbers[0]); columnNumber++ {
		allColMarked := true
		for _, row := range b.numbers {
			if row[columnNumber].marked == false {
				allColMarked = false
				break
			}
		}
		if allColMarked {
			return true
		}
	}

	return false
}

func (b *Board) SumUnMarked() int {
	sum := 0
	for _, row := range b.numbers {
		for _, cell := range row {
			if !cell.marked {
				sum = sum + cell.number
			}
		}
	}
	return sum
}

func (b Board) String() string {
	var s string
	for _, row := range b.numbers {
		for _, cell := range row {
			s = s + fmt.Sprintf("{%v, %v}", cell.number, cell.marked)
		}
		s = s + fmt.Sprintf("\n")
	}
	return s
}

func main() {
	text := utils.ReadFile("day4_input")

	randomNumbers := utils.StringArrayToIntArray(strings.Split(text[0], ","))

	boards := []Board{}
	currentBoard := Board{}
	for i := 2; i < len(text); i++ {
		if text[i] != "" {
			currentBoard.AddRow(text[i])
		}
		if i == len(text)-1 || text[i+1] == "" {
			boards = append(boards, currentBoard)
			currentBoard = Board{}
		}
	}

outside:
	for _, numberDrawn := range randomNumbers {
		for _, board := range boards {
			board.Mark(numberDrawn)
			if board.HasWin() {
				fmt.Println("This board has a win:")
				fmt.Println(board)
				fmt.Println("Answer:", board.SumUnMarked()*numberDrawn)

				break outside
			}
		}
	}
}
