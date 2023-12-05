package main

import (
	"strconv"
	"strings"
)

type Index int

type Alamac [7][]Mapping

func (a Alamac) SeedLocationIndex(seed Index) Index {
	location := seed
	for i := 0; i < 7; i++ {
		for _, m := range a[i] {
			var ok bool
			location, ok = m.MapIndex(location)
			if ok {
				break
			}
		}
	}
	return location
}

func (a Alamac) SeedLocationRange(seeds Range) []Range {
	locations := []Range{seeds}

	for i := 0; i < 7; i++ {
		queue := append([]Range{}, locations...)
		locations = []Range{}

		for len(queue) > 0 {
			next := queue[0]
			queue = queue[1:]
			didMapping := false

			for _, m := range a[i] {
				output, overflow := m.MapRange(next)

				if output != (Range{}) {
					locations = append(locations, output)
					queue = append(queue, overflow...)
					didMapping = true
					break
				}
			}

			if !didMapping {
				locations = append(locations, next)
			}
		}
	}

	return locations
}

func AlamacFromInput(input []string) Alamac {
	alamac := Alamac{}
	mapping := 0

	for _, line := range input[2:] {
		if line == "" {
			mapping++
			continue
		}
		if alamac[mapping] == nil {
			alamac[mapping] = []Mapping{}
			continue
		}
		split := strings.Split(line, " ")
		dest, _ := strconv.Atoi(split[0])
		src, _ := strconv.Atoi(split[1])
		len, _ := strconv.Atoi(split[2])
		alamac[mapping] = append(alamac[mapping], Mapping{Index(dest), Index(src), len})
	}

	return alamac
}
