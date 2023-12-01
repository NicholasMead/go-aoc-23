package main

import (
	"testing"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

func getTestInput() []string] {
	return inputFile.ReadInputFile("./sample.txt")
}

func TestPart1(t *testing.T) {
	i := getTestInput()
	_ := part1(i)
}

func TestPart2(t *testing.T) {
	i := getTestInput()
	_ := part2(i)
}
