package volcano

type Beam struct {
	Position  Vector
	Direction Vector
}

func (b Beam) Move() Beam {
	return Beam{
		b.Position.Add(b.Direction),
		b.Direction,
	}
}
