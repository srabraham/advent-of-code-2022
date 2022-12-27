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

func assertLen[T any, S ~[]T](s S, length int) S {
	if len(s) != length {
		log.Panicf("expected length %v got %v", length, s)
	}
	return s
}

func mustInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func part1(in string) int64 {
	var coveredCount int64
	for _, row := range strings.Split(in, "\n") {
		if row == "" {
			continue
		}
		cols := assertLen(strings.Split(row, ","), 2)

		ass1 := assertLen(strings.Split(cols[0], "-"), 2)
		ass1Lower := mustInt(ass1[0])
		ass1Upper := mustInt(ass1[1])

		ass2 := assertLen(strings.Split(cols[1], "-"), 2)
		ass2Lower := mustInt(ass2[0])
		ass2Upper := mustInt(ass2[1])

		if ass2Lower >= ass1Lower && ass2Upper <= ass1Upper {
			coveredCount++
		} else if ass1Lower >= ass2Lower && ass1Upper <= ass2Upper {
			coveredCount++
		}
	}
	return coveredCount
}

func part2(in string) int64 {
	var coveredCount int64
	for _, row := range strings.Split(in, "\n") {
		if row == "" {
			continue
		}
		cols := assertLen(strings.Split(row, ","), 2)

		ass1 := assertLen(strings.Split(cols[0], "-"), 2)
		ass1Lower := mustInt(ass1[0])
		ass1Upper := mustInt(ass1[1])

		ass2 := assertLen(strings.Split(cols[1], "-"), 2)
		ass2Lower := mustInt(ass2[0])
		ass2Upper := mustInt(ass2[1])

		if (ass1Lower <= ass2Lower && ass2Lower <= ass1Upper) ||
			(ass1Lower <= ass2Upper && ass2Upper <= ass1Upper) ||
			(ass2Lower <= ass1Lower && ass1Lower <= ass2Upper) ||
			(ass2Lower <= ass1Upper && ass1Upper <= ass2Upper) {
			coveredCount++
		}
	}
	return coveredCount
}

func main() {

}
