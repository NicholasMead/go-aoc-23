package main

import (
	"fmt"
	"os"

	"github.com/NicholasMead/go-aoc-23/common"
	"github.com/NicholasMead/go-aoc-23/common/inputFile"
	"github.com/NicholasMead/go-aoc-23/y23d19/gears"
)

var inputPath = "./y23d19/input.txt"
var samplePath = "./y23d19/sample.txt"

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
		workflow, inGears := parse(input)
		p1, p2 = part1(workflow, inGears), part2(workflow)
	})

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
	fmt.Printf("Time: %vus\n", d.Microseconds())
}

func parse(input []string) (map[string]gears.Workflow, []gears.Gear) {
	workflows := map[string]gears.Workflow{}
	inGears := []gears.Gear{}

	for _, line := range input {
		if line == "" {
			break
		}

		workflow := gears.ParseWorkflow(line)
		workflows[workflow.Id] = workflow
	}

	for _, line := range input[len(workflows)+1:] {
		part := gears.ParseGear(line)
		inGears = append(inGears, part)
	}

	return workflows, inGears
}

func part1(workflows map[string]gears.Workflow, inGears []gears.Gear) any {
	var (
		accepted       = []gears.Gear{}
		accept, reject = "A", "R"
	)

	for _, gear := range inGears {
		next := "in"
		seen := map[string]bool{}

		for !seen[next] {
			seen[next] = true
			next = workflows[next].Execute(gear)

			if next == accept || next == reject {
				if next == accept {
					accepted = append(accepted, gear)
				}
				break
			}
		}

		if next != accept && next != reject {
			panic("Workflow infinite loop")
		}
	}

	ans := 0
	for _, part := range accepted {
		ans += part.Value()
	}
	return ans
}

func part2(workflows map[string]gears.Workflow) any {
	var (
		queue    = []gears.ComboGear{gears.DefaultCombo("in")}
		accepted = []gears.ComboGear{}

		accept, reject = "A", "R"
	)

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		workflow := workflows[next.Location]
		output := workflow.ExecuteCombo(next)

		for _, combo := range output {
			switch combo.Location {
			case accept:
				accepted = append(accepted, combo)
			case reject:
				//drop
			default:
				queue = append(queue, combo)
			}
		}
	}

	ans := 0
	for _, part := range accepted {
		ans += part.Combinations()
	}
	return ans
}
