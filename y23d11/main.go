package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y23d11/input.txt"
var samplePath = "./y23d11/sample.txt"

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

		p1, p2 = parts(input)
	})
	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

type Coord [2]int

func Abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func Max(a ...int) int {
	max := a[0]

	for _, aa := range a[1:] {
		if aa > max {
			max = aa
		}
	}

	return max
}

func Min(a ...int) int {
	min := a[0]

	for _, aa := range a[1:] {
		if aa < min {
			min = aa
		}
	}

	return min
}

type Universe struct {
	Galaxies  map[Coord]interface{}
	Expansion [2]map[int]interface{} // [dimension][i]
}

func NewUniverse() Universe {
	return Universe{
		Galaxies: make(map[Coord]interface{}),
		Expansion: [2]map[int]interface{}{
			make(map[int]interface{}),
			make(map[int]interface{}),
		},
	}
}

func (u Universe) Distance(a, b Coord, scale int) int {
	d := [2]int{0, 0}
	for i := range d {
		// regular space
		d[i] = Abs(a[i] - b[i])

		// Get range of dimesion
		min, max := Min(a[i], b[i]), Max(a[i], b[i])

		// Add expanded space
		for ii := min; ii <= max; ii++ {
			if _, expanded := u.Expansion[i][ii]; expanded {
				d[i] += scale - 1
			}
		}
	}

	return d[0] + d[1]
}

func parts(input []string) (any, any) {
	U := NewUniverse()

	// Populate galaxies and search for empty rows
	X, Y := 0, len(input)
	for y, line := range input {
		clearX := true
		for x, char := range line {
			if char == '#' {
				U.Galaxies[Coord{x, y}] = *new(interface{})
				clearX = false
				if x > X {
					X = x
				}
			}
		}
		if clearX {
			U.Expansion[1][y] = *new(interface{})
		}
	}

	// Search for empty columns
	for x := 0; x < X; x++ {
		clearY := true
		for y := 0; y < Y; y++ {
			if _, galaxy := U.Galaxies[Coord{x, y}]; galaxy {
				clearY = false
				break
			}
		}
		if clearY {
			U.Expansion[0][x] = *new(interface{})
		}
	}

	G := []Coord{} // Create slice for indexing
	for galaxy := range U.Galaxies {
		G = append(G, galaxy)
	}

	ans := [4]int{}
	scale := [4]int{2, 10, 100, 1_000_000}
	for a := range ans {
		for i := range G {
			for j := range G[i+1:] {
				g1, g2 := G[i], G[i+1+j]

				ans[a] += U.Distance(g1, g2, scale[a])
			}
		}
	}

	return ans[0], ans[3]
}
