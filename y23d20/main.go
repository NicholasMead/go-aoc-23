package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d20/machine"
)

var inputPath = "./y23d20/input.txt"
var samplePath = "./y23d20/sample.txt"

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
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

func part1(input []string) any {
	schemas := []machine.Schema{}

	for _, line := range input {
		schemas = append(schemas, machine.ParseSchema(line))
	}

	circuit := machine.NewCircuit(schemas)
	high, low := 0, 0

	for i := 0; i < 1_000; i++ {
		signals := circuit.Input("broadcaster", machine.LOW)
		for _, sig := range signals {
			switch sig.Pulse {
			case machine.HIGH:
				high++
			case machine.LOW:
				low++
			}
		}
	}

	return high * low
}

func part2(input []string) any {
	schemas := []machine.Schema{}

	for _, line := range input {
		schemas = append(schemas, machine.ParseSchema(line))
	}

	circuit := machine.NewCircuit(schemas)

	//sub graphs... check input (I don't like it either)
	watch := []string{"gc", "sz", "cm", "xf"}
	period := map[string]int{
		"gc": 0,
		"sz": 0,
		"cm": 0,
		"xf": 0,
	}

	for count := 0; count < 1_000_000; count++ {
		if count%1_000_000 == 0 {
			fmt.Println(count/1_000_000, "Million")
		}

		signals := circuit.Input("broadcaster", machine.LOW)

		for _, s := range signals {
			if s.Pulse == machine.LOW && slices.Contains(watch, s.To) {
				period[s.To] = count

				ans := 1
				for _, p := range period {
					ans *= p
				}
				if ans > 0 {
					return ans
				}
			}

			if s.To == "rx" {
				if s.Pulse == machine.LOW {
					return count
				}
			}
		}
	}

	return 0
}
