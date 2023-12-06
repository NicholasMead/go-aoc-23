package main

import (
	"strconv"
	"strings"
)

type Command interface {
	TryExecute(Register) (success bool)
}

func ParseCmd(line string) Command {
	parts := strings.Split(line, " -> ")
	input, dest := parts[0], parts[1]

	inParts := strings.Split(input, " ")
	switch len(inParts) {
	case 1:
		return write{
			src:  ParseWire(inParts[0]),
			dest: dest,
		}

	case 2:
		return not{
			src:  ParseWire(inParts[1]),
			dest: dest,
		}

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

type write struct {
	src  Wire
	dest string
}

func (cmd write) TryExecute(r Register) bool {
	if v, ok := cmd.src.Value(r); ok {
		r[cmd.dest] = v
		return true
	}
	return false
}

type and struct {
	a, b Wire
	dest string
}

func (cmd and) TryExecute(r Register) (ok bool) {
	var aa, bb Value

	if aa, ok = cmd.a.Value(r); !ok {
		return false
	}
	if bb, ok = cmd.b.Value(r); !ok {
		return false
	}

	r[cmd.dest] = aa & bb
	return true
}

type or struct {
	a, b Wire
	dest string
}

func (cmd or) TryExecute(r Register) (ok bool) {
	var aa, bb Value

	if aa, ok = cmd.a.Value(r); !ok {
		return false
	}
	if bb, ok = cmd.b.Value(r); !ok {
		return false
	}

	r[cmd.dest] = aa | bb
	return true
}

type lShift struct {
	src  Wire
	i    int
	dest string
}

func (cmd lShift) TryExecute(r Register) bool {
	if value, ok := cmd.src.Value(r); ok {
		r[cmd.dest] = value << Value(cmd.i)
		return true
	} else {
		return false
	}
}

type rShift struct {
	src  Wire
	i    int
	dest string
}

func (cmd rShift) TryExecute(r Register) bool {
	if value, ok := cmd.src.Value(r); ok {
		r[cmd.dest] = value >> Value(cmd.i)
		return true
	} else {
		return false
	}
}

type not struct {
	src  Wire
	dest string
}

func (cmd not) TryExecute(r Register) bool {
	if value, ok := cmd.src.Value(r); ok {
		r[cmd.dest] = value ^ 0xffff
		return true
	} else {
		return false
	}
}
