package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y21d3/input.txt"
var samplePath = "./y21d3/sample.txt"

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

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func bitCount(input []string) []int {
	length := len(input[0])
	bitCount := make([]int, length)

	for _, line := range input {
		for i := range bitCount {
			switch line[i] {
			case '1':
				bitCount[i]++
			case '0':
				bitCount[i]--
			}
		}
	}

	return bitCount
}

func part1(input []string) int {
	bitCount := bitCount(input)

	var a, b int
	for i, count := range bitCount {
		if count >= 0 {
			a |= 1 >> i
		} else {
			b |= 1 >> i
		}
	}
	return a * b
}

func part2(input []string) any {
	bitCount := bitCount(input)
	slices.Reverse(bitCount)

	var oxy, co2 []string

	//hydrate
	for _, line := range input {
		oxy = append(oxy, line)
		co2 = append(co2, line)
	}

	for i, count := range bitCount {
		var _oxy []string
		for _, v := range oxy {
			var expect byte
			if count >= 0 {
				expect = '1'
			} else {
				expect = '0'
			}

			if v[i] == expect {
				_oxy = append(_oxy, v)
			}
		}

		var _co2 []string
		for _, v := range oxy {
			var expect byte
			if count >= 0 {
				expect = '0'
			} else {
				expect = '1'
			}

			if v[i] == expect {
				_co2 = append(_co2, v)
			}
		}
		if len(oxy) > 2 {
			oxy = _oxy
		}
		if len(co2) > 2 {
			co2 = _co2
		}
	}

	a, _ := strconv.ParseInt(oxy[0], 2, 64)
	b, _ := strconv.ParseInt(co2[0], 2, 64)

	return a * b
}
