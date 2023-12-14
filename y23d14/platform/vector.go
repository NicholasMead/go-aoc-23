package platform

type Vector [2]int

func (v Vector) Add(a Vector) Vector {
	sum := Vector{}

	for i := range sum {
		sum[i] = v[i] + a[i]
	}

	return sum
}
