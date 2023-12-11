package main

import (
	"strings"
)

const expanseFactor = 1000000 - 1 // change to 1 for part 1 answer

type point struct {
	x, y int
}

func (p point) dist(o point) int {
	return abs(p.x-o.x) + abs(p.y-o.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	lines := strings.Split(input, "\n")
	var grid [][]byte
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}
	var galaxies []point
	markRows := make([]int, len(grid)) // square matrix
	markCols := make([]int, len(grid))
	for i := 0; i < len(grid); i++ {
		markRows[i] = expanseFactor
		markCols[i] = expanseFactor
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '#' {
				galaxies = append(galaxies, point{x, y})
				markRows[y] = 0
				markCols[x] = 0
			}
		}
	}
	rowDeltas := make([]int, len(grid))
	colDeltas := make([]int, len(grid))
	exCount := 0
	for idx, expand := range markRows {
		exCount += expand
		rowDeltas[idx] = exCount
	}
	exCount = 0
	for idx, expand := range markCols {
		exCount += expand
		colDeltas[idx] = exCount
	}
	for i := 0; i < len(galaxies); i++ {
		galaxies[i].x += colDeltas[galaxies[i].x]
		galaxies[i].y += rowDeltas[galaxies[i].y]
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += galaxies[i].dist(galaxies[j])
		}
	}
	println(sum)
}
