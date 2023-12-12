package main

import (
	"strconv"
	"strings"
)

type ScratchCard struct {
	id      int
	winners map[int]interface{}
	found   map[int]interface{}
}

func ParseScratchCard(line string) ScratchCard {
	card := ScratchCard{
		winners: make(map[int]interface{}),
		found:   make(map[int]interface{}),
	}
	colon, pipe := strings.IndexRune(line, ':'), strings.IndexRune(line, '|')

	headers := strings.Split(line[:colon], " ")
	id, err := strconv.Atoi(headers[len(headers)-1])
	if err != nil {
		panic(err)
	} else {
		card.id = id
	}

	winnersStrings := strings.Split(
		strings.Trim(line[colon+1:pipe], " "),
		" ")

	for _, winner := range winnersStrings {
		if winner == " " || winner == "" {
			continue
		}

		value, err := strconv.Atoi(winner)
		if err != nil {
			panic(err)
		}
		card.winners[value] = new(interface{})
	}

	foundStrings := strings.Split(
		strings.Trim(line[pipe+1:], " "),
		" ")

	for _, found := range foundStrings {
		if found == " " || found == "" {
			continue
		}

		value, err := strconv.Atoi(found)
		if err != nil {
			panic(err)
		}
		card.found[value] = new(interface{})
	}

	return card
}

func (sc *ScratchCard) Score() int {

	if m := sc.Matches(); m == 0 {
		return 0
	} else {
		return 1 << (m - 1)
	}
}

func (sc *ScratchCard) Matches() int {
	matches := 0

	for f := range sc.found {
		if _, ok := sc.winners[f]; ok {
			matches++
		}
	}

	return matches
}
