package crucible

type Vector [2]int

func (a Vector) Add(b Vector) Vector {
	return Vector{
		a[0] + b[0],
		a[1] + b[1],
	}
}

func (a Vector) Rotate(turns int) Vector {
	out := a
	for turns < 0 {
		turns = turns + 4
	}
	for i := 0; i < turns; i++ {
		out = Vector{
			-out[1],
			out[0],
		}
	}
	return out
}

func (a Vector) Dist(b Vector) int {
	dist := 0
	for i := range a {
		dist += abs(a[i] - b[i])
	}
	return dist
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

type Grid map[Vector]int

func (g Grid) Max() Vector {
	max := Vector{}
	for pos := range g {
		if pos[0] > max[0] || pos[1] > max[1] {
			max = pos
		}
	}
	return max
}

type Crucible struct {
	Position Vector
	Velocity Vector
	Distance int

	min, max int
}

func NewCrucible(min, max int) Crucible {
	return Crucible{
		Position: Vector{},
		Velocity: Vector{1, 0},
		Distance: 0,
		min:      min,
		max:      max,
	}
}

func (c Crucible) Moves() (moves []Crucible) {
	if c.Distance < c.max {
		moves = append(moves,
			Crucible{
				c.Position.Add(c.Velocity),
				c.Velocity,
				c.Distance + 1,
				c.min, c.max,
			},
		)
	}

	if c.Distance >= c.min {
		moves = append(moves,
			Crucible{
				c.Position.Add(c.Velocity.Rotate(1)),
				c.Velocity.Rotate(1),
				1,
				c.min, c.max,
			},
			Crucible{
				c.Position.Add(c.Velocity.Rotate(-1)),
				c.Velocity.Rotate(-1),
				1,
				c.min, c.max,
			},
		)
	}
	return moves
}

func (c Crucible) CanStop() bool {
	return c.Distance >= c.min
}
