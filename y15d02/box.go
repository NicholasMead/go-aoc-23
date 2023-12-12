package main

import (
	"strconv"
	"strings"
)

type Box [3]int

func (b Box) Sides() []Side {
	return []Side{
		{b[0], b[1]},
		{b[1], b[2]},
		{b[2], b[0]},
	}
}

func (b Box) Volume() int {
	return b[0] * b[1] * b[2]
}

func parse(dim string) (out Box) {
	for i := range [2]struct{}{} {
		index := strings.IndexRune(dim, 'x')
		out[i], _ = strconv.Atoi(dim[:index])
		dim = dim[index+1:]
	}
	out[2], _ = strconv.Atoi(dim)
	return
}
