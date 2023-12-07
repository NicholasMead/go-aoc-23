package main

import "testing"

func TestHand(t *testing.T) {
	cases := map[Hand]Score{
		NewHand("32T3K"): onePair,
		NewHand("T55J5"): threeOfAKind,
		NewHand("KK677"): twoPair,
		NewHand("KTJJT"): twoPair,
		NewHand("QQQJA"): threeOfAKind,
		NewHand("QQQQA"): fourOfAKind,
		NewHand("QQQQQ"): fiveOfAKind,
		NewHand("QAQAQ"): fullHouse,
	}

	for input, output := range cases {
		score := input.Score()

		if score != output {
			t.Error(input, output, score)
		}
	}
}
