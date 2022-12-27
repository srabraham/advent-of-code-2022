package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPart1(t *testing.T) {
	Convey("For the part1", t, func() {
		Convey("For demo input", func() {
			in := input0
			Convey("We get the right answer", func() {
				So(part1(in), ShouldEqual, 157)
			})
		})
		Convey("For real input", func() {
			in := input1
			Convey("We get the right answer", func() {
				So(part1(in), ShouldEqual, 7980)
			})
		})
	})
}

func TestPart2(t *testing.T) {
	Convey("For the part2", t, func() {
		Convey("For demo input", func() {
			in := input0
			Convey("We get the right answer", func() {
				So(part2(in), ShouldEqual, 70)
			})
		})
		Convey("For real input", func() {
			in := input1
			Convey("We get the right answer", func() {
				So(part2(in), ShouldEqual, 2881)
			})
		})
	})
}
