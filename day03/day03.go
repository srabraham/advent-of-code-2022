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

func part2(in string) int64 {
	var sumPrio int64
	rows := strings.Split(in, "\n")
	for i := 0; i < len(rows); i += 3 {
		if rows[i] == "" {
			break
		}
		m0 := make(map[byte]int64)
		m1 := make(map[byte]int64)
		m2 := make(map[byte]int64)
		for j := 0; j < len(rows[i]); j++ {
			m0[rows[i][j]]++
		}
		for j := 0; j < len(rows[i+1]); j++ {
			m1[rows[i+1][j]]++
		}
		for j := 0; j < len(rows[i+2]); j++ {
			m2[rows[i+2][j]]++
		}
		for k := range m0 {
			if m1[k] > 0 && m2[k] > 0 {
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

func main() {

}
