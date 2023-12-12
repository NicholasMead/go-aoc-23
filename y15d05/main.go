package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y15d5/input.txt"

func main() {
	var p1, p2 any = "", ""
	d := common.Timer(func() {
		args := os.Args[1:]
		path := inputPath
		if len(args) > 0 {
			switch args[0] {
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
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

type StringIsNice func(string) bool

func hasEnoughVowels(s string) bool {
	const vowels = "aeiou"
	count := 0

	for _, r := range s {
		if strings.IndexRune(vowels, r) >= 0 {
			count++

			if count >= 3 {
				return true
			}
		}
	}
	return false
}

func hasDoubleRune(s string) bool {
	for i, r := range s[1:] {
		if rune(s[i]) == r {
			return true
		}
	}

	return false
}

func excludesNaughtyStrings(s string) bool {
	naughty := [4]string{
		"ab",
		"cd",
		"pq",
		"xy",
	}
	for _, n := range naughty {
		if strings.Index(s, n) >= 0 {
			return false
		}
	}
	return true
}

func hasSandwichedRune(s string) bool {
	for i := 1; i < len(s)-1; i++ {
		if s[i-1] == s[i+1] {
			return true
		}
	}
	return false
}

func hasDuplicatePair(s string) bool {
	seen := map[[2]byte]int{}

	for i := range s[1:] {
		pair := [2]byte{s[i], s[i+1]}

		if index, ok := seen[pair]; ok {
			if index <= i-2 {
				return true
			}
		} else {
			seen[pair] = i
		}
	}

	return false
}

func part1(input []string) any {
	count := 0
	tests := []StringIsNice{
		hasEnoughVowels,
		hasDoubleRune,
		excludesNaughtyStrings,
	}

	for _, line := range input {
		nice := true

		for _, test := range tests {
			if !test(line) {
				nice = false
				break
			}
		}

		if nice {
			count++
		}
	}

	return count
}

func part2(input []string) any {
	count := 0
	tests := []StringIsNice{
		hasSandwichedRune,
		hasDuplicatePair,
	}

	for _, line := range input {
		nice := true

		for _, test := range tests {
			if !test(line) {
				nice = false
				break
			}
		}

		if nice {
			count++
		}
	}

	return count
}
