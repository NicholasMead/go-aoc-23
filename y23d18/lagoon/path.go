package lagoon

type Path []Vector

func (p Path) IsClosed() bool {
	n := len(p)
	return p[0] == p[n-1]
}

func (p Path) Area() int {
	if len(p) == 0 {
		panic("Zero length path cannot have area")
	}

	if !p.IsClosed() {
		p = append(p, p[0])
	}

	//Shoelace formula
	area := 0
	perimeter := 0

	for i := 0; i < len(p)-1; i++ {
		left, right := p[i], p[i+1]
		dx, dy := abs(left[0]-right[0]), abs(left[1]-right[1])

		area += left[0]*right[1] - left[1]*right[0]
		perimeter += dx + dy
	}

	return (area+perimeter)/2 + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
