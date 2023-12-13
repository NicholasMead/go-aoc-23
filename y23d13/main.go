package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y23d13/input.txt"
var samplePath = "./y23d13/sample.txt"

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
		lavas := parseLavas(input)

		p1, p2 = part1(lavas), part2(lavas)
	})

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

func parseLavas(input []string) []Lava {
	lavas := []Lava{}
	lava := Lava{}

	for _, line := range input {
		if line == "" {
			lavas = append(lavas, lava)
			lava = Lava{}
		} else {
			lava = append(lava, line)
		}
	}
	return append(lavas, lava)
}

func part1(input []Lava) any {
	rowCount, columnCount := 0, 0

	for i, lava := range input {
		rows, columns := lava.Mirrors()

		for _, row := range rows {
			rowCount += row
		}
		for _, column := range columns {
			columnCount += column
		}

		if len(rows)+len(columns) == 0 {
			log.Panicf("No mirror in lava:%v\n%v", i, lava)
		}
	}

	return (100 * rowCount) + columnCount
}

func part2(input []Lava) any {
	rowCount, columnCount := 0, 0

	for i, lava := range input {
		rows, columns := lava.Mirrors()

		didClean := func() bool {
			for y := 0; y < len(lava); y++ {
				for x := 0; x < len(lava[y]); x++ {
					clean := lava.Clean(x, y)

					cleanRows, cleanColumns := clean.Mirrors()

					for _, cleanRow := range cleanRows {
						if !slices.Contains(rows, cleanRow) {
							rowCount += cleanRow
							return true
						}
					}
					for _, cleanColumn := range cleanColumns {
						if !slices.Contains(columns, cleanColumn) {
							columnCount += cleanColumn
							return true
						}
					}
				}
			}
			return false
		}()

		if !didClean {
			log.Panicf("No clean mirror in lava:%v\n%v", i, lava)
		}
	}

	return (100 * rowCount) + columnCount
}
