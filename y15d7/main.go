package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./input.txt"
// var inputPath = "./y15d7/input.txt"
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
	register := Register{}
	commands := append([]Command{}, input...)
	count := 0

	for len(commands) > 0 {
		cmd := commands[0]
		commands = commands[1:]
		count++

		if !cmd.TryExecute(register) {
			commands = append(commands, cmd)
		}
	}

	return register["a"]
}

func part2(input []Command) any {
	register := Register{}
	commands := append([]Command{}, input...)

	for len(commands) > 0 {
		cmd := commands[0]
		commands = commands[1:]

		if !cmd.TryExecute(register) {
			commands = append(commands, cmd)
		}
	}

	register = Register{
		"b": register["a"],
	}
	commands = append([]Command{}, input...)
	for i, cmd := range commands {
		switch c := cmd.(type) {
		case write:
			if c.dest == "b" {
				slices.Delete(commands, i, i+1)
			}
			break
		}
	}

	for len(commands) > 0 {
		cmd := commands[0]
		commands = commands[1:]

		if !cmd.TryExecute(register) {
			commands = append(commands, cmd)
		}
	}
	return register["a"]
}
