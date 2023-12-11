package main

func Abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func Max(a ...int) int {
	max := a[0]

	for _, aa := range a[1:] {
		if aa > max {
			max = aa
		}
	}

	return max
}

func Min(a ...int) int {
	min := a[0]

	for _, aa := range a[1:] {
		if aa < min {
			min = aa
		}
	}

	return min
}
