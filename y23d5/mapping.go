package main

type Mapping struct {
	dest, src Index
	len       int
}

func (m Mapping) MapIndex(from Index) (Index, bool) {
	if from < m.src || from >= m.src+Index(m.len) {
		return from, false
	}

	return from + (m.dest - m.src), true
}

func (m Mapping) MapRange(from Range) (output Range, remainder []Range) {
	output, remainder = from.Intersect(Range{m.src, m.len})

	// if intersecting, do mapping
	if output.len > 0 {
		output.start += (m.dest - m.src)
	}

	return
}
