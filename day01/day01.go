package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

//go:embed input_0.txt
var input0 string

//go:embed input_1.txt
var input1 string

//go:embed input_2.txt
var input2 string

func part1(in string) int64 {

	spl := strings.Split(in, "\n")
	var cals []int64
	var curr int64
	for _, s := range spl {
		if s == "" {
			cals = append(cals, curr)
			curr = 0
			continue
		}
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		curr += i
	}
	var max int64
	for _, c := range cals {
		if c > max {
			max = c
		}
	}
	return max
}

func part2(in string) []int64 {
	spl := strings.Split(in, "\n")
	var cals []int64
	var curr int64
	for _, s := range spl {
		if s == "" {
			cals = append(cals, curr)
			curr = 0
			continue
		}
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		curr += i
	}
	slices.Sort(cals)
	return cals
}

func main() {

}
