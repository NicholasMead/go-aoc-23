package elves

import (
	"slices"
	"testing"
)

func TestGame(t *testing.T) {
	games := []Game{
		// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		{
			Id: 1,
			Rounds: []Round{
				{Red: 4, Green: 0, Blue: 3},
				{Red: 1, Green: 2, Blue: 6},
				{Red: 0, Green: 2, Blue: 0},
			},
		},
		// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
		{
			Id: 2,
			Rounds: []Round{
				{Red: 0, Green: 2, Blue: 1},
				{Red: 1, Green: 3, Blue: 4},
				{Red: 0, Green: 1, Blue: 1},
			},
		},
		// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
		{
			Id: 3,
			Rounds: []Round{
				{Red: 20, Green: 8, Blue: 6},
				{Red: 4, Green: 13, Blue: 5},
				{Red: 1, Green: 5, Blue: 0},
			},
		},
		// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
		{
			Id: 4,
			Rounds: []Round{
				{Red: 3, Green: 1, Blue: 6},
				{Red: 6, Green: 3, Blue: 0},
				{Red: 14, Green: 3, Blue: 15},
			},
		},
		// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
		{
			Id: 5,
			Rounds: []Round{
				{Red: 6, Green: 3, Blue: 1},
				{Red: 1, Green: 2, Blue: 2},
			},
		},
	}

	t.Run("CanPlay", func(t *testing.T) {
		tests := []struct {
			game   Game
			expect bool
		}{
			{games[0], true},
			{games[1], true},
			{games[2], false},
			{games[3], false},
			{games[4], true},
		}

		bag := Bag{12, 13, 14}

		for _, test := range tests {
			result := test.game.CanPlay(bag)

			if result != test.expect {
				t.Errorf("%v: got %v expected %v", test.game.Id, result, test.expect)
			}
		}
	})

	t.Run("MinimumPlayableBag", func(t *testing.T) {
		tests := []struct {
			game   Game
			expect Bag
		}{
			{games[0], Bag{4, 2, 6}},
			{games[1], Bag{1, 3, 4}},
			{games[2], Bag{20, 13, 6}},
			{games[3], Bag{14, 3, 15}},
			{games[4], Bag{6, 3, 2}},
		}

		for _, test := range tests {
			result := test.game.MinimumPlayableBag()

			if result != test.expect {
				t.Errorf("%v: got %v expected %v", test.game.Id, result, test.expect)
			}
		}
	})

	t.Run("ParseGame", func(t *testing.T) {
		lines := []string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		}

		for i, line := range lines {
			game := ParseGame(line)

			if !slices.Equal(game.Rounds, games[i].Rounds) {
				t.Errorf("got Rounds %v expected %v for %v", game.Rounds, games[i].Rounds, line)
			}
			if game.Id != games[i].Id {
				t.Errorf("got Id %v expected %v for %v", game.Id, games[i].Id, line)
			}
		}
	})
}
