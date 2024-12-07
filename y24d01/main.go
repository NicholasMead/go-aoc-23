package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y24d01/input.txt"
var samplePath = "./y24d01/sample.txt"

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

func parseLine(line string) (left, right int) {
	for strings.Contains(line, "  ") {
		line = strings.ReplaceAll(line, "  ", " ")
	}

	parts := strings.Split(line, " ")

	leftStr := parts[0]
	rightStr := parts[1]

 	left, _ = strconv.Atoi(leftStr)
	right, _ = strconv.Atoi(rightStr)
	return
}

func parseLines(lines []string) (lefts, rights []int) {
	lefts = make([]int, len(lines))
	rights = make([]int, len(lines))

	for i, line := range lines {
		lefts[i], rights[i] = parseLine(line)
	}

	return
}

func part1(input []string) any {
	lefts, rights := parseLines(input)

	slices.Sort(lefts)
	slices.Sort(rights)

	diff := 0
	for i := 0; i < len(lefts); i++ {
		if lefts[i] > rights[i] {
			diff += lefts[i] - rights[i]
		} else {
			diff += rights[i] - lefts[i]
		}
	}

	return diff
}

func part2(input []string) any {
	lefts, rights := parseLines(input)
	answer := 0

	for l := 0; l < len(lefts); l++ {
		count := 0
		for r := 0; r < len(rights); r++ {
			if lefts[l] == rights[r] {
				count++
			}
		}
		answer += lefts[l] * count
	}

	return answer
}
