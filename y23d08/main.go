package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y23d8/input.txt"
var samplePath = "./y23d8/sample.txt"

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
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

type Node [2]string

func part1(input []string) any {
	instructions := []rune{}
	network := map[string]Node{}

	for _, r := range input[0] {
		instructions = append(instructions, r)
	}

	for _, line := range input[2:] {
		base := line[0:3]
		left := line[7:10]
		right := line[12:15]

		network[base] = Node{left, right}
	}

	return period(instructions, network, "AAA")
}

func period(instructions []rune, network map[string]Node, start string) int {
	at := start

	for count := 1; count < 100_000_000; count++ {
		next := instructions[0]
		instructions = append(instructions[1:], next)

		switch next {
		case 'L':
			at = network[at][0]
		case 'R':
			at = network[at][1]
		default:
			panic(next)
		}

		if at[2] == 'Z' {
			return count
		}
	}

	panic("Overflow")
}

func part2(input []string) any {
	instructions := []rune{}
	network := map[string]Node{}

	for _, r := range input[0] {
		instructions = append(instructions, r)
	}

	for _, line := range input[2:] {
		base := line[0:3]
		left := line[7:10]
		right := line[12:15]

		network[base] = Node{left, right}
	}

	cycles := []int{}

	for node := range network {
		if node[2] == 'A' {
			p := period(instructions, network, node)

			cycles = append(cycles, p)
		}
	}

	return lowerCommonMultiple(cycles)
}

func lowerCommonMultiple(values []int) int {
	primes := map[int]interface{}{}

	for _, value := range values {
		factors := primeFactors(value)
		for _, f := range factors {
			primes[f] = struct{}{}
		}
	}

	lcm := 1
	for prime := range primes {
		lcm *= prime
	}
	return lcm
}

func primeFactors(value int) (primes []int) {
	candidates := factors(value)

	for _, c := range candidates {
		if len(factors(c)) == 2 {
			primes = append(primes, c)
		}
	}

	return
}

func factors(value int) (fact []int) {
	if value <= 0 {
		return
	}

	if value == 1 {
		return []int{1}
	} else {
		fact = append(fact, 1)

		for i := 2; i <= value/2; i++ {
			if value%i == 0 {
				fact = append(fact, i)
			}
		}

		fact = append(fact, value)
		return
	}
}
