package main

import (
	_ "embed"
	"fmt"
	"golang.org/x/exp/slices"
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

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func part1(in string) string {
	rows := strings.Split(in, "\n")
	// 0th element in each value is top of stack
	state := make(map[int][]byte)
	var rowInd int
	for true {
		row := rows[rowInd]
		if strings.HasPrefix(row, " 1") {
			break
		}
		for stackNum := 1; stackNum*4-3 < len(row); stackNum++ {
			colInd := stackNum*4 - 3
			v := row[colInd]
			if v == ' ' {
				continue
			}
			state[stackNum] = append(state[stackNum], v)
		}
		rowInd++
	}
	rowInd += 2

	// e.g.
	// move 3 from 1 to 3
	for ; rowInd < len(rows); rowInd++ {
		row := rows[rowInd]
		if row == "" {
			break
		}
		spl := strings.Split(row, " ")
		assertLen(spl, 6)
		move := mustInt(spl[1])
		from := mustInt(spl[3])
		to := mustInt(spl[5])
		for i := 0; i < move; i++ {
			toMove := state[from][0]
			state[from] = slices.Delete(state[from], 0, 1)
			state[to] = slices.Insert(state[to], 0, toMove)
		}
	}

	log.Print(state)

	result := ""
	for i := 1; i <= len(state); i++ {
		result = fmt.Sprintf("%s%c", result, state[i][0])
	}

	return result
}

func part2(s string) int64 {
	return -1
}

func main() {

}
