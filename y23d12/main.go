package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d12/springs"
)

var inputPath = "./y23d12/input.txt"
var samplePath = "./y23d12/sample.txt"

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
		p1, p2 = part1(input), 0 //part2(input)
	})

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

func part1(input []string) any {
	ans := 0

	for l, line := range input {
		lineParts := strings.Split(line, " ")

		history := springs.History(lineParts[0])
		record := springs.NewRecordFromString(lineParts[1])

		counter := springs.Counter{}
		counter.History = history
		counter.Record = record

		count := counter.Count()

		ans += count
		fmt.Printf("(%v/%v): %v [ans = %v]\n", l+1, len(input), count, ans)
	}

	return ans
}

func part2(input []string) any {
	ans := 0

	for l, line := range input {
		lineParts := strings.Split(line, " ")
		recordParts := strings.Split(lineParts[1], ",")

		history := springs.History(lineParts[0]).Normalise()
		record := make(springs.Record, len(recordParts))

		for i, r := range recordParts {
			record[i], _ = strconv.Atoi(r)
		}

		history, record = springs.Unfold(history, record)

		fmt.Printf("\n%v\n%v\n", history, record.DrawHistory())

		count := springs.NewHistoryFinder(record).CountValid(history)
		ans += count

		fmt.Printf("(%v/%v): %v [ans = %v]\n", l+1, len(input), count, ans)
	}

	return ans
}