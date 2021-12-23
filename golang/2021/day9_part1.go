package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
)

type heightMap struct {
	lines []string
}

func (h *heightMap) getHeightAt(line int, index int) (height int) {
	height = 9
	if line >= 0 && line < len(h.lines) {
		lineStr := h.lines[line]
		if index >= 0 && index < len(h.lines[0]) {
			height = utils.Atoi(string(lineStr[index]))
		}
	}
	return height
}

func main() {
	heightMap := heightMap{utils.ReadFileIntoArrayOfStrings("day9_input")}

	var lowPoints []int
	for i, line := range heightMap.lines {
		for j, height := range line {
			curPosHeight := int(height - 48)
			above := heightMap.getHeightAt(i-1, j)
			below := heightMap.getHeightAt(i+1, j)
			left := heightMap.getHeightAt(i, j-1)
			right := heightMap.getHeightAt(i, j+1)
			if curPosHeight < above && curPosHeight < below && curPosHeight < left && curPosHeight < right {
				lowPoints = append(lowPoints, curPosHeight)
			}
		}
	}

	var sumRiskLevels int
	for _, lowPoint := range lowPoints {
		sumRiskLevels = sumRiskLevels + lowPoint + 1
	}

	fmt.Println(sumRiskLevels)
}
