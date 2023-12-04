package main

import (
	"testing"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

func getTestInput() []string {
	return inputFile.ReadInputFile("./sample.txt")
}

func TestPart1(t *testing.T) {
	input := getTestInput()
	ans := part1(input)

	if ans != 198 {
		t.Fatal("expected 198 got", ans)
	}
}

func TestPart2(t *testing.T) {
	input := getTestInput()
	ans := part2(input)

	if ans != 230 {
		t.Fatal("expected 230 got", ans)
	}
}
