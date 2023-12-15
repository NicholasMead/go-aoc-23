package main

import (
	"strings"
	"testing"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

func getTestInput() []string {
	return inputFile.ReadInputFile("./sample.txt")
}

func TestPart1(t *testing.T) {
	i := strings.Split(getTestInput()[0], ",")
	ans := part1(i)

	if ans != 1320 {
		t.Fatal(ans)
	}
}

func TestPart2(t *testing.T) {
	i := strings.Split(getTestInput()[0], ",")
	ans := part2(i)

	if ans != 145 {
		t.Fatal(ans)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := getTestInput()
		data := strings.Split(input[0], ",")
		part1(data)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := getTestInput()
		data := strings.Split(input[0], ",")
		part2(data)
	}
}
