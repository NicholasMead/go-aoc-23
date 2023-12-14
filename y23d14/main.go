package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d14/platform"
)

var inputPath = "./y23d14/input.txt"
var samplePath = "./y23d14/sample.txt"

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
	height, width := len(input), len(input[0])
	p := platform.NewPlatform(height, width)

	for y, line := range input {
		for x, char := range line {
			if !strings.ContainsRune(platform.RockTypes, char) {
				continue
			}

			p.AddRock(platform.Vector{x, y}, platform.RockType(char))
		}
	}
	fmt.Println(p)

	p.ApplyTilt(platform.Vector{0, -1})

	fmt.Println(p)

	ans := 0
	for _, rock := range p.GetRocksOfType(platform.RoundRock) {
		ans += height - rock[1]
	}
	return ans
}

func part2(input []string) any {
	height, width := len(input), len(input[0])
	p := platform.NewPlatform(height, width)

	for y, line := range input {
		for x, char := range line {
			if !strings.ContainsRune(platform.RockTypes, char) {
				continue
			}

			p.AddRock(platform.Vector{x, y}, platform.RockType(char))
		}
	}

	max := 3 //1_000_000_000
	for i := 0; i < max; i++ {
		p.ApplyTilt(platform.Vector{0, -1})
		p.ApplyTilt(platform.Vector{-1, 0})
		p.ApplyTilt(platform.Vector{0, 1})
		p.ApplyTilt(platform.Vector{1, 0})
	}
	fmt.Println(p)

	ans := 0
	for _, rock := range p.GetRocksOfType(platform.RoundRock) {
		ans += height - rock[1]
	}
	return ans
}
