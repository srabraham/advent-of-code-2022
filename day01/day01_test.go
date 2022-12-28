package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPart1(t *testing.T) {
	Convey("For the demo input", t, func() {
		in := input0
		Convey("We get the right answer", func() {
			So(part1(in), ShouldEqual, 24000)
		})
	})
	Convey("For the real input", t, func() {
		in := input1
		Convey("We get the right answer", func() {
			So(part1(in), ShouldEqual, 72478)
		})
	})
}

func TestPart2(t *testing.T) {
	Convey("For the demo input", t, func() {
		in := input0
		Convey("We get the right answer", func() {
			sorted := part2(in)
			ans := sorted[len(sorted)-3] + sorted[len(sorted)-2] + sorted[len(sorted)-1]
			So(ans, ShouldEqual, 45000)
		})
	})
	Convey("For the real input", t, func() {
		in := input1
		Convey("We get the right answer", func() {
			sorted := part2(in)
			ans := sorted[len(sorted)-3] + sorted[len(sorted)-2] + sorted[len(sorted)-1]
			So(ans, ShouldEqual, 210367)
		})
	})
}
