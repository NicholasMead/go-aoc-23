package volcano

import (
	"errors"
	"fmt"
)

type Mirror interface {
	Reflect(Beam) []Beam
}

func NewMirrorFromRune(r rune) Mirror {
	switch r {
	case '/':
		return rightMirror{}
	case '\\':
		return leftMirror{}
	case '|':
		return vericalSpliter{}
	case '-':
		return horizontalSpliter{}
	}

	panic(errors.Join(
		fmt.Errorf("unknown rune: %s", string(r)),
		errors.ErrUnsupported,
	))
}

// Mirror '/'
type rightMirror struct{}

// Reflect implements Mirror.
func (rightMirror) Reflect(b Beam) []Beam {
	return []Beam{{
		Position: b.Position,
		Direction: Vector{
			-b.Direction[1],
			-b.Direction[0],
		},
	}}
}

// Mirror '\'
type leftMirror struct{}

// Reflect implements Mirror.
func (leftMirror) Reflect(b Beam) []Beam {
	return []Beam{{
		Position: b.Position,
		Direction: Vector{
			b.Direction[1],
			b.Direction[0],
		},
	}}
}

// Splitter '|'
type vericalSpliter struct {
}

// Reflect implements Mirror.
func (vericalSpliter) Reflect(b Beam) []Beam {
	if b.Direction[0] == 0 {
		return []Beam{b}
	}

	return []Beam{
		{
			Position:  b.Position,
			Direction: Vector{0, 1},
		},
		{
			Position:  b.Position,
			Direction: Vector{0, -1},
		},
	}
}

// Splitter '-'
type horizontalSpliter struct{}

// Reflect implements Mirror.
func (horizontalSpliter) Reflect(b Beam) []Beam {
	if b.Direction[1] == 0 {
		return []Beam{b}
	}

	return []Beam{
		{
			Position:  b.Position,
			Direction: Vector{1, 0},
		},
		{
			Position:  b.Position,
			Direction: Vector{-1, 0},
		},
	}
}
