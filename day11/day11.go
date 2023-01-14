package main

import (
	_ "embed"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input_0.txt
var input0 string

//go:embed input_1.txt
var input1 string

//go:embed input_2.txt
var input2 string

var (
	monkeyRowRE     = regexp.MustCompile("Monkey ([0-9]+):")
	startingItemsRE = regexp.MustCompile("  Starting items: ([0-9,\\s]+)")
	operationRE     = regexp.MustCompile("  Operation: new = (old|[0-9]+) ([+*]) (old|[0-9]+)")
	testRE          = regexp.MustCompile("  Test: divisible by ([0-9]+)")
	ifTrueRE        = regexp.MustCompile("    If true: throw to monkey ([0-9]+)")
	ifFalseRE       = regexp.MustCompile("    If false: throw to monkey ([0-9]+)")
)

func mustInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func mustSlInt(s []string) []int64 {
	var res []int64
	for _, ss := range s {
		res = append(res, mustInt(ss))
	}
	return res
}

func mustLen(s []string, n int) []string {
	if len(s) != n {
		log.Fatalf("expected %v to have length %v, but no", s, n)
	}
	return s
}

func getOnlyElement(s []int64) int64 {
	if len(s) != 1 {
		log.Fatalf("expected 1 element, got %v", s)
	}
	return s[0]
}

type Monkey struct {
	Number     int
	HasItems   []int64
	Operation  func(old int64) int64
	Test       func(n int64) bool
	TestDivBy  int64
	TrueThrow  int
	FalseThrow int

	InspectionCount int64
}

func (m Monkey) String() string {
	return fmt.Sprintf("%v: %v", m.Number, m.HasItems)
}

func part1(in string) int64 {
	rounds := 20

	rows := strings.Split(in, "\n")
	var monkeys []*Monkey
	for i := 0; i < len(rows); i++ {
		if rows[i] == "" {
			log.Printf("done reading %v monkeys", len(monkeys))
			break
		}
		monkeyMatch := mustLen(monkeyRowRE.FindStringSubmatch(rows[i]), 2)
		monkeyNum := mustInt(monkeyMatch[1])
		i++
		startingItemsMatch := mustLen(startingItemsRE.FindStringSubmatch(rows[i]), 2)
		items := mustSlInt(strings.Split(startingItemsMatch[1], ", "))
		i++
		opMatch := mustLen(operationRE.FindStringSubmatch(rows[i]), 4)
		i++
		testMatch := mustLen(testRE.FindStringSubmatch(rows[i]), 2)
		testDivisBy := mustInt(testMatch[1])
		i++
		trueMatch := mustLen(ifTrueRE.FindStringSubmatch(rows[i]), 2)
		trueThrowTo := mustInt(trueMatch[1])
		i++
		falseMatch := mustLen(ifFalseRE.FindStringSubmatch(rows[i]), 2)
		falseThrowTo := mustInt(falseMatch[1])
		i++
		if rows[i] != "" {
			log.Fatalf("expected empty row after %v, got %v", rows[i-1], rows[i])
		}
		monkeys = append(monkeys, &Monkey{
			Number:   int(monkeyNum),
			HasItems: items,
			Operation: func(old int64) int64 {
				left := old
				if opMatch[1] != "old" {
					left = mustInt(opMatch[1])
				}
				right := old
				if opMatch[3] != "old" {
					right = mustInt(opMatch[3])
				}
				switch opMatch[2] {
				case "+":
					return left + right
				case "*":
					return left * right
				default:
					log.Fatalf("bad op %v", opMatch[2])
					return -1
				}
			},
			Test: func(n int64) bool {
				return n%testDivisBy == 0
			},
			TrueThrow:  int(trueThrowTo),
			FalseThrow: int(falseThrowTo),
		})
	}
	for round := 0; round < rounds; round++ {
		for _, monk := range monkeys {
			for len(monk.HasItems) > 0 {
				worryLevel := monk.HasItems[0]
				monk.HasItems = slices.Delete(monk.HasItems, 0, 1)
				//log.Printf("on monkey %v, have hasitems %v", monk.Number, monk.HasItems)
				// inspection begins...
				worryLevel = monk.Operation(worryLevel)
				monk.InspectionCount++
				// relief
				worryLevel = worryLevel / 3
				// test
				if monk.Test(worryLevel) {
					monkeys[monk.TrueThrow].HasItems = append(monkeys[monk.TrueThrow].HasItems, worryLevel)
				} else {
					monkeys[monk.FalseThrow].HasItems = append(monkeys[monk.FalseThrow].HasItems, worryLevel)
				}
			}
		}
	}
	log.Printf("have monkeys %v", monkeys)
	var inspections []int64
	for _, m := range monkeys {
		inspections = append(inspections, m.InspectionCount)
	}
	slices.Sort(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func part2(in string) int64 {
	rounds := 10000

	rows := strings.Split(in, "\n")
	var monkeys []*Monkey
	for i := 0; i < len(rows); i++ {
		if rows[i] == "" {
			log.Printf("done reading %v monkeys", len(monkeys))
			break
		}
		monkeyMatch := mustLen(monkeyRowRE.FindStringSubmatch(rows[i]), 2)
		monkeyNum := mustInt(monkeyMatch[1])
		i++
		startingItemsMatch := mustLen(startingItemsRE.FindStringSubmatch(rows[i]), 2)
		items := mustSlInt(strings.Split(startingItemsMatch[1], ", "))
		i++
		opMatch := mustLen(operationRE.FindStringSubmatch(rows[i]), 4)
		i++
		testMatch := mustLen(testRE.FindStringSubmatch(rows[i]), 2)
		testDivisBy := mustInt(testMatch[1])
		i++
		trueMatch := mustLen(ifTrueRE.FindStringSubmatch(rows[i]), 2)
		trueThrowTo := mustInt(trueMatch[1])
		i++
		falseMatch := mustLen(ifFalseRE.FindStringSubmatch(rows[i]), 2)
		falseThrowTo := mustInt(falseMatch[1])
		i++
		if rows[i] != "" {
			log.Fatalf("expected empty row after %v, got %v", rows[i-1], rows[i])
		}
		monkeys = append(monkeys, &Monkey{
			Number:   int(monkeyNum),
			HasItems: items,
			Operation: func(old int64) int64 {
				left := old
				if opMatch[1] != "old" {
					left = mustInt(opMatch[1])
				}
				right := old
				if opMatch[3] != "old" {
					right = mustInt(opMatch[3])
				}
				switch opMatch[2] {
				case "+":
					return left + right
				case "*":
					return left * right
				default:
					log.Fatalf("bad op %v", opMatch[2])
					return -1
				}
			},
			Test: func(n int64) bool {
				return n%testDivisBy == 0
			},
			TestDivBy:  testDivisBy,
			TrueThrow:  int(trueThrowTo),
			FalseThrow: int(falseThrowTo),
		})
	}
	moduloRelief := int64(1)
	for _, m := range monkeys {
		moduloRelief = moduloRelief * m.TestDivBy
	}
	for round := 0; round < rounds; round++ {
		for _, monk := range monkeys {
			for len(monk.HasItems) > 0 {
				worryLevel := monk.HasItems[0]
				monk.HasItems = slices.Delete(monk.HasItems, 0, 1)
				//log.Printf("on monkey %v, have hasitems %v", monk.Number, monk.HasItems)
				// inspection begins...
				worryLevel = monk.Operation(worryLevel)
				monk.InspectionCount++
				// relief.. part 2 style
				worryLevel = worryLevel % moduloRelief
				// test
				if monk.Test(worryLevel) {
					monkeys[monk.TrueThrow].HasItems = append(monkeys[monk.TrueThrow].HasItems, worryLevel)
				} else {
					monkeys[monk.FalseThrow].HasItems = append(monkeys[monk.FalseThrow].HasItems, worryLevel)
				}
			}
		}
	}
	log.Printf("have monkeys %v", monkeys)
	var inspections []int64
	for _, m := range monkeys {
		inspections = append(inspections, m.InspectionCount)
	}
	slices.Sort(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func main() {

}
