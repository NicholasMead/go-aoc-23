package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d12/springs"
)

var inputPath = "./y23d12/input.txt"
var samplePath = "./y23d12/sample.txt"

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
	ans := 0

	for _, line := range input {
		lineParts := strings.Split(line, " ")

		history := springs.History(lineParts[0])
		record := springs.RecordFromString(lineParts[1])

		count := springs.CountValid(history, record)
		ans += count
	}

	return ans
}

func part2(input []string) any {
	ans := 0

	for _, line := range input {
		lineParts := strings.Split(line, " ")

		history := springs.History(lineParts[0])
		record := springs.RecordFromString(lineParts[1])

		history, record = springs.Unfold(history, record)

		count := springs.CountValid(history, record)
		ans += count
	}

	return ans
}
