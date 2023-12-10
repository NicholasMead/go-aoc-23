package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d10/grid"
)

var inputPath = "./y23d10/input.txt"
var samplePath = "./y23d10/sample.txt"

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
	fmt.Printf("Time: %vms\n", d.Milliseconds())
}

func PipeMoves(g grid.Grid, c grid.Coord) []grid.Coord {
	x, y := c[0], c[1]
	switch g[x][y] {
	case '-':
		return []grid.Coord{{x - 1, y}, {x + 1, y}}
	case '|':
		return []grid.Coord{{x, y + 1}, {x, y - 1}}
	case 'F':
		return []grid.Coord{{x, y + 1}, {x + 1, y}}
	case '7':
		return []grid.Coord{{x - 1, y}, {x, y + 1}}
	case 'J':
		return []grid.Coord{{x, y - 1}, {x - 1, y}}
	case 'L':
		return []grid.Coord{{x, y - 1}, {x + 1, y}}
	case 'S':
		return c.Moves()
	case '.':
		return []grid.Coord{}

	default:
		panic(string(g[x][y]))
	}
}

func FindLoop(g grid.Grid, start grid.Coord) map[grid.Coord]rune {
	memo := map[grid.Coord]int{start: 0}
	queue := []grid.Coord{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		moves := PipeMoves(g, current)
		for _, move := range moves {
			if !slices.Contains(PipeMoves(g, move), current) {
				continue //not a mutually connected pipe
			}

			cost := memo[current] + 1
			if prevCost, found := memo[move]; found && cost >= prevCost {
				continue
			}

			memo[move] = cost
			queue = append(queue, move)
		}
	}

	loop := map[grid.Coord]rune{}
	for c := range memo {
		loop[c] = g.At(c)
	}
	return loop
}

func part1(input []string) any {
	g, S := grid.FromInput(input)
	loop := FindLoop(g, S)
	return len(loop) / 2
}

func isBound(g grid.Grid, memo map[grid.Coord]bool, X, Y int, start grid.Coord) (result bool) {
	// Check memory
	if exits, found := memo[start]; found {
		return exits
	}

	// Keep track of all seen coordinates
	// and memorise the result for all seen coordinates
	// i.e. all connected nodes have the same 'canExits' value
	seen := map[grid.Coord]interface{}{start: struct{}{}}
	defer func() {
		for s := range seen {
			memo[s] = result
		}
	}()

	// edge cases, literally
	x, y := start[0], start[1]
	if x <= 0 || y <= 0 || x >= X-1 || y >= Y-1 {
		return false
	}

	queue := []grid.Coord{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		moves := current.Moves()

		for _, move := range moves {
			// Check if pre computed
			if moveExits, found := memo[move]; found {
				return moveExits
			}

			mx, my := move[0], move[1]

			// Check if we're at the edge
			if mx < 0 || my < 0 || mx >= X || my >= Y {
				return false
			}

			// Check if valid move
			if g[mx][my] != '.' {
				continue
			}

			// Check if duplicate
			if _, seen := seen[move]; seen {
				continue
			}
			queue = append(queue, move)
			seen[move] = struct{}{}
		}
	}

	// If we made it here, we ran out of un-explored spaces
	// before we found the edge... No route out!
	return true
}

func expand(pipe rune) (out [3][3]rune) {
	// expansions below are in [y][x] formart to make them readable
	// we transpose the output on returning
	defer func() {
		swap := [3][3]rune{}
		for x := range swap {
			for y := range swap {
				swap[x][y] = out[y][x]
			}
		}
		out = swap
	}()

	switch pipe {
	case '.':
		out = [3][3]rune{
			{'.', '.', '.'},
			{'.', '.', '.'},
			{'.', '.', '.'},
		}
	case '-':
		out = [3][3]rune{
			{'.', '.', '.'},
			{'-', '-', '-'},
			{'.', '.', '.'},
		}
	case '|':
		out = [3][3]rune{
			{'.', '|', '.'},
			{'.', '|', '.'},
			{'.', '|', '.'},
		}
	case 'F':
		out = [3][3]rune{
			{'.', '.', '.'},
			{'.', 'F', '-'},
			{'.', '|', '.'},
		}
	case '7':
		out = [3][3]rune{
			{'.', '.', '.'},
			{'-', '7', '.'},
			{'.', '|', '.'},
		}
	case 'J':
		out = [3][3]rune{
			{'.', '|', '.'},
			{'-', 'J', '.'},
			{'.', '.', '.'},
		}
	case 'L':
		out = [3][3]rune{
			{'.', '|', '.'},
			{'.', 'L', '-'},
			{'.', '.', '.'},
		}
	case 'S':
		out = [3][3]rune{
			{'.', '|', '.'},
			{'-', 'S', '-'},
			{'.', '|', '.'},
		}
	default:
		panic(string(pipe))
	}
	return
}

func part2(input []string) any {
	// recompute part 1
	g, s := grid.FromInput(input)
	X, Y := g.Dimension()
	l := FindLoop(g, s)

	// expand x,y grid to xx,yy (to include space between pipes)
	gg := grid.EmptyGrid(X*3, Y*3)
	XX, YY := gg.Dimension()
	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			// do not include none-looping pipes in expanded grid
			// replace them with empty space
			pipe, found := l[grid.Coord{x, y}]
			if !found {
				pipe = '.'
			}

			expanded := expand(pipe)

			for ex := 0; ex < 3; ex++ {
				for ey := 0; ey < 3; ey++ {
					xx, yy := x*3+ex, y*3+ey
					gg[xx][yy] = expanded[ex][ey]
				}
			}
		}
	}

	// find all positions (not part of the pipe) that have paths to the edge (bound)
	bound := map[grid.Coord]bool{}
	func() {
		for xx := 0; xx < XX; xx++ {
			for yy := 0; yy < YY; yy++ {
				coord := grid.Coord{xx, yy}
				pipe := gg.At(coord)
				if pipe != '.' {
					continue
				}

				if isBound(gg, bound, XX, YY, coord) {
					return // if we found a bound coord, we've memoed them all!
				}
			}
		}
	}()

	// cross reference positions in origional x,y grid with exits value from xx,yy
	ans := 0
	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			xx, yy := (x*3)+1, (y*3)+1

			_, onLoop := l[grid.Coord{x, y}]

			if !onLoop && bound[grid.Coord{xx, yy}] {
				ans++
			}
		}
	}
	return ans
}
