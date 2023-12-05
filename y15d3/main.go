package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y15d3/input.txt"
var samplePath = "./y15d3/sample.txt"

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

		input := inputFile.ReadInputFile(path)[0]

		p1, p2 = part1(input), part2(input)
	})
	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

type coord [2]int

func (c coord) Move(d rune) coord {
	switch d {
	case '^':
		return coord{c[0], c[1] + 1}
	case 'v':
		return coord{c[0], c[1] - 1}
	case '>':
		return coord{c[0] + 1, c[1]}
	case '<':
		return coord{c[0] - 1, c[1]}
	default:
		panic(d)
	}
}

func part1(input string) any {
	houses := map[coord]int{
		{}: 1,
	}
	santa := coord{}

	for _, d := range input {
		santa = santa.Move(d)
		houses[santa] += 1
	}

	return len(houses)
}

func part2(input string) any {
	houses := map[coord]int{
		{}: 2,
	}
	santa := coord{}
	robot := coord{}

	for i, d := range input {
		if i%2 == 0 {
			santa = santa.Move(d)
			houses[santa] += 1
		} else {
			robot = robot.Move(d)
			houses[robot] += 1
		}
	}

	return len(houses)
}
