package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

// var inputPath = "./input.txt"

var inputPath = "./y15d7/input.txt"
var samplePath = "./y15d7/sample.txt"

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
		commands := []Command{}
		for _, line := range input {
			commands = append(commands, ParseCmd(line))
		}

		p1, p2 = part1(commands), part2(commands)
	})
	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

func part1(input []Command) any {
	register := NewRegister(input...)

	return register.Value("a")
}

func part2(input []Command) any {
	register := NewRegister(input...)

	a := register.Value("a")

	register = NewRegister(input...)
	register.Add(write{staticWire(a), dest("b")})

	return register.Value("a")
}
