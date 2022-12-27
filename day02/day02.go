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

func part1(in string) int64 {
	spl := strings.Split(in, "\n")
	var score int64
	for _, s := range spl {
		if s == "" {
			continue
		}
		switch s[2] {
		case 'X':
			// 1 point for rock
			score += 1
			switch s[0] {
			case 'A':
				score += 3
			case 'B':
				score += 0
			case 'C':
				score += 6
			default:
				log.Panicf("unexpected s0 %v", s)
			}
		case 'Y':
			// 2 points for paper
			score += 2
			switch s[0] {
			case 'A':
				score += 6
			case 'B':
				score += 3
			case 'C':
				score += 0
			default:
				log.Panicf("unexpected s0 %v", s)
			}
		case 'Z':
			// 3 points for scissors
			score += 3
			switch s[0] {
			case 'A':
				score += 0
			case 'B':
				score += 6
			case 'C':
				score += 3
			default:
				log.Panicf("unexpected s0 %v", s)
			}
		default:
			log.Panicf("unexpected s2 %v", s)
		}
	}
	return score
}

func part2(in string) int64 {
	spl := strings.Split(in, "\n")
	var score int64
	for _, s := range spl {
		if s == "" {
			continue
		}
		switch s[2] {
		case 'X':
			// you must lose
			score += 0
			switch s[0] {
			case 'A':
				score += p2YouPickABCScore('C')
			case 'B':
				score += p2YouPickABCScore('A')
			case 'C':
				score += p2YouPickABCScore('B')
			default:
				log.Panicf("unexpected s0 %v", s)
			}
		case 'Y':
			// you must draw
			score += 3
			switch s[0] {
			case 'A':
				score += p2YouPickABCScore('A')
			case 'B':
				score += p2YouPickABCScore('B')
			case 'C':
				score += p2YouPickABCScore('C')
			default:
				log.Panicf("unexpected s0 %v", s)
			}
		case 'Z':
			// you must win
			score += 6
			switch s[0] {
			case 'A':
				score += p2YouPickABCScore('B')
			case 'B':
				score += p2YouPickABCScore('C')
			case 'C':
				score += p2YouPickABCScore('A')
			default:
				log.Panicf("unexpected s0 %v", s)
			}
		default:
			log.Panicf("unexpected s2 %v", s)
		}
	}
	return score
}

func p2YouPickABCScore(c uint8) int64 {
	switch c {
	case 'A':
		return 1
	case 'B':
		return 2
	case 'C':
		return 3
	default:
		log.Panicf("unexpected c %v", c)
	}
	return -1
}

func main() {

}
