package main

import (
	"slices"
	"strings"
)

type Lava []string

func (lava Lava) Mirrors() (rows []int, columns []int) {
	rows = lava.MirrorRow()
	columns = lava.MirrorColumn()
	return
}

func (lava Lava) MirrorRow() []int {
	mirrors := []int{}

	for row := 1; row < len(lava); row++ {
		if isMirror(lava[:row], lava[row:]) {
			mirrors = append(mirrors, row)
		}
	}

	return mirrors
}

func (lava Lava) MirrorColumn() []int {
	return lava.Transpose().MirrorRow()
}

func (lava Lava) Transpose() Lava {
	X := len(lava[0])
	Y := len(lava)

	transposed := make(Lava, X)

	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			transposed[x] += string(lava[y][x])
		}
	}

	return transposed
}

func (lava Lava) Clean(x, y int) Lava {
	cleaned := Lava{}

	for yy, line := range lava {
		cleaned = append(cleaned, "")
		for xx, rock := range line {
			next := rock
			if xx == x && yy == y {
				switch rock {
				case '.':
					next = '#'
				case '#':
					next = '.'
				default:
					panic(string(next))
				}
			}
			cleaned[yy] += string(next)
		}
	}

	return cleaned
}

func (lava Lava) String() string {
	return strings.Join(lava, "\n")
}

func isMirror(left, right []string) bool {
	L, R := len(left), len(right)
	I := min(L, R)

	left = append([]string{}, left...)
	slices.Reverse(left)

	var (
		leftExpanded  = strings.Join(left[:I], "\n")
		rightExpanded = strings.Join(right[:I], "\n")
	)
	return leftExpanded == rightExpanded
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
