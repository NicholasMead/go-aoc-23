package main

import (
	"strconv"
)

type Wire interface {
	Value(ReadRegister) (value Value)
}

func ParseWire(s string) Wire {
	if num, err := strconv.Atoi(s); err == nil {
		return staticWire(num)
	} else {
		return dynamicWire(s)
	}
}

type staticWire int

func (static staticWire) Value(_ ReadRegister) Value {
	return Value(static)
}

type dynamicWire string

func (dynamic dynamicWire) Value(r ReadRegister) Value {
	return r.Value(string(dynamic))
}
