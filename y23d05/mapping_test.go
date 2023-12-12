package main

import "testing"

func TestMapping(t *testing.T) {
	maps := []Mapping{
		{50, 98, 2},
		{52, 50, 48},
	}

	t.Run("Index", func(t *testing.T) {
		seeds := map[Index]Index{
			79: 81, // Seed 79, soil 81.
			14: 14, // Seed 14, soil 14.
			55: 57, // Seed 55, soil 57.
			13: 13, // Seed 13, soil 13.
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
}
