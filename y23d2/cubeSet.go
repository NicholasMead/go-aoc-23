package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type cubeSet struct {
	red, green, blue int
}

func parseCubeSet(game string) cubeSet {
	set := cubeSet{}

	for _, colour := range []struct {
		word string
		val  *int
	}{
		{"red", &set.red},
		{"green", &set.green},
		{"blue", &set.blue},
	} {
		regx := regexp.MustCompile(fmt.Sprintf(`\d+ %v`, colour.word))
		match := regx.FindString(game)
		if match != "" {
			val := regexp.MustCompile(`\d+`).FindString(match)
			*colour.val, _ = strconv.Atoi(val)
		}
	}

	return set
}
