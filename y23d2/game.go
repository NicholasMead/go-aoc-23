package main

import (
	"strconv"
	"strings"
)

type game struct {
	num    int
	rounds []cubeSet
}

func parseGame(line string) (g game) {

	// example:
	//   line = 		"game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	//   parts = 		["game 1","3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"]
	//   headerParts = 	["game","1"]
	//   rounds = 	 	["3 blue, 4 red","1 red, 2 green, 6 blue","2 green"]

	//split header from rounds
	parts := strings.Split(line, ": ")

	//get game num from header
	headerParts := strings.Split(parts[0], " ")
	g.num, _ = strconv.Atoi(headerParts[1])

	//split and parse rounds
	rounds := strings.Split(parts[1], "; ")
	for _, round := range rounds {
		cubes := cubeSetFromRoundString(round)
		g.rounds = append(g.rounds, cubes)
	}
	return
}
