package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d02/elves"
)

var inputPath = "./y23d02/input.txt"
var samplePath = "./y23d02/sample.txt"

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
	bag := elves.Bag{Red: 12, Green: 13, Blue: 14}
	ans := 0

	for _, line := range input {
		game := elves.ParseGame(line)
		if game.CanPlay(bag) {
			ans += game.Id
		}
	}

	return ans
}

func part2(input []string) any {
	ans := 0

	for _, line := range input {
		game := elves.ParseGame(line)
		ans += game.MinimumPlayableBag().Power()
	}

	return ans
}
