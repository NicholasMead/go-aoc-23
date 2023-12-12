package main

import (
	"strconv"
	"strings"
)

type Command interface {
	Wire
	Dest() string
}

func ParseCmd(line string) Command {
	parts := strings.Split(line, " -> ")
	input, dest := parts[0], dest(parts[1])

	inParts := strings.Split(input, " ")
	switch len(inParts) {
	case 1:
		return write{ParseWire(inParts[0]), dest}

	case 2:
		return not{ParseWire(inParts[1]), dest}

	case 3:
		switch inParts[1] {
		case "AND":
			return and{
				ParseWire(inParts[0]), ParseWire(inParts[2]),
				dest,
			}
		case "OR":
			return or{
				ParseWire(inParts[0]), ParseWire(inParts[2]),
				dest,
			}
		case "LSHIFT":
			i, _ := strconv.Atoi(inParts[2])
			return lShift{
				ParseWire(inParts[0]), i,
				dest,
			}
		case "RSHIFT":
			i, _ := strconv.Atoi(inParts[2])
			return rShift{
				ParseWire(inParts[0]), i,
				dest,
			}
		}
	}
	panic(line)
}

type dest string

func (w dest) Dest() string {
	return string(w)
}

type write struct {
	a Wire
	dest
}

func (cmd write) Value(r ReadRegister) Value {
	return cmd.a.Value(r)
}

type and struct {
	a, b Wire
	dest
}

func (cmd and) Value(r ReadRegister) Value {
	a := cmd.a.Value(r)
	b := cmd.b.Value(r)

	return a & b
}

type or struct {
	a, b Wire
	dest
}

func (cmd or) Value(r ReadRegister) Value {
	aa := cmd.a.Value(r)
	bb := cmd.b.Value(r)

	return aa | bb
}

type lShift struct {
	src Wire
	i   int
	dest
}

func (cmd lShift) Value(r ReadRegister) Value {
	value := cmd.src.Value(r)
	return value << Value(cmd.i)
}

type rShift struct {
	src Wire
	i   int
	dest
}

func (cmd rShift) Value(r ReadRegister) Value {
	value := cmd.src.Value(r)
	return value >> Value(cmd.i)
}

type not struct {
	src Wire
	dest
}

func (cmd not) Value(r ReadRegister) Value {
	value := cmd.src.Value(r)
	return value ^ 0xffff
}
