package platform

import "errors"

type Rock interface {
	Position() Vector
	Move(Vector) error
	Type() RockType
}

type RockType rune

const (
	RoundRock  RockType = 'O'
	SquareRock RockType = '#'
	RockTypes  string   = "O#"
)

func makeRock(platform *platform, position Vector, rockType RockType) Rock {
	rock := rock{
		platform,
		position,
		rockType,
	}

	switch rockType {
	case RoundRock:
		return roundRock{rock}

	case SquareRock:
		return squareRock{rock}
	}

	panic(string(rockType))
}

type rock struct {
	platform *platform
	position Vector
	rockType RockType
}

func (r rock) Position() Vector {
	return r.position
}

func (r rock) Type() RockType {
	return r.rockType
}

type roundRock struct {
	rock
}

func (r roundRock) Move(v Vector) error {
	nextPos := r.position.Add(v)

	inBounds := true &&
		nextPos[0] >= 0 &&
		nextPos[1] >= 0 &&
		nextPos[0] < r.platform.width &&
		nextPos[1] < r.platform.height
	if !inBounds {
		return errors.New("out of bounds")
	}

	_, taken := r.platform.rocks[nextPos]
	if taken {
		return errors.New("path blocked")
	}

	r.platform.rocks[nextPos] = roundRock{
		rock: rock{
			r.platform,
			r.position.Add(v),
			r.rockType,
		},
	}
	delete(r.platform.rocks, r.position)
	return nil
}

type squareRock struct {
	rock
}

func (r squareRock) Move(v Vector) error {
	return errors.ErrUnsupported
}
