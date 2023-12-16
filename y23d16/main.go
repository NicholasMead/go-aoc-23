package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d16/volcano"
)

var inputPath = "./y23d16/input.txt"
var samplePath = "./y23d16/sample.txt"

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
	cave := volcano.NewCave(len(input), len(input[0]))

	for y, line := range input {
		for x, space := range line {
			if space == '.' {
				continue
			}
			cave.AddMirror(
				volcano.Vector{x, y},
				volcano.NewMirrorFromRune(space),
			)
		}
	}

	beamStart := volcano.Beam{
		Position:  volcano.Vector{-1, 0},
		Direction: volcano.Vector{1, 0},
	}
	enriched := cave.EnrichedFromBeam(beamStart)

	return len(enriched)
}

func part2(input []string) any {
	height, width := len(input), len(input[0])
	cave := volcano.NewCave(height, width)

	for y, line := range input {
		for x, space := range line {
			if space == '.' {
				continue
			}
			cave.AddMirror(
				volcano.Vector{x, y},
				volcano.NewMirrorFromRune(space),
			)
		}
	}

	beams := []volcano.Beam{}
	for y := 0; y < height; y++ {
		beams = append(beams,
			volcano.Beam{
				Position:  volcano.Vector{-1, y},
				Direction: volcano.Vector{1, 0},
			},
			volcano.Beam{
				Position:  volcano.Vector{width, y},
				Direction: volcano.Vector{-1, 0},
			},
		)
	}
	for x := 0; x < height; x++ {
		beams = append(beams,
			volcano.Beam{
				Position:  volcano.Vector{x, -1},
				Direction: volcano.Vector{0, 1},
			}, volcano.Beam{
				Position:  volcano.Vector{x, height},
				Direction: volcano.Vector{0, -1},
			},
		)
	}

	max := 0
	for _, beam := range beams {
		enriched := len(cave.EnrichedFromBeam(beam))

		if enriched > max {
			max = enriched
		}
	}

	return max
}
