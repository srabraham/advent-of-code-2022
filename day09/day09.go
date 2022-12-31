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

//type visits map[int]map[int]struct{}
//
//func (v *visits) visit(x, y int64){
//	if v == nil {
//		*v = make(map[int]map[int]struct{})
//	}
//	if _, present := v[x]; !present {
//		v[x] = make(map[int]struct{})
//	}
//}

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

	// map y coord -> x coord -> value
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
				//case abs(headX, tailX) == 2 && headY == tailY:
				//	tailX = (headX + tailX) / 2
				//case headX == tailX && abs(headY, tailY) == 2:
				//	tailY = (headY + tailY) / 2
				//case abs(headX, tailX) >= 1 && (abs(headX, tailX)+abs(headY, tailY)) == 3:
				tailX += sign(headX - tailX)
				tailY += sign(headY - tailY)
				//tailY = (headY + tailY + sign(headY-tailY)) / 2
				//default:
				//	log.Fatalf("unexpected case at %v", row)
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

	return -1
}

func main() {

}
