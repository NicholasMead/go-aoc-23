package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

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

type box [3]int

func (b box) min() (min int) {
	m := math.MaxInt

	for _, v := range b {
		if v < m {
			m = v
		}
	}

	return m
}

func parse(dim string) (out box) {
	for i := range [2]struct{}{} {
		index := strings.IndexRune(dim, 'x')
		out[i], _ = strconv.Atoi(dim[:index])
		dim = dim[index+1:]
	}
	out[2], _ = strconv.Atoi(dim)
	return
}

func part1(input []string) any {
	total := 0
	for _, line := range input {
		box := parse(line)

		area := 2 * (box[0]*box[1] + box[1]*box[2] + box[2]*box[0])
		total += area
		total += box.min()
	}
	return total
}

func part2(input []string) any {
	return "_"
}
