package main

import (
	"testing"

	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

func TestMapping(t *testing.T) {
	maps := []Mapping{
		{50, 98, 2},
		{52, 50, 48},
	}

	t.Run("Index", func(t *testing.T) {
		seeds := map[Index]Index{
			79: 81, // Seed number 79 corresponds to soil number 81.
			14: 14, // Seed number 14 corresponds to soil number 14.
			55: 57, // Seed number 55 corresponds to soil number 57.
			13: 13, // Seed number 13 corresponds to soil number 13.
		}

		for from, to := range seeds {
			var result Index = from
			var didMap bool = false

			for _, m := range maps {
				result, didMap = m.MapIndex(result)
				if didMap {
					break
				}
			}

			if result != to {
				t.Errorf("got %v for %v expected %v", result, from, to)
			}
		}
	})

	// t.Run("Range", func(t *testing.T) {
	// 	ranges := map[Range][]Range{
	// 		{99, 2}: {{51, 1}, {100, 1}},
	// 	}

	// 	for input, output := range ranges {
	// 		result, didMap := maps[0].MapRange(input)

	// 		if !slices.Equal(result, output){
	// 			t.Errorf("Got %v for %v expected %v", result, input, output)
	// 		}
	// 	}
	// })
}

func TestAlamac(t *testing.T) {
	input := inputFile.ReadInputFile("./sample.txt")
	// Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
	// Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
	// Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, humidity 82, location 86.
	// Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, humidity 35, location 35.
	seeds := map[Index]Index{
		79: 82, // Seed number 79 corresponds to soil number 81.
		14: 43, // Seed number 14 corresponds to soil number 14.
		55: 86, // Seed number 55 corresponds to soil number 57.
		13: 35, // Seed number 13 corresponds to soil number 13.
	}
	alamac := AlamacFromInput(input)

	for from, to := range seeds {
		result := alamac.SeedLocationIndex(from)

		if result != to {
			t.Errorf("Got %v from %v expected %v", result, from, to)
		}
	}
}
