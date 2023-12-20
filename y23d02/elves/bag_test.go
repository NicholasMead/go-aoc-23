package elves

import (
	"fmt"
	"testing"
)

func TestBag(t *testing.T) {

	t.Run("CanDraw", func(t *testing.T) {
		cases := []struct {
			r, g, b int
			expect  bool
		}{
			{0, 0, 0, true},
			{1, 1, 1, true},
			{12, 13, 14, true},
			{13, 13, 14, false},
			{12, 14, 14, false},
			{12, 13, 15, false},
		}
		bag := Bag{12, 13, 14}

		for _, test := range cases {
			t.Run(fmt.Sprint(test), func(t *testing.T) {

				result := bag.CanDraw(test.r, test.g, test.b)

				if result != test.expect {
					t.Errorf("got %v expected %v", result, test.expect)
				}
			})
		}
	})

	t.Run("Power", func(t *testing.T) {
		tests := []struct {
			bag   Bag
			power int
		}{
			{Bag{4, 2, 6}, 48},
			{Bag{1, 3, 4}, 12},
			{Bag{20, 13, 6}, 1560},
			{Bag{14, 3, 15}, 630},
			{Bag{6, 3, 2}, 36},
		}

		for _, test := range tests {
			result := test.bag.Power()

			if result != test.power {
				t.Errorf("got %v expected %v", result, test.power)
			}
		}
	})
}
