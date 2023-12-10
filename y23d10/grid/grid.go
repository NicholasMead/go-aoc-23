package grid

type Grid [][]rune

func (g Grid) Dimension() (x, y int) {
	x, y = len(g), len(g[0])
	return
}

func (g Grid) At(c Coord) rune {
	x, y := c[0], c[1]
	return g[x][y]
}

func EmptyGrid(x, y int) Grid {
	g := make([][]rune, x)
	for xx := range g {
		g[xx] = make([]rune, y)
	}
	return g
}

func FromInput(from []string) (Grid, Coord) {
	X, Y := len(from[0]), len(from)
	G := EmptyGrid(X, Y)
	S := Coord{}

	// populate grid
	for y, line := range from {
		for x, r := range line {
			G[x][y] = r

			if r == 'S' {
				S = Coord{x, y}
			}
		}
	}

	return G, S
}
