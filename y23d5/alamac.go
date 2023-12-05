package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Index int

type Range struct {
	start Index
	len   int
}

func (r Range) Split(at int) [2]Range {
	if at <= 0 || at > r.len {
		panic("Out of range")
	}

	return [2]Range{
		{r.start, at},
		{r.start + Index(at), r.len - at},
	}
}

func (r Range) Cap() Index {
	return r.start + Index(r.len)
}

func (r Range) Expand() []Index {
	indexes := make([]Index, r.len)
	for i := range indexes {
		indexes[i] = r.start + Index(i)
	}
	return indexes
}

type Mapping struct {
	dest, src Index
	len       int
}

func (m Mapping) Src() Range {
	return Range{m.src, m.len}
}

func (m Mapping) MapIndex(from Index) (Index, bool) {
	if from < m.src || from >= m.src+Index(m.len) {
		return from, false
	}

	return from + (m.dest - m.src), true
}

func (m Mapping) MapRange(from Range) (Range, []Range, bool) {
	// Underflow
	if from.start < m.src {
		// total underflow
		if from.Cap() < m.src {
			return Range{}, []Range{from}, false
		}

		// partial underflow
		split := from.Split(int(m.src) - int(from.start))
		next, overflow, _ := m.MapRange(split[1])

		return next, append(overflow, split[0]), true
	}

	// total overflow
	if from.start >= m.Src().Cap() {
		return Range{}, []Range{from}, false
	}

	// no overflow
	if from.Cap() <= m.Src().Cap() {
		next := Range{
			from.start - m.src + m.dest,
			from.len,
		}
		return next, []Range{}, true
	}

	// partial overlapping
	split := from.Split(from.len - int(from.Cap()-m.Src().Cap()))
	next := split[0]
	next.start -= m.src
	next.start += m.dest
	return next, []Range{split[1]}, true
}

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
		search := append([]Range{}, locations...)
		locations = []Range{}

		for len(search) > 0 {
			var next Range
			next, search = search[0], search[1:]

			for _, m := range a[i] {
				output, overflow, ok := m.MapRange(next)

				fmt.Sprintln(output, overflow, ok)

				if ok {
					locations = append(locations, output)
					search = append(search, overflow...)
					break
				}
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
