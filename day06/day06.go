package main

import (
	_ "embed"
	"log"
	"strconv"
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

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return i
}

var dummy struct{}

func allDiff(a, b, c, d byte) bool {
	m := make(map[byte]struct{})
	m[a] = dummy
	m[b] = dummy
	m[c] = dummy
	m[d] = dummy
	return len(m) == 4
}

func part1(in string) int {
	for i := range in {
		if i < 3 {
			continue
		}
		if allDiff(in[i-3], in[i-2], in[i-1], in[i]) {
			return i + 1
		}
	}
	log.Panicf("found no four diffs in a row %v", in)
	return -1
}

func part2(s string) int64 {
	return -1
}

func main() {

}
