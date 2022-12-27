package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPart1(t *testing.T) {
	Convey("For the demo input", t, func() {
		in := input0
		Convey("We get the right answer", t, func() {
			So(part1(in), ShouldEqual, "")
		})
	})
}

func TestPart2(t *testing.T) {
	Convey("For the demo input", t, func() {
		in := input0
		Convey("We get the right answer", t, func() {
			So(part2(in), ShouldEqual, "")
		})
	})
}
