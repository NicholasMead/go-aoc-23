package volcano

type Vector [2]int

func (a Vector) Add(b Vector) (out Vector) {
	for i := range out {
		out[i] = a[i] + b[i]
	}
	return
}
