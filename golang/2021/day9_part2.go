package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"sort"
)

type heightMap struct {
	lines []string
}

type point struct {
	x, y int
}

func (h *heightMap) getHeightAt(p point) (height int) {
	height = 9
	if p.y >= 0 && p.y < len(h.lines) {
		lineStr := h.lines[p.y]
		if p.x >= 0 && p.x < len(h.lines[0]) {
			height = utils.Atoi(string(lineStr[p.x]))
		}
	}
	return height
}

func (h *heightMap) lowPoints() (lowPoints []point) {
	for y, line := range h.lines {
		for x, height := range line {
			curPosHeight := int(height - 48)
			above := h.getHeightAt(point{x, y - 1})
			below := h.getHeightAt(point{x, y + 1})
			left := h.getHeightAt(point{x - 1, y})
			right := h.getHeightAt(point{x + 1, y})
			if curPosHeight < above && curPosHeight < below && curPosHeight < left && curPosHeight < right {
				lowPoints = append(lowPoints, point{x, y})
			}
		}
	}
	return lowPoints
}

func (h *heightMap) exploreBasin(p point, basin map[point]bool) {
	pointAbove := point{p.x, p.y - 1}
	pointBelow := point{p.x, p.y + 1}
	pointLeft := point{p.x - 1, p.y}
	pointRight := point{p.x + 1, p.y}

	currentPointLevel := h.getHeightAt(point{p.x, p.y})
	levelAbove := h.getHeightAt(pointAbove)
	levelBelow := h.getHeightAt(pointBelow)
	levelLeft := h.getHeightAt(pointLeft)
	levelRight := h.getHeightAt(pointRight)

	if levelAbove != 9 && levelAbove > currentPointLevel {
		if _, explored := basin[pointAbove]; !explored {
			basin[pointAbove] = true
			h.exploreBasin(pointAbove, basin)
		}
	}
	if levelBelow != 9 && levelBelow > currentPointLevel {
		if _, explored := basin[pointBelow]; !explored {
			basin[pointBelow] = true
			h.exploreBasin(pointBelow, basin)
		}
	}
	if levelLeft != 9 && levelLeft > currentPointLevel {
		if _, explored := basin[pointLeft]; !explored {
			basin[pointLeft] = true
			h.exploreBasin(pointLeft, basin)
		}
	}
	if levelRight != 9 && levelRight > currentPointLevel {
		if _, explored := basin[pointRight]; !explored {
			basin[pointRight] = true
			h.exploreBasin(pointRight, basin)
		}
	}
}

func main() {
	heightMap := heightMap{utils.ReadFileIntoArrayOfStrings("day9_input")}

	basinSizes := []int{}
	for _, lowPoint := range heightMap.lowPoints() {
		basin := make(map[point]bool)
		basin[lowPoint] = true
		heightMap.exploreBasin(lowPoint, basin)
		basinSizes = append(basinSizes, len(basin))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])
}
