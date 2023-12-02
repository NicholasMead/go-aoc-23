package main

import (
	"regexp"
	"strconv"
)

type game struct {
	num    int
	rounds []cubeSet
}

func parseGame(line string) game {
	g := game{}
	g.num, _ = strconv.Atoi(
		regexp.MustCompile(`\d+`).FindString(line))

	line = regexp.MustCompile(`game \d+: `).ReplaceAllString(line, "")

	rounds := regexp.MustCompile(`[\w ,]+;?`).FindAllString(line, -1)

	for _, round := range rounds {
		cubes := parseCubeSet(round)
		g.rounds = append(g.rounds, cubes)
	}
	
	return g
}
