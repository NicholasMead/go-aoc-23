package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
)

var inputPath = "./y15d4/input.txt"
var samplePath = "./y15d4/sample.txt"

func CalcMd5(input string, key int) string {
	input = fmt.Sprint(input, key)

	hash := md5.Sum([]byte(input))
	return fmt.Sprintf("%0x", hash)
}

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

		input := inputFile.ReadInputFile(path)[0]

		p1, p2 = part1(input), part2(input)
	})
	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vms\n", d.Milliseconds())
}

func part1(input string) any {
	for key := 0; key < math.MaxInt; key++ {
		hash := CalcMd5(input, key)
		if strings.HasPrefix(hash, "00000") {
			return key
		}
	}
	return -1
}

func part2(input string) any {
	for key := 0; key < math.MaxInt; key++ {
		hash := CalcMd5(input, key)
		if strings.HasPrefix(hash, "000000") {
			return key
		}
	}
	return -1
}
