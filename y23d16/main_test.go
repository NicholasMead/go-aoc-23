package main

import (
	"testing"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

func getTestInput() []string {
	return inputFile.ReadInputFile("./input.txt")
}

func TestPart1(t *testing.T) {
	i := getTestInput()
	ans := part1(i)

	if ans != 8901 {
		t.Fatal(ans)
	}
}

func TestPart2(t *testing.T) {
	i := getTestInput()
	ans := part2(i)

	if ans != 9064 {
		t.Fatal(ans)
	}
}
