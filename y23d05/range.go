package main

type Range struct {
	start Index
	len   int
}

func (r Range) Split(at int) [2]Range {
	if at <= 0 || at > r.len {
		panic("Out of range")
	}

	return [2]Range{
		{r.start, at},
		{r.start + Index(at), r.len - at},
	}
}

func (r Range) Cap() Index {
	return r.start + Index(r.len)
}

func (r Range) Expand() []Index {
	indexes := make([]Index, r.len)
	for i := range indexes {
		indexes[i] = r.start + Index(i)
	}
	return indexes
}

func (r Range) Intersect(at Range) (intersect Range, remainder []Range) {
	// total underflow / overflow (trivial)
	if r.Cap() <= at.start || at.Cap() <= r.start {
		remainder = append(remainder, r)
		return
	}

	// partial underflow
	if r.start < at.start {
		diff := int(at.start - r.start)
		remainder = append(remainder, Range{
			r.start,
			diff,
		})
		r.start += Index(diff)
		r.len -= diff
	}

	// partial overflow
	if r.Cap() > at.Cap() {
		diff := int(r.Cap() - at.Cap())
		remainder = append(remainder, Range{
			at.Cap(),
			diff,
		})
		r.len -= diff
	}

	return r, remainder
}
