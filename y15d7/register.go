package main

type Register interface {
	ReadRegister
	Add(Command)
}

type ReadRegister interface {
	Value(of string) Value
}

type Value uint16

type register struct {
	cache map[string]Value
	wires map[string]Wire
}

func NewRegister(commands ...Command) Register {
	r := &register{
		make(map[string]Value),
		make(map[string]Wire),
	}
	for _, cmd := range commands {
		r.Add(cmd)
	}
	return r
}

func (r register) Value(of string) Value {
	if c, ok := r.cache[of]; ok {
		return c
	} else {
		wire := r.wires[of]
		r.cache[of] = wire.Value(r)
		return r.cache[of]
	}
}

func (r *register) Add(cmd Command) {
	r.wires[cmd.Dest()] = cmd
}
