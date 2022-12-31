package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPart1(t *testing.T) {
	Convey("For the part1", t, func() {
		Convey("For demo input", func() {
			in := input0
			Convey("We get the right answer", func() {
				So(part1(in), ShouldEqual, 21)
			})
		})
		Convey("For real input", func() {
			in := input1
			Convey("We get the right answer", func() {
				So(part1(in), ShouldEqual, 1805)
			})
		})
	})
}

func TestPart2(t *testing.T) {
	Convey("For the part2", t, func() {
		Convey("For demo input", func() {
			in := input0
			Convey("We get the right answer", func() {
				So(part2(in), ShouldEqual, 8)
			})
		})
		Convey("For real input", func() {
			in := input1
			Convey("We get the right answer", func() {
				So(part2(in), ShouldEqual, 444528)
			})
		})
	})
}
