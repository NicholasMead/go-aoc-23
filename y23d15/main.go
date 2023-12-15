package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d15/hascii"
)

var inputPath = "./y23d15/input.txt"
var samplePath = "./y23d15/sample.txt"

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
		input := strings.Split(inputFile.ReadInputFile(path)[0], ",")
		p1, p2 = part1(input), part2(input)
	})

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

func part1(input []string) any {
	ans := 0

	for _, code := range input {
		ans += hascii.Hash(code)
	}

	return ans
}

func part2(input []string) any {
	boxes := hascii.Hashmap{}

	for _, code := range input {
		isAddOperation := strings.ContainsRune(code, '=')
		if isAddOperation {
			parts := strings.Split(code, "=")
			focalLength, _ := strconv.Atoi(parts[1])
			lens := hascii.Lens{
				Label:       parts[0],
				FocalLength: focalLength,
			}

			boxes.Add(lens)
		}

		isDeleteOperation := strings.ContainsRune(code, '-')
		if isDeleteOperation {
			label := strings.TrimRight(code, "0")
			boxes.Delete(label)
		}
	}

	ans := 0
	for b, box := range boxes {
		for l, lens := range box {
			ans += (b + 1) * (l + 1) * lens.FocalLength
		}
	}
	return ans
}
