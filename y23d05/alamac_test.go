package main

import (
	"testing"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

func TestAlamac(t *testing.T) {
	input := inputFile.ReadInputFile("./sample.txt")
	seeds := map[Index]Index{
		79: 82, // Seed 79, location 82.
		14: 43, // Seed 14, location 43.
		55: 86, // Seed 55, location 86.
		13: 35, // Seed 13, location 35.
	}
	alamac := AlamacFromInput(input)

	for from, to := range seeds {
		result := alamac.SeedLocationIndex(from)

		if result != to {
			t.Errorf("Got %v from %v expected %v", result, from, to)
		}
	}
}
