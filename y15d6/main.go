package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y15d6/input.txt"
var samplePath = "./y15d6/sample.txt"

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
	fmt.Printf("Time: %vs\n", d.Seconds())
}

func toLight(s string) Light {
	xy := strings.Split(s, ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	return Light{x, y}
}

func part1(input []string) any {
	instructions := []Instruction{}

	for _, line := range input {
		words := strings.Split(line, " ")

		switch words[0] {
		case "toggle":
			start := toLight(words[1])
			end := toLight(words[3])
			instructions = append(instructions, Instruction{
				Bound{start, end}, -1, true,
			})

		case "turn":
			start := toLight(words[2])
			end := toLight(words[4])
			state := 0
			if words[1] == "on" {
				state = 1
			}
			instructions = append(instructions, Instruction{
				Bound{start, end},
				state,
				false,
			})

		default:
			panic(line)
		}
	}

	ans := 0
	n := len(instructions)
	for i := 0; i < 1_000; i++ {
		for j := 0; j < 1_000; j++ {
			state := instructions[n-1].Trace(Light{i, j}, instructions[:n-1])
			ans += state
		}
	}

	return ans
}

func part2(input []string) any {
	instructions := []Instruction{}

	for _, line := range input {
		words := strings.Split(line, " ")

		switch words[0] {
		case "toggle":
			start := toLight(words[1])
			end := toLight(words[3])
			instructions = append(instructions, Instruction{
				Bound{start, end}, -1, true,
			})

		case "turn":
			start := toLight(words[2])
			end := toLight(words[4])
			state := -1
			if words[1] == "on" {
				state = 1
			}
			instructions = append(instructions, Instruction{
				Bound{start, end},
				state,
				false,
			})

		default:
			panic(line)
		}
	}

	ans := 0
	for i := 0; i < 1_000; i++ {
		for j := 0; j < 1_000; j++ {
			state := 0
			for _, next := range instructions {
				state = next.Apply(Light{i, j}, state)
			}
			ans += state
		}
	}

	return ans
}
