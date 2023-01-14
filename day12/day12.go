package main

import (
	_ "embed"
	"golang.org/x/exp/slices"
	"math"
	"strings"
)

//go:embed input_0.txt
var input0 string

//go:embed input_1.txt
var input1 string

//go:embed input_2.txt
var input2 string

type pos struct {
	x int
	y int
}

func part1(in string) int {
	rows := strings.Split(in, "\n")
	grid := make(map[pos]rune)
	shortestDist := make(map[pos]int)
	var start, end pos
	for y, row := range rows {
		for x, r := range row {
			grid[pos{x, y}] = r
			shortestDist[pos{x, y}] = math.MaxInt
			if r == 'S' {
				start = pos{x, y}
				shortestDist[pos{x, y}] = 0
				grid[pos{x, y}] = 'a'
			}
			if r == 'E' {
				end = pos{x, y}
				grid[pos{x, y}] = 'z'
			}
		}
	}

	toVisit := []pos{start}
	for len(toVisit) > 0 {
		atPos := toVisit[0]
		//log.Printf("visiting pos %v", atPos)
		toVisit = slices.Delete(toVisit, 0, 1)
		movesSoFar := shortestDist[atPos]

		left := pos{atPos.x - 1, atPos.y}
		right := pos{atPos.x + 1, atPos.y}
		up := pos{atPos.x, atPos.y - 1}
		down := pos{atPos.x, atPos.y + 1}
		if canMoveAdjacent(grid, atPos, left) {
			if shortestDist[left] > movesSoFar+1 {
				shortestDist[left] = movesSoFar + 1
				toVisit = append(toVisit, left)
			}
		}
		if canMoveAdjacent(grid, atPos, right) {
			if shortestDist[right] > movesSoFar+1 {
				shortestDist[right] = movesSoFar + 1
				toVisit = append(toVisit, right)
			}
		}
		if canMoveAdjacent(grid, atPos, up) {
			if shortestDist[up] > movesSoFar+1 {
				shortestDist[up] = movesSoFar + 1
				toVisit = append(toVisit, up)
			}
		}
		if canMoveAdjacent(grid, atPos, down) {
			if shortestDist[down] > movesSoFar+1 {
				shortestDist[down] = movesSoFar + 1
				toVisit = append(toVisit, down)
			}
		}
	}
	// breadth-first search, starting at start

	//log.Print(grid)
	//log.Printf("x,y = %v,%v", startX, startY)
	//log.Print(shortestDist)
	return shortestDist[end]
}

func canMoveAdjacent(grid map[pos]rune, from pos, to pos) bool {
	// from and to are assumed to be adjacent

	// check if out of bounds
	if _, valid := grid[to]; !valid {
		return false
	}
	diff := grid[to] - grid[from]
	return diff <= 1
}

func part2(in string) int {
	rows := strings.Split(in, "\n")
	grid := make(map[pos]rune)
	shortestDist := make(map[pos]int)
	var end pos
	var starts []pos
	for y, row := range rows {
		for x, r := range row {
			grid[pos{x, y}] = r
			shortestDist[pos{x, y}] = math.MaxInt
			if r == 'a' || r == 'S' {
				starts = append(starts, pos{x, y})
				shortestDist[pos{x, y}] = 0
				grid[pos{x, y}] = 'a'
			}
			if r == 'E' {
				end = pos{x, y}
				grid[pos{x, y}] = 'z'
			}
		}
	}

	minOverAllStarts := math.MaxInt
	for _, start := range starts {
		toVisit := []pos{start}
		for len(toVisit) > 0 {
			atPos := toVisit[0]
			//log.Printf("visiting pos %v", atPos)
			toVisit = slices.Delete(toVisit, 0, 1)
			movesSoFar := shortestDist[atPos]

			left := pos{atPos.x - 1, atPos.y}
			right := pos{atPos.x + 1, atPos.y}
			up := pos{atPos.x, atPos.y - 1}
			down := pos{atPos.x, atPos.y + 1}
			if canMoveAdjacent(grid, atPos, left) {
				if shortestDist[left] > movesSoFar+1 {
					shortestDist[left] = movesSoFar + 1
					toVisit = append(toVisit, left)
				}
			}
			if canMoveAdjacent(grid, atPos, right) {
				if shortestDist[right] > movesSoFar+1 {
					shortestDist[right] = movesSoFar + 1
					toVisit = append(toVisit, right)
				}
			}
			if canMoveAdjacent(grid, atPos, up) {
				if shortestDist[up] > movesSoFar+1 {
					shortestDist[up] = movesSoFar + 1
					toVisit = append(toVisit, up)
				}
			}
			if canMoveAdjacent(grid, atPos, down) {
				if shortestDist[down] > movesSoFar+1 {
					shortestDist[down] = movesSoFar + 1
					toVisit = append(toVisit, down)
				}
			}
		}
		if shortestDist[end] < minOverAllStarts {
			minOverAllStarts = shortestDist[end]
		}
	}
	// breadth-first search, starting at start

	return minOverAllStarts
}

func main() {

}
