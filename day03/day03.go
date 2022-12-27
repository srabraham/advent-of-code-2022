package main

import (
	_ "embed"
	"strings"
)

//go:embed input_0.txt
var input0 string

//go:embed input_1.txt
var input1 string

//go:embed input_2.txt
var input2 string

func part1(in string) int64 {
	var sumPrio int64
	for _, s := range strings.Split(in, "\n") {
		if s == "" {
			continue
		}
		m1 := make(map[byte]int64)
		m2 := make(map[byte]int64)
		for i := 0; i < len(s)/2; i++ {
			m1[s[i]]++
		}
		for i := len(s) / 2; i < len(s); i++ {
			m2[s[i]]++
		}
		for k := range m1 {
			if m2[k] > 0 {
				if k <= 'z' && k >= 'a' {
					sumPrio += int64(k - 'a' + 1)
				}
				if k <= 'Z' && k >= 'A' {
					sumPrio += int64(k - 'A' + 27)
				}
			}
		}
	}
	return sumPrio
}

func part2(s string) int64 {
	return -1
}

func main() {

}
