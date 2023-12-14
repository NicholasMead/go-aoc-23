package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d14/platform"
)

var inputPath = "./y23d14/input.txt"
var samplePath = "./y23d14/sample.txt"

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
	fmt.Printf("Time: %vs\n", d.Seconds())
}

func part1(input []string) any {
	p := platform.LoadPlatform(input)
	p.ApplyTilt(platform.Vector{0, -1})
	return p.NorthWeight()
}

func part2(input []string) any {
	generator := NewSpinCycleHashGenerator(platform.LoadPlatform(input))

	goal := 1_000_000_000
	offset, cycle := DetectCycle(generator, goal)
	relativePositionAtGoal := offset + ((goal - offset) % cycle)

	p := platform.LoadPlatformFromHash(generator(relativePositionAtGoal))
	return p.NorthWeight()
}
