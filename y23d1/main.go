package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y23d1/input.txt"
var samplePath = "./y23d1/sample.txt"

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

func extract(line string) int {
	numString := ""

	for _, letter := range line {
		if letter >= '0' && letter <= '9' {
			numString += string(letter)
		}
	}

	switch len(numString) {
	case 0:
		return 0

	case 1:
		num := int(numString[0] - '0')
		num = (num * 10) + num
		return num

	case 2:
		num, err := strconv.ParseInt(numString, 10, 32)
		if err != nil {
			panic(err)
		}
		return int(num)

	default:
		numString = string(numString[0]) +
			string(numString[len(numString)-1])

		num, err := strconv.ParseInt(numString, 10, 32)
		if err != nil {
			panic(err)
		}
		return int(num)
	}
}

func part1(input []string) any {
	value := 0

	for _, line := range input {
		value += extract(line)
	}

	return value
}

func part2(input []string) any {
	value := 0

	for _, line := range input {
		spelling := []struct {
			word  string
			digit string
		}{
			{"one", "o1e"},
			{"two", "t2o"},
			{"three", "t3e"},
			{"four", "f4r"},
			{"five", "f5e"},
			{"six", "s6x"},
			{"seven", "s7n"},
			{"eight", "e8t"},
			{"nine", "n9e"},
		}

		for _, v := range spelling {
			line = strings.ReplaceAll(line, v.word, v.digit)
		}

		value += extract(line)
	}

	return value
}
