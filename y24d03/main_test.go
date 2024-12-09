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
	ans := part1(i)

	if ans != "_" {
		t.Fatal(ans)
	}
}

func TestPart2(t *testing.T) {
	i := getTestInput()
	ans := part2(i)

	if ans != "_" {
		t.Fatal(ans)
	}
}
