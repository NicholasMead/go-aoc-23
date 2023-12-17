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
	grid := NewGrid(i)

	ans := part1(grid)

	if ans != 102 {
		t.Fatal(ans)
	}
}

func TestPart2(t *testing.T) {
	i := getTestInput()
	grid := NewGrid(i)
	ans := part2(grid)

	if ans != "_" {
		t.Fatal(ans)
	}
}

func TestPart2Sample2(t *testing.T) {
	i := inputFile.ReadInputFile("./sample2.txt")
	grid := NewGrid(i)
	ans := part2(grid)

	if ans != "_" {
		t.Fatal(ans)
	}
}
