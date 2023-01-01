package main

import (
	_ "embed"
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

func mustInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func nextCommand(program string) func() string {
	commands := strings.Split(program, "\n")
	i := -1
	return func() string {
		i++
		if i >= len(commands) {
			return ""
		}
		return commands[i]
	}
}

func part1(in string) int64 {
	nextCmd := nextCommand(in)

	xReg := int64(1)
	running := false
	cmd := "noop"
	cmdDoneAfterCycle := int64(0)
	signalStrength := int64(0)

	for cycle := int64(1); ; cycle++ {
		// start of the cycle...if nothing is running, read a new command
		if !running {
			cmd = nextCmd()
			if cmd == "" {
				log.Printf("no commands remaining...breaking")
				break
			}
			running = true
			switch {
			case cmd == "noop":
				cmdDoneAfterCycle = cycle
			case strings.HasPrefix(cmd, "addx "):
				cmdDoneAfterCycle = cycle + 1
			default:
				log.Fatalf("bad command %v", cmd)
			}
		}

		// in middle of cycle
		if cycle%40 == 20 {
			log.Printf("during cycle %v, command %v, xReg %v", cycle, cmd, xReg)
			signalStrength += cycle * xReg
		}

		// end of cycle...if command is done, let it finish
		if cycle == cmdDoneAfterCycle {
			running = false
			switch {
			case cmd == "noop":
				break
			case strings.HasPrefix(cmd, "addx "):
				xReg += mustInt(strings.TrimPrefix(cmd, "addx "))
			default:
				log.Fatalf("unexpected cmd %v", cmd)
			}
		}

	}
	return signalStrength
}
