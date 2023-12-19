package gears

import "strings"

type Workflow struct {
	Id    string
	Rules []Rule
}

func ParseWorkflow(s string) Workflow {
	s = strings.TrimRight(s, "}")
	parts := strings.Split(s, "{")

	id := parts[0]
	rules := []Rule{}

	for _, ruleString := range strings.Split(parts[1], ",") {
		rule := ParseRule(ruleString)
		rules = append(rules, rule)
	}

	return Workflow{id, rules}
}

func (w Workflow) Execute(gear Gear) string {
	for _, rule := range w.Rules {
		ok, target := rule.Execute(gear)

		if ok {
			return target
		}
	}
	panic("Workflow did not terminate")
}

func (w Workflow) ExecuteCombo(in ComboGear) []ComboGear {
	queue := []ComboGear{in}
	out := []ComboGear{}

	for _, rule := range w.Rules {
		remainder := []ComboGear{}
		for _, next := range queue {
			result := rule.ExecuteCombo(next)

			out = append(out, result[0])
			remainder = append(remainder, result[1:]...)
		}
		queue = remainder
	}

	return out
}
