package main

import (
	"slices"
	"testing"
)

func TestRange(t *testing.T) {
	t.Run("Split", func(t *testing.T) {
		input := Range{10, 20}

		output := input.Split(10)

		if output[0] != (Range{10, 10}) {
			t.Error(output[0])
		}
		if output[1] != (Range{20, 10}) {
			t.Error(output[1])
		}
	})

	t.Run("Cap", func(t *testing.T) {
		input := Range{10, 20}

		cap := input.Cap()

		if cap != 30 {
			t.Error(cap)
		}
	})

	t.Run("Expand", func(t *testing.T) {
		input := Range{5, 3}

		expanded := input.Expand()

		if !slices.Equal([]Index{5, 6, 7}, expanded) {
			t.Error(expanded)
		}
	})

	t.Run("Intersect", func(t *testing.T) {
		input := Range{10, 10}
		cases := map[Range]struct {
			intersect Range
			remainder []Range
		}{
			{0, 10}: {Range{}, []Range{{10, 10}}},
			{5, 10}: {Range{10, 5}, []Range{{15, 5}}},
			{12, 5}: {Range{12, 5}, []Range{{10, 2}, {17, 3}}},
			{18, 5}: {Range{18, 2}, []Range{{10, 8}}},
			{21, 5}: {Range{}, []Range{{10, 10}}},
		}

		for test, result := range cases {
			intersect, remainder := input.Intersect(test)

			if result.intersect != intersect {
				t.Error("Intersect", intersect, "expected", result.intersect)
			}
			if !slices.Equal(remainder, result.remainder) {
				t.Error("remainder", remainder, "expected", result.remainder)
			}
		}
	})
}
