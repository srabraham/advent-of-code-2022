package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPart1(t *testing.T) {
	Convey("For the demo input", t, func() {
		in := input0
		Convey("We get the right answer", func() {
			So(part1(in), ShouldEqual, -1)
		})
	})
}

func TestPart2(t *testing.T) {
	Convey("For the demo input", t, func() {
		in := input0
		Convey("We get the right answer", func() {
			So(part2(in), ShouldResemble, []int64{-1})
		})
	})
}
