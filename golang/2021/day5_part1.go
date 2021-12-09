package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"regexp"
)

type Point struct {
	x, y int
}

type Line struct {
	start, end Point
}

func pointsForLine(line Line) []Point {
	var points []Point
	if line.start.x == line.end.x {
		var start, end int
		if line.start.y > line.end.y {
			start = line.end.y
			end = line.start.y
		} else {
			start = line.start.y
			end = line.end.y
		}
		for y := start; y <= end; y++ {
			points = append(points, Point{line.start.x, y})
		}
	} else if line.start.y == line.end.y {
		var start, end int
		if line.start.x > line.end.x {
			start = line.end.x
			end = line.start.x
		} else {
			start = line.start.x
			end = line.end.x
		}
		for x := start; x <= end; x++ {
			points = append(points, Point{x, line.start.y})
		}
	} else {
		// ignore
	}
	return points
}

func addPoints(line Line, pointCounts map[Point]int) map[Point]int {
	for _, point := range pointsForLine(line) {
		count, _ := pointCounts[point]
		pointCounts[point] = count + 1
	}
	return pointCounts
}

func main() {
	lines := []Line{}
	for _, line := range utils.ReadFileIntoArrayOfStrings("day5_input") {
		if line == "" {
			continue
		}
		re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
		matches := re.FindStringSubmatch(line)
		point1 := Point{utils.Atoi(matches[1]), utils.Atoi(matches[2])}
		point2 := Point{utils.Atoi(matches[3]), utils.Atoi(matches[4])}
		lines = append(lines, Line{point1, point2})
	}

	var pointCounts = make(map[Point]int)
	for _, line := range lines {
		pointCounts = addPoints(line, pointCounts)
	}

	var answer int
	for _, count := range pointCounts {
		if count > 1 {
			answer = answer + 1
		}
	}
	fmt.Println(answer)
}
