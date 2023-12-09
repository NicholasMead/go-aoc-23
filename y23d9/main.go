package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y23d9/input.txt"
var samplePath = "./y23d9/sample.txt"

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

		input := parse(inputFile.ReadInputFile(path))

		p1, p2 = part1(input), part2(input)
	})
	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

func parse(raw []string) [][]int {
	input := [][]int{}

	for _, line := range raw {
		s := []int{}

		for _, word := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(word)
			s = append(s, num)
		}

		input = append(input, s)
	}

	return input
}

func diff(s []int) []int {
	diff := make([]int, len(s)-1)

	for i := 0; i < len(s)-1; i++ {
		diff[i] = s[i+1] - s[i]
	}

	return diff
}

func part1(input [][]int) any {
	// compute all sub sequences (inc zeros)
	// 1   3   6  10  15  21
	// 2   3   4   5   6
	// 1   1   1   1
	// 0   0   0
	// 0   0
	// 0
	// buttom up next elements
	// 1   3   6  10  15  21  _28_
	// 2   3   4   5   6 _7_
	// 1   1   1   1  _1_
	// 0   0   0  _0_
	// 0   0  _0_
	// 0  _0_
	ans := 0

	for _, s := range input {
		var (
			memo = [][]int{append([]int{}, s...)}
			n    = len(s)
		)

		// compute all sub sequences (inc zeros)
		for i := 0; i < n-1; i++ {
			memo = append(memo, diff(memo[i]))
		}

		// buttom up next elements
		next := 0
		for i := n - 1; i >= 0; i-- {
			next = memo[i][n-i-1] + next
		}

		ans += next
	}

	return ans
}

func part2(input [][]int) any {
	// compute all sub sequences (inc zeros)
	// 1   3   6  10  15  21
	// 2   3   4   5   6
	// 1   1   1   1
	// 0   0   0
	// 0   0
	// 0
	// buttom up prev elements
	// _0_ 1   3   6  10  15  21
	// _1_ 2   3   4   5   6
	// _1_ 1   1   1   1
	// _0_ 0   0   0
	// _0_ 0   0
	// _0_ 0
	ans := 0

	for _, s := range input {
		var (
			n    = len(s)
			memo = [][]int{append([]int{}, s...)}
		)

		// compute all sub sequences (inc zeros)
		for i := 0; i < n-1; i++ {
			memo = append(memo, diff(memo[i]))
		}

		// buttom up prev elements
		next := 0
		for i := n - 1; i >= 0; i-- {
			next = memo[i][0] - next
		}

		ans += next
	}

	return ans
}
