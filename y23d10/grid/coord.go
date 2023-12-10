package grid

type Coord [2]int

func (c Coord) Moves() []Coord {
	x, y := c[0], c[1]
	return []Coord{
		{x, y + 1},
		{x, y - 1},
		{x + 1, y},
		{x - 1, y},
	}
}
