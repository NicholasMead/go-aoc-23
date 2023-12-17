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

		candidates := []*Path{}
		moves := current.Crucible.Moves()
		for _, move := range moves {
			if _, found := grid[move.Position]; !found {
				continue
			}
			if current.From != nil && move == current.From.Crucible {
				continue
			}

			cost := current.Cost + grid[move.Position]
			candidates = append(candidates, &Path{
				Crucible: move,
				From:     current,
				Cost:     cost,
				Weight:   cost + move.Position.Dist(target),
			})
		}

		for _, next := range candidates {
			if mem, found := memo[next.Crucible]; found && next.Cost >= mem {
				continue
			} else {
				memo[next.Crucible] = next.Cost
			}

			queue = append(queue, next)
		}

		slices.SortFunc(queue, func(a, b *Path) int {
			return a.Weight - b.Weight
		})
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
