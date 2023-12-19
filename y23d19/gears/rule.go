package gears

import (
	"strconv"
	"strings"
)

type Rule struct {
	Property  Property
	Operation Operation
	Value     int
	Target    string
}

type Operation rune

const (
	GreaterThan Operation = '>'
	LessThan    Operation = '<'
	Goto        Operation = 0
)

func (r Rule) Execute(gear Gear) (ok bool, target string) {
	if r.Operation == Goto {
		return true, r.Target
	}

	switch r.Operation {
	case GreaterThan:
		ok = gear.Select(r.Property) > r.Value
	case LessThan:
		ok = gear.Select(r.Property) < r.Value
	default:
		panic("unknown operation: " + string(r.Operation))
	}

	if ok {
		return ok, r.Target
	} else {
		return ok, ""
	}
}

func (r Rule) ExecuteCombo(combo ComboGear) []ComboGear {
	if r.Operation == Goto {
		copy := combo.Copy()
		copy.Location = r.Target
		return []ComboGear{copy}
	}

	var ok, remainder ComboGear

	switch r.Operation {
	case GreaterThan:
		remainder, ok = combo.Split(r.Property, r.Value+1)
	case LessThan:
		ok, remainder = combo.Split(r.Property, r.Value)
	default:
		panic("unknown operation: " + string(r.Operation))
	}

	ok.Location = r.Target
	return []ComboGear{ok, remainder}
}

func ParseRule(s string) Rule {
	parts := strings.Split(s, ":")

	isGoto := len(parts) == 1
	if isGoto {
		return Rule{
			Operation: Goto,
			Target:    parts[0],
		}
	}

	property := Property(parts[0][0])
	operation := Operation(parts[0][1])
	value, _ := strconv.Atoi(parts[0][2:])

	return Rule{
		Property:  property,
		Operation: operation,
		Value:     value,
		Target:    parts[1],
	}
}
