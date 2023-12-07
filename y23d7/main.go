package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y23d7/input.txt"
var samplePath = "./y23d7/sample.txt"

type bid struct {
	hand  Hand
	value int
}

func main() {
	var p1, p2 any = "", ""
	d := common.Timer(func() {
		args := os.Args[1:]
		path := inputPath
		if len(args) > 0 {
			switch args[0] {
			case "sample":
				path = samplePath
			case "input":
				path = inputPath
			default:
				path = args[1]
			}
		}

		input := inputFile.ReadInputFile(path)

		p1, p2 = part1(input), part2(input)
	})
	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vms\n", d.Milliseconds())
}

func part1(input []string) any {
	bids := []bid{}

	for _, line := range input {
		parts := strings.Split(line, " ")
		bid := bid{}
		bid.hand = NewHand(parts[0])
		bid.value, _ = strconv.Atoi(parts[1])
		bids = append(bids, bid)
	}

	cmp := func(a, b bid) int {
		return HandCompare(a.hand, b.hand)
	}

	slices.SortStableFunc(bids, cmp)

	ans := 0
	for i, bid := range bids {
		ans += (i + 1) * bid.value
	}
	return ans
}

func part2(input []string) any {
	bids := []bid{}

	for _, line := range input {
		parts := strings.Split(line, " ")
		bid := bid{}
		bid.hand = NewWildHand(parts[0])
		bid.value, _ = strconv.Atoi(parts[1])
		bids = append(bids, bid)
	}

	cmp := func(a, b bid) int {
		return HandCompare(a.hand, b.hand)
	}

	slices.SortStableFunc(bids, cmp)

	ans := 0
	for i, bid := range bids {
		ans += (i + 1) * bid.value
	}
	return ans
}
