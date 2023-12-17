package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d17/crucible"
)

var inputPath = "./y23d17/input.txt"
var samplePath = "./y23d17/sample.txt"

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
		grid := NewGrid(input)

		p1, p2 = part1(grid), part2(grid)
	})

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vs\n", d.Seconds())
}

func NewGrid(input []string) crucible.Grid {
	grid := crucible.Grid{}

	for y, line := range input {
		for x, a := range line {
			pos := crucible.Vector{x, y}
			grid[pos] = int(a - '0')
		}
	}

	return grid
}

type Path struct {
	Crucible crucible.Crucible
	From     *Path
	Cost     int
	Weight   int
}

func FindOptimumPath(grid crucible.Grid, start crucible.Crucible) Path {
	var (
		target = grid.Max()
		queue  = []*Path{{Crucible: start, From: nil}}
		memo   = map[crucible.Crucible]int{}

		current *Path = nil
		final   *Path = nil
	)

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if current.Crucible.Position == target && current.Crucible.CanStop() {
			final = current
			break
		}

		moves := current.Crucible.Moves()
		for _, move := range moves {
			if _, found := grid[move.Position]; !found {
				continue
			}
			if current.From != nil && move == current.From.Crucible {
				continue
			}

			cost := current.Cost + grid[move.Position]
			if mem, found := memo[move]; found && cost >= mem {
				continue
			} else {
				memo[move] = cost
			}

			next := &Path{
				Crucible: move,
				From:     current,
				Cost:     cost,
				Weight:   cost + move.Position.Dist(target),
			}
			index := slices.IndexFunc(queue, func(queued *Path) bool {
				return queued.Weight > next.Weight
			})
			if index != -1 {
				queue = slices.Insert(queue, index, next)
			} else {
				queue = append(queue, next)
			}
		}
	}

	if final == nil {
		panic("No route found")
	}
	return *final
}

func part1(grid crucible.Grid) any {
	return FindOptimumPath(grid, crucible.NewCrucible(0, 3)).Cost
}

func part2(grid crucible.Grid) any {
	return FindOptimumPath(grid, crucible.NewCrucible(4, 10)).Cost
}
