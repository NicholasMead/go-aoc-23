package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y23d2/input.txt"
var samplePath = "./y23d2/sample.txt"

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
	games := []game{}

	for _, line := range input {
		g := parseGame(line)
		games = append(games, g)
	}

	fmt.Printf("Part 1: %v\n", part1(games))
	fmt.Printf("Part 2: %v\n", part2(games))
}

func part1(input []game) any {
	cap := cubeSet{12, 13, 14}
	ans := 0

	for _, game := range input {
		success := true
		for _, round := range game.rounds {
			if round.red > cap.red || round.green > cap.green || round.blue > cap.blue {
				success = false
				break
			}
		}
		if success {
			ans += game.num
		}
	}

	return ans
}

func part2(input []game) any {
	ans := 0

	for _, game := range input {
		min := cubeSet{}

		for _, round := range game.rounds {
			if round.red > min.red {
				min.red = round.red
			}
			if round.green > min.green {
				min.green = round.green
			}
			if round.blue > min.blue {
				min.blue = round.blue
			}
		}

		pow := min.red * min.green * min.blue
		ans += pow
	}

	return ans
}
