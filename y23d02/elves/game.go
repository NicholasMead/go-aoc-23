package elves

import (
	"strconv"
	"strings"
)

type Game struct {
	Id     int
	Rounds []Round
}

type Round struct {
	Red, Green, Blue int
}

func (g Game) CanPlay(bag Bag) bool {
	for _, round := range g.Rounds {
		if !bag.CanDraw(round.Red, round.Green, round.Blue) {
			return false
		}
	}
	return true
}

func (g Game) MinimumPlayableBag() Bag {
	bag := Bag{0, 0, 0}

	for _, r := range g.Rounds {
		bag.Red = max(bag.Red, r.Red)
		bag.Green = max(bag.Green, r.Green)
		bag.Blue = max(bag.Blue, r.Blue)
	}

	return bag
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func ParseGame(s string) Game {
	// s : "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	// mainParts :["Game 1", "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"]
	mainParts := strings.Split(s, ": ")

	// idParts :["Game", "1"]
	idParts := strings.Split(mainParts[0], " ")
	id, _ := strconv.Atoi(idParts[1])

	rounds := []Round{}

	// roundParts :["3 blue, 4 red", "1 red, 2 green, 6 blue", "2 green"]
	roundParts := strings.Split(mainParts[1], "; ")
	for _, roundString := range roundParts {
		round := parseRound(roundString)
		rounds = append(rounds, round)
	}

	return Game{
		Id:     id,
		Rounds: rounds,
	}
}

func parseRound(s string) Round {
	// s :["3 blue, 4 red"]
	// parts :["3", "blue"]
	round := Round{}
	for _, parts := range strings.Split(s, ", ") {
		wordParts := strings.Split(parts, " ")
		switch wordParts[1] {
		case "red":
			round.Red, _ = strconv.Atoi(wordParts[0])
		case "green":
			round.Green, _ = strconv.Atoi(wordParts[0])
		case "blue":
			round.Blue, _ = strconv.Atoi(wordParts[0])
		}
	}
	return round
}
