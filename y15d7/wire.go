package main

import "strconv"

type Wire interface {
	Value(Register) (value Value, ok bool)
}

func ParseWire(s string) Wire {
	if num, err := strconv.Atoi(s); err == nil {
		return staticWire(num)
	} else {
		return dynamicWire(s)
	}
}

type staticWire Value

func (static staticWire) Value(_ Register) (Value, bool) {
	return Value(static), true
}

type dynamicWire string

func (dynamics dynamicWire) Value(r Register) (Value, bool) {
	v, ok := r[string(dynamics)]
	return v, ok
}
