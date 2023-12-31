package main

import (
	"strconv"
)

type schematic struct {
	parts   []part
	symbols map[coord]rune
}

type part struct {
	value      int
	start, end coord
}

func (p part) inRange(coord coord) bool {
	return p.start[0]-1 <= coord[0] &&
		p.start[1]-1 <= coord[1] &&
		p.end[0]+1 >= coord[0] &&
		p.end[1]+1 >= coord[1]
}

type gear [2]int

func (g gear) ratio() int {
	return g[0] * g[1]
}

type coord [2]int

func parseScematic(doc []string) schematic {
	s := schematic{
		parts:   []part{},
		symbols: make(map[coord]rune),
	}

	for y, line := range doc {
		var current *part = nil

		for x, r := range line {

			if v, err := strconv.Atoi(string(r)); err == nil {
				if current == nil {
					current = &part{
						v,
						coord{x, y},
						coord{x, y},
					}
				} else {
					current.value *= 10
					current.value += v
					current.end = coord{x, y}
				}
			} else {
				if current != nil {
					s.parts = append(s.parts, *current)
					current = nil
				}

				if r != '.' {
					s.symbols[coord{x, y}] = r
				}
			}
		}

		if current != nil {
			s.parts = append(s.parts, *current)
		}
	}

	return s
}

func (s *schematic) validParts() []part {
	parts := map[part]interface{}{}

	for coord := range s.symbols {
		for _, p := range s.parts {
			if p.inRange(coord) {
				parts[p] = new(interface{})
			}
		}
	}

	keys := make([]part, 0, len(parts))
	for p := range parts {
		keys = append(keys, p)
	}
	return keys
}

func (s *schematic) gears() (gears []gear) {
	for coord, symb := range s.symbols {
		if symb != '*' {
			continue
		}

		parts := []part{}

		for _, p := range s.parts {
			if p.inRange(coord) {
				parts = append(parts, p)
			}
		}

		if len(parts) == 2 {
			gears = append(gears, gear{parts[0].value, parts[1].value})
		}
	}
	return
}
