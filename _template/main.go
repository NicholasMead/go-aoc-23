package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./_template/input.txt"
var samplePath = "./_template/sample.txt"

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

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []string) any {
	return "_"
}

func part2(input []string) any {
	return "_"
}
