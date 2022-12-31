package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input_0.txt
var input0 string

//go:embed input_1.txt
var input1 string

//go:embed input_2.txt
var input2 string

func mustInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func abs(a, b int64) int64 {
	if a > b {
		return a - b
	}
	return b - a
}

func sign(a int64) int64 {
	if a > 0 {
		return 1
	}
	if a < 0 {
		return -1
	}
	return 0
}

func part1(in string) int64 {

	tailVisits := make(map[int64]map[int64]struct{})
	visited := struct{}{}

	var headX, headY, tailX, tailY int64
	for _, row := range strings.Split(in, "\n") {
		if row == "" {
			continue
		}
		splits := strings.Split(row, " ")
		if len(splits) != 2 {
			log.Fatalf("wanted two splits, got %v", splits)
		}
		dir := splits[0]
		steps := mustInt(splits[1])
		for i := int64(0); i < steps; i++ {
			// move the head one spot
			switch dir {
			case "L":
				headX--
			case "R":
				headX++
			case "U":
				headY++
			case "D":
				headY--
			default:
				log.Fatalf("unexpected direction %v", dir)
			}
			// figure out if the tail needs to move, and if so, move it
			switch {
			case abs(headX, tailX) <= 1 && abs(headY, tailY) <= 1:
				// do nothing
				break
			default:
				tailX += sign(headX - tailX)
				tailY += sign(headY - tailY)
			}
			//log.Printf("got head at (%v,%v) and tail at (%v,%v)", headX, headY, tailX, tailY)
			if _, found := tailVisits[tailX]; !found {
				tailVisits[tailX] = make(map[int64]struct{})
			}
			tailVisits[tailX][tailY] = visited
		}
	}
	var count int64
	for _, m := range tailVisits {
		for range m {
			count++
		}
	}
	return count
}

func part2(in string) int64 {

	tailVisits := make(map[int64]map[int64]struct{})
	visited := struct{}{}

	ropeLen := 10

	xPos := make([]int64, ropeLen)
	yPos := make([]int64, ropeLen)
	for _, row := range strings.Split(in, "\n") {
		if row == "" {
			continue
		}
		splits := strings.Split(row, " ")
		if len(splits) != 2 {
			log.Fatalf("wanted two splits, got %v", splits)
		}
		dir := splits[0]
		steps := mustInt(splits[1])
		for i := int64(0); i < steps; i++ {
			// move the head one spot
			switch dir {
			case "L":
				xPos[0]--
			case "R":
				xPos[0]++
			case "U":
				yPos[0]++
			case "D":
				yPos[0]--
			default:
				log.Fatalf("unexpected direction %v", dir)
			}
			for j := 1; j < ropeLen; j++ {
				// figure out if the tail needs to move, and if so, move it
				switch {
				case abs(xPos[j-1], xPos[j]) <= 1 && abs(yPos[j-1], yPos[j]) <= 1:
					// do nothing
					break
				default:
					// if we need to move, we move 1 unit in each direction toward the head
					xPos[j] += sign(xPos[j-1] - xPos[j])
					yPos[j] += sign(yPos[j-1] - yPos[j])
				}
			}
			if _, found := tailVisits[xPos[ropeLen-1]]; !found {
				tailVisits[xPos[ropeLen-1]] = make(map[int64]struct{})
			}
			tailVisits[xPos[ropeLen-1]][yPos[ropeLen-1]] = visited
		}
	}
	var count int64
	for _, m := range tailVisits {
		for range m {
			count++
		}
	}
	return count
}

func main() {

}
