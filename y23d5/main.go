package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

// var inputPath = "./input.txt"

var inputPath = "./y23d5/input.txt"
var samplePath = "./y23d5/sample.txt"

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

	var p1, p2 any = "", ""
	input := inputFile.ReadInputFile(path)

	d := common.Timer(func() {
		p1, p2 = part1(input), part2(input)
	})

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

func runAlamac(alamac Alamac, seeds []Index) Index {
	min := Index(math.MaxInt)

	for _, seed := range seeds {
		loc := alamac.SeedLocationIndex(seed)
		if loc < min {
			min = loc
		}
	}

	return min
}

func part1(input []string) Index {
	alamac := AlamacFromInput(input)
	seedString := strings.Split(input[0][7:], " ")
	seeds := []Index{}
	for _, s := range seedString {
		i, _ := strconv.Atoi(s)
		seeds = append(seeds, Index(i))
	}

	return runAlamac(alamac, seeds)
}

func part2(input []string) any {
	alamac := AlamacFromInput(input)
	seedString := strings.Split(input[0][7:], " ")
	seedRanges := []Range{}
	for i := 0; i < len(seedString)-1; i += 2 {
		start, _ := strconv.Atoi(seedString[i])
		length, _ := strconv.Atoi(seedString[i+1])
		seedRanges = append(seedRanges, Range{Index(start), length})
	}

	min := Index(math.MaxInt)

	for _, seeds := range seedRanges {
		locations := alamac.SeedLocationRange(seeds)

		for _, loc := range locations {
			if loc.start < min {
				min = loc.start
			}
		}
	}

	return min
}
