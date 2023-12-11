package main

type Coord [2]int

type Universe struct {
	Galaxies  map[Coord]interface{}
	Expansion [2]map[int]interface{} // [dimension][i]
}

func NewUniverse() Universe {
	return Universe{
		Galaxies: make(map[Coord]interface{}),
		Expansion: [2]map[int]interface{}{
			make(map[int]interface{}),
			make(map[int]interface{}),
		},
	}
}

func (u Universe) Distance(a, b Coord, scale int) int {
	d := [2]int{0, 0}
	for i := range d {
		// regular space
		d[i] = Abs(a[i] - b[i])

		// Get range of dimesion
		min, max := Min(a[i], b[i]), Max(a[i], b[i])

		// Add expanded space
		for ii := min; ii <= max; ii++ {
			if _, expanded := u.Expansion[i][ii]; expanded {
				d[i] += scale - 1
			}
		}
	}

	return d[0] + d[1]
}
