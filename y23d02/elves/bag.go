package elves

type Bag struct {
	Red, Green, Blue int
}

func (bag Bag) CanDraw(red, green, blue int) bool {
	return red <= bag.Red && green <= bag.Green && blue <= bag.Blue
}

func (b Bag) Power() int {
	return b.Red * b.Green * b.Blue
}
