package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y23d4/input.txt"
var samplePath = "./y23d4/sample.txt"

func main() {
	var p1, p2 any = "", ""
	d := common.Timer(func() {
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

		p1, p2 = part1(input), part2(input)
	})
	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

func part1(input []string) any {
	ans := 0

	for _, line := range input {
		card := ParseScratchCard(line)
		ans += card.Score()
	}

	return ans
}

func part2(input []string) (ans int) {
	copies := map[int]int{}

	for _, line := range input {
		card := ParseScratchCard(line)
		id := card.id
		matches := card.Matches()

		if _, ok := copies[id]; !ok {
			copies[id] = 1
		} else {
			copies[id]++
		}

		for i := 0; i < matches; i++ {
			copies[id+i+1] += copies[id]
		}

		ans += copies[id]
	}

	return
}
