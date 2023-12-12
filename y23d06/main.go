package main

import (
	"fmt"
	"math"
	"os"

	"github.com/NicholasMead/go-aoc-23/common"
)

var inputData = [][]Race{{
	{41, 214},
	{96, 1789},
	{88, 1128},
	{94, 1055},
}, {
	{41_968_894, 214_178_911_271_055},
}}
var sampleData = [][]Race{{
	{7, 9},
	{15, 40},
	{30, 200},
}, {
	{71_530, 940_200},
}}

type Race struct {
	Time   int
	Record int
}

func (r Race) Compete(charge int) (dist int, win bool) {
	if charge >= r.Time {
		return
	}

	dist = (r.Time - charge) * charge
	if dist > r.Record {
		win = true
	}
	return
}

func (r Race) Bounds() (low, high int) {
	// quadratic eq
	a := float64(-1)
	b := float64(r.Time)
	c := float64(-r.Record)

	x1 := (-b + math.Sqrt((b*b)-(4*a*c))) / (2 * a)
	x2 := (-b - math.Sqrt((b*b)-(4*a*c))) / (2 * a)

	low = int(math.Ceil(math.Min(x1, x2)))
	high = int(math.Floor(math.Max(x1, x2)))
	return
}

func main() {
	var p1, p2 any = "", ""
	d := common.Timer(func() {
		args := os.Args[1:]
		input := inputData
		if len(args) > 0 {
			switch args[0] {
			case "sample":
				input = sampleData
			case "input":
				input = inputData
			}
		}

		p1, p2 = Calculate(input[0]), Calculate(input[1])
	})
	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vns\n", d.Nanoseconds())
}

func Simulate(input []Race) int {
	ans := 1

	for _, race := range input {
		wins := 0

		for charge := 1; charge < race.Time; charge++ {
			_, win := race.Compete(charge)

			if win {
				wins += 1
			}
		}

		ans *= wins
	}

	return ans
}

func Calculate(input []Race) int {
	var ans int = 1

	for _, race := range input {
		low, high := race.Bounds()

		wins := high - low + 1 // inclusive range
		ans *= wins
	}

	return ans
}
