package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d14/platform"
)

var inputPath = "./y23d14/input.txt"
var samplePath = "./y23d14/sample.txt"

func main() {
	var p1, p2 any = "", ""
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

	d := common.Timer(func() {
		input := inputFile.ReadInputFile(path)
		p1, p2 = part1(input), part2(input)
	})

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vs\n", d.Seconds())
}

func NorthWeight(p platform.Platform) int {
	ans := 0
	for _, rock := range p.GetRocksOfType(platform.RoundRock) {
		ans += p.Height() - rock[1]
	}
	return ans
}

func part1(input []string) any {
	p := platform.LoadPlatform(input)
	p.ApplyTilt(platform.Vector{0, -1})
	return NorthWeight(p)
}

func detectCycle(p platform.Platform) (offset int, period int) {
	memo := map[string]int{
		p.Hash(): 0,
	}

	for i := 1; i < 1_000_000_000; i++ {
		p.SpinCycle()
		hash := p.Hash()

		if index, found := memo[hash]; found {
			return index, i - index
		} else {
			memo[hash] = i
		}
	}

	panic("No cycle found")
}

func part2(input []string) any {
	p := platform.LoadPlatform(input)

	offset, cycle := detectCycle(p)

	current := offset + cycle
	cap := 1_000_000_000

	//sudo-iterate to close to the cap...
	for current+cycle <= cap {
		current += cycle
	}

	//iterate the last bit
	for i := 0; i < cap-current; i++ {
		p.SpinCycle()
	}

	return NorthWeight(p)
}
