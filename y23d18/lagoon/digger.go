package lagoon

type Digger interface {
	Position() Vector
	Path() Path

	Follow(Instruction)
}

type digger struct {
	position Vector
	path     []Vector
}

// Path implements Digger.
func (d *digger) Path() Path {
	return append(Path{}, d.path...)
}

// Follow implements Digger.
func (d *digger) Follow(i Instruction) {
	d.position = d.position.Add(i.Vector())
	d.path = append(d.path, d.position)
}

// Position implements Digger.
func (d *digger) Position() Vector {
	return d.position
}

func NewDigger(start Vector) Digger {
	return &digger{
		position: start,
		path:     []Vector{start},
	}
}
