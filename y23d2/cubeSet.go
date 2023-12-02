package main

import (
	"strconv"
	"strings"
)

type cubeSet struct {
	red, green, blue int
}

func cubeSetFromRoundString(round string) (set cubeSet) {
	partMap := map[string]*int{
		"red":   &set.red,
		"green": &set.green,
		"blue":  &set.blue,
	}

	// example:
	//   round:     "3 blue, 4 red"
	//   cubes:     ["3 blue","4 red"]
	//   cubeParts: ["3","blue"]
	
	cubes := strings.Split(round, ", ")
	for _, cube := range cubes {
		cubeParts := strings.Split(cube, " ")
		name, val := cubeParts[1], cubeParts[0]
		*partMap[name], _ = strconv.Atoi(val)
	}
	return
}
