package main

import (
	"testing"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

func getTestInput() [2][]string {
	return [2][]string{
		inputFile.ReadInputFile("./sample1.txt"),
		inputFile.ReadInputFile("./sample2.txt"),
	}
}

func TestPart1(t *testing.T) {
	inputs := getTestInput()
	for i, expect := range []int{32000000, 11687500} {
		input := inputs[i]
		ans := part1(input)

		if ans != expect {
			t.Errorf("%v: got %v expected %v", i, ans, expect)
		}
	}
}

// func TestPart2(t *testing.T) {
// 	i := getTestInput()
// 	ans := part2(i)

// 	if ans != "_" {
// 		t.Fatal(ans)
// 	}
// }
