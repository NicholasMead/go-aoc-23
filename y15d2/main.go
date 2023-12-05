package main

import (
	"fmt"
	"math"
	"os"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y15d2/input.txt"
var samplePath = "./y15d2/sample.txt"

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

func min(values ...int) int {
	m := math.MaxInt
	for _, v := range values {
		if v < m {
			m = v
		}
	}
	return m
}

func sum(values ...int) (s int) {
	for _, v := range values {
		s += v
	}
	return
}

func part1(input []string) any {
	total := 0
	for _, line := range input {
		sides := parse(line).Sides()
		area := 0
		slack := math.MaxInt

		for _, s := range sides {
			area += 2 * s.Area()
			if s.Area() < slack {
				slack = s.Area()
			}
		}

		total += area + slack
	}
	return total
}

func part2(input []string) any {
	total := 0
	for _, line := range input {
		box := parse(line)
		sides := box.Sides()
		ribbon := math.MaxInt

		for _, s := range sides {
			if s.Perimeter() < ribbon {
				ribbon = s.Perimeter()
			}
		}

		total += ribbon + box.Volume()
	}
	return total
}
