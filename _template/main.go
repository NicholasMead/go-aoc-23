package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string

	if file, err := os.OpenFile("./input.txt"); err != nil {
		bufio.NewReader(file).rea

	}
	input, err := bufio.Reader(os.ReadFile("./input.txt"))
	if err != nil {
		panic(err)
	}

	fmt.Println(input)
	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input any) any {
	return "_"
}

func part2(input any) any {
	return "_"
}
