package volcano

type Cave struct {
	height, width int
	mirrors       map[Vector]Mirror
}

func NewCave(height, width int) Cave {
	return Cave{
		height:  height,
		width:   width,
		mirrors: make(map[Vector]Mirror),
	}
}

func (c *Cave) AddMirror(pos Vector, mirror Mirror) {
	c.mirrors[pos] = mirror
}

func (c *Cave) EnrichedFromBeam(beam Beam) map[Vector]int {
	queue := []Beam{beam}
	memo := map[Beam]bool{}
	enriched := map[Vector]int{}

	duplicates, exits := 0,0

	for len(queue) > 0 {
		beam, queue = queue[0], queue[1:]
		beam = beam.Move()

		// Boundry conditions
		if !c.inBounds(beam.Position) {
			exits++
			continue // out of bounds
		}
		if _, found := memo[beam]; found {
			duplicates ++
			continue // duplicate
		}

		// Process
		enriched[beam.Position] += 1
		memo[beam] = true

		// Expand
		if mirror, found := c.mirrors[beam.Position]; found {
			reflections := mirror.Reflect(beam)
			queue = append(queue, reflections...)
		} else {
			queue = append(queue, beam)
		}
	}

	return enriched
}

func (c Cave) inBounds(pos Vector) bool {
	extent := Vector{c.width, c.height}
	for i := range extent {
		if pos[i] < 0 || pos[i] >= extent[i] {
			return false
		}
	}
	return true
}
