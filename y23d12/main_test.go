package main

import (
	"testing"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

func getTestInput() []string {
	return inputFile.ReadInputFile("./sample.txt")
}

func TestPart1(t *testing.T) {
	i := getTestInput()
	part1(i[:1])
}

func TestPart2(t *testing.T) {
	i := getTestInput()
	part2(i)
}
