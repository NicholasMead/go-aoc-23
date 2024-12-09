package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y24d02/input.txt"
var samplePath = "./y24d02/sample.txt"

func main() {
	var p1, p2 any = "", ""
	args := os.Args[1:]
	path := inputPath
	if len(args) > 0 {
		switch args[0] {
		case "sample":
			path = samplePath
		case "input":
			path = inputPath
		default:
			path = args[1]
		}
	}

	d := common.Timer(func() {
		input := inputFile.ReadInputFile(path)
		p1, p2 = part1(input), part2(input)
	})

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

type Frame []int

func newFrame(line string) Frame {
	parts := strings.Split(line, " ")

	frame := make(Frame, len(parts))
	for i, part := range parts {
		frame[i], _ = strconv.Atoi(part)
	}
	return frame
}

func (f Frame) isSafe(tollerance int)	bool {
	prev := 0
	errors := 0

	for i := 0; i < len(f)-1; i++ {
		diff := f[i] - f[i+1]

		// Must increase or decrease
		if diff == 0 {
			errors++
			continue
		}

		// Must not increase or decrease by more than 3
		if diff < -3 || diff > 3 {
			errors++
			continue
		}

		// Must not change direction
		if diff * prev < 0 {
			errors++
			continue
		}

		prev = diff
	}

	return errors <= tollerance
}

func part1(input []string) any {
	safe := 0

	for _, line := range input {
		frame := newFrame(line)
		if frame.isSafe(0) {
			safe++
		} else {
		}
	}

	return safe
}

func part2(input []string) any {
	safe := 0

	for _, line := range input {
		frame := newFrame(line)
		if frame.isSafe(1) {
			safe++
		} else {
		}
	}

	return safe
}
