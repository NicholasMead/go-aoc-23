package machine

import "strings"

type Schema struct {
	InLabel   string
	Module    Module
	OutLabels []string
}

func ParseSchema(line string) Schema {
	schema := Schema{
		"",
		nil,
		[]string{},
	}

	parts := strings.Split(line, " -> ")

	switch parts[0][0] {
	case '%':
		schema.InLabel = parts[0][1:]
		schema.Module = &FlipFlop{}
	case '&':
		schema.InLabel = parts[0][1:]
		schema.Module = &Conjunction{}

	default:
		schema.InLabel = parts[0]
		schema.Module = Broadcast{}
	}

	schema.OutLabels = strings.Split(parts[1], ", ")

	return schema
}
