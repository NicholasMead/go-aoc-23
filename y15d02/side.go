package main

type Side [2]int

func (s Side) Area() int {
	return s[0] * s[1]
}

func (s Side) Perimeter() int {
	return 2 * (s[0] + s[1])
}
