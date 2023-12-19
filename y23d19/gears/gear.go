package gears

import (
	"strconv"
	"strings"
)

type Gear map[Property]int

func (gear Gear) Value() int {
	value := 0
	for _, v := range gear {
		value += v
	}
	return value
}

func (gear Gear) Select(p Property) int {
	return gear[p]
}

func ParseGear(s string) Gear {
	s = strings.Trim(s, "{")
	s = strings.Trim(s, "}")

	gear := Gear{}

	for _, property := range strings.Split(s, ",") {
		parts := strings.Split(property, "=")

		prop := Property(parts[0][0])
		value, _ := strconv.Atoi(parts[1])

		gear[prop] = value
	}

	return gear
}
