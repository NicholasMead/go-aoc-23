package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d18/lagoon"
)

var inputPath = "./y23d18/input.txt"
var samplePath = "./y23d18/sample.txt"

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

func part1(input []string) any {
	digger := lagoon.NewDigger(lagoon.Vector{0, 0})

	for _, line := range input {
		instruction := lagoon.InstructionFromInput(line)
		digger.Follow(instruction)
	}

	path := digger.Path()

	return path.Area()
}

func part2(input []string) any {
	digger := lagoon.NewDigger(lagoon.Vector{0, 0})

	for _, line := range input {
		instruction := lagoon.InstructionFromInput(line).Debug()
		digger.Follow(instruction)
	}

	path := digger.Path()

	return path.Area()
}

func getFactors(integer int) (factors []int) {
	for i := 1; i <= integer; i++ {
		if integer%i == 0 {
			factors = append(factors, i)
		}
	}
	return
}

func getCommon(a, b []int) []int {
	common := []int{}

	for _, a := range a {
		if slices.Contains(b, a) {
			common = append(common, a)
		}
	}

	return common
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
