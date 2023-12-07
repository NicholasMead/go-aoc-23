package main

type Light [2]int // [x,y]

type Bound [2]Light // [[start_x, start_y], [end_x, end_y]]

func (light Light) InRange(bound Bound) bool {
	for dim, pos := range light {
		if pos < bound[0][dim] || pos > bound[1][dim] {
			return false
		}
	}
	return true
}

type Instruction struct {
	Bound  Bound
	Set    int
	Toggle bool
}

// Back traces the light state
func (i Instruction) Trace(pos Light, prev []Instruction) int {
	n := len(prev)

	if !pos.InRange(i.Bound) {
		if n == 0 {
			return 0
		}
		return prev[n-1].Trace(pos, prev[:n-1])
	}

	if i.Toggle {
		if n == 0 {
			return 0
		}
		return 1 - prev[len(prev)-1].Trace(pos, prev[:n-1])
	}

	return i.Set
}

// applies the instruction
func (i Instruction) Apply(pos Light, state int) int {
	if !pos.InRange(i.Bound) {
		return state
	}

	newState := state

	if i.Toggle {
		newState += 2
	} else {
		newState += i.Set
	}

	if newState < 0 {
		return 0
	} else {
		return newState
	}
}
