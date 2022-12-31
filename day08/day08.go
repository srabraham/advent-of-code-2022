package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input_0.txt
var input0 string

//go:embed input_1.txt
var input1 string

//go:embed input_2.txt
var input2 string

func assertLen[T any, S ~[]T](s S, length int) S {
	if len(s) != length {
		log.Panicf("expected length %v got %v", length, s)
	}
	return s
}

func leftVisi(grid map[int]map[int]int, x, y int) bool {
	for x0 := 0; x0 < x; x0++ {
		if grid[y][x0] >= grid[y][x] {
			return false
		}
	}
	return true
}

func rightVisi(grid map[int]map[int]int, x, y int) bool {
	for x0 := len(grid[y]) - 1; x0 > x; x0-- {
		if grid[y][x0] >= grid[y][x] {
			return false
		}
	}
	return true
}

func upVisi(grid map[int]map[int]int, x, y int) bool {
	for y0 := 0; y0 < y; y0++ {
		if grid[y0][x] >= grid[y][x] {
			return false
		}
	}
	return true
}

func downVisi(grid map[int]map[int]int, x, y int) bool {
	for y0 := len(grid) - 1; y0 > y; y0-- {
		if grid[y0][x] >= grid[y][x] {
			return false
		}
	}
	return true
}

func part1(in string) int64 {

	// map y coord -> x coord -> value
	grid := make(map[int]map[int]int)

	for y, row := range strings.Split(in, "\n") {
		if row == "" {
			break
		}
		newRow := make(map[int]int)
		for x, b := range []byte(row) {
			newRow[x] = int(b - '0')
		}
		grid[y] = newRow
	}

	visiCount := int64(0)
	for y := range grid {
		for x := range grid[y] {
			if leftVisi(grid, x, y) || rightVisi(grid, x, y) || upVisi(grid, x, y) || downVisi(grid, x, y) {
				visiCount++
			}
		}
	}

	return visiCount
}

func leftVisiDist(grid map[int]map[int]int, x, y int) int64 {
	trees := int64(0)
	for x0 := x - 1; x0 >= 0; x0-- {
		trees++
		if grid[y][x0] >= grid[y][x] {
			break
		}
	}
	return trees
}

func rightVisiDist(grid map[int]map[int]int, x, y int) int64 {
	trees := int64(0)
	for x0 := x + 1; x0 < len(grid[y]); x0++ {
		trees++
		if grid[y][x0] >= grid[y][x] {
			break
		}
	}
	return trees
}

func upVisiDist(grid map[int]map[int]int, x, y int) int64 {
	trees := int64(0)
	for y0 := y - 1; y0 >= 0; y0-- {
		trees++
		if grid[y0][x] >= grid[y][x] {
			break
		}
	}
	return trees
}

func downVisiDist(grid map[int]map[int]int, x, y int) int64 {
	trees := int64(0)
	for y0 := y + 1; y0 < len(grid); y0++ {
		trees++
		if grid[y0][x] >= grid[y][x] {
			break
		}
	}
	return trees
}

func part2(in string) int64 {

	// map y coord -> x coord -> value
	grid := make(map[int]map[int]int)

	for y, row := range strings.Split(in, "\n") {
		if row == "" {
			break
		}
		newRow := make(map[int]int)
		for x, b := range []byte(row) {
			newRow[x] = int(b - '0')
		}
		grid[y] = newRow
	}

	maxTreesVis := int64(0)
	for y := range grid {
		for x := range grid[y] {
			treesVis := leftVisiDist(grid, x, y) * rightVisiDist(grid, x, y) * upVisiDist(grid, x, y) * downVisiDist(grid, x, y)
			if treesVis > maxTreesVis {
				log.Printf("found max at y,x %v,%v with %v", y, x, grid[y][x])
				maxTreesVis = treesVis
			}
		}
	}

	return maxTreesVis
}

func main() {

}
