package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y15d1/input.txt"
var samplePath = "./y15d1/sample.txt"

func main() {
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

	input := inputFile.ReadInputFile(path)

	fmt.Printf("Part 1: %v\n", part1(input[0]))
	fmt.Printf("Part 2: %v\n", part2(input[0]))
}

func decode(c rune) int {
	switch c {
	case '(':
		return +1
	case ')':
		return -1
	default:
		panic(c)
	}
}

func part1(input string) any {
	floor := 0
	for _, c := range input {
		floor += decode(c)
	}
	return floor
}

func part2(input string) any {
	floor := 0
	for i, c := range input {
		floor += decode(c)

		if floor == -1 {
			return i + 1
		}
	}
	panic("No basement")
}
