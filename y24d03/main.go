package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y24d03/input.txt"
var samplePath = "./y24d03/sample.txt"

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

type instruction interface {
	compute() int
}

type multiply [2]int

func (i multiply) compute() int {
	return i[0] * i[1]
}

func parseMultiply(input string) multiply {
	r := regexp.MustCompile(`\d{1,3}`)

	matches := r.FindAllString(input, -1)

	m := multiply{}
	for i, v := range matches {
		n, _ := strconv.Atoi(v)
		m[i] = n
	}
	return m
}

type enable bool

func (e enable) compute() int {
	if e {
		return 1
	}
	return 0
}

func parseEnable(input string) enable {
	return enable(!strings.HasPrefix(input, "don't"))
}

func parseInstructions(input string) []instruction {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do(n't)?\(\)`)

	matches := r.FindAllString(input, -1)

	instructions := make([]instruction, len(matches))

	for i, m := range matches {
		if strings.HasPrefix(m, "mul") {
			instructions[i] = parseMultiply(m)
		} else if strings.HasPrefix(m, "do") {
			instructions[i] = parseEnable(m)
		}
	}


	return instructions
}

func part1(input []string) any {
	instructions := make([]instruction, 0)

	for _, line := range input {
		instructions = append(instructions, parseInstructions(line)...)
	}

	answer := 0
	for _, i := range instructions {
		switch i.(type) {
		case multiply:
			answer += i.compute()
		}
	}

	return answer
}

func part2(input []string) any {
	instructions := make([]instruction, 0)

	for _, line := range input {
		instructions = append(instructions, parseInstructions(line)...)
	}

	answer := 0
	enabled := 1
	for _, i := range instructions {
		switch i.(type) {
		case enable:
			enabled = i.compute()
			break
		case multiply:
			answer += enabled * i.compute()
			break
		}
	}

	return answer
}
