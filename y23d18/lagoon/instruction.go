package lagoon

import (
	"strconv"
	"strings"
)

type Instruction struct {
	Distance  int
	Direction Vector
	Colour    string
}

// Performs the 'debug' operation in day 18 part 2
func (i Instruction) Debug() Instruction {
	var (
		distance  int64
		direction Vector
	)

	switch i.Colour[5] {
	case '0':
		direction = Vector{1, 0}
	case '1':
		direction = Vector{0, 1}
	case '2':
		direction = Vector{-1, 0}
	case '3':
		direction = Vector{0, -1}
	}

	distance, _ = strconv.ParseInt(i.Colour[0:5], 16, 32)

	return Instruction{
		Distance:  int(distance),
		Direction: direction,
	}
}

func (i Instruction) Vector() Vector {
	return i.Direction.Scale(i.Distance)
}

func InstructionFromInput(s string) Instruction {
	var (
		parts     = strings.Split(s, " ")
		distance  int
		direction Vector
		colour    string
	)

	// Distance
	distance, _ = strconv.Atoi(parts[1])

	// Direction
	switch parts[0] {
	case "U":
		direction = Vector{+0, -1}
	case "R":
		direction = Vector{+1, +0}
	case "D":
		direction = Vector{+0, +1}
	case "L":
		direction = Vector{-1, +0}
	default:
		panic(parts[0])
	}

	// Colour
	colour = strings.TrimLeft(parts[2], "(#")
	colour = strings.TrimRight(colour, ")")

	return Instruction{
		distance,
		direction,
		colour,
	}
}
