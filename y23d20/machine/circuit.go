package machine

type Wire interface {
	Send(Pulse) *Pulse
}

type Module interface {
	Connect(string)
	GetConnection(string) Wire

	Type() string
}

type Circuit struct {
	modules map[string]Module
	routes  map[string][]string
}

type Signal struct {
	From, To string
	Pulse    Pulse
}

func NewCircuit(schemas []Schema) Circuit {
	circuit := Circuit{
		modules: make(map[string]Module),
		routes:  make(map[string][]string),
	}

	for _, schema := range schemas {
		circuit.modules[schema.InLabel] = schema.Module
	}

	for _, schema := range schemas {
		inLabel := schema.InLabel
		routes := []string{}

		for _, outLabel := range schema.OutLabels {
			next := circuit.modules[outLabel]
			if next != nil {
				next.Connect(inLabel)
			}
			routes = append(routes, outLabel)
		}
		circuit.routes[inLabel] = routes
	}

	return circuit
}

func (c Circuit) Input(label string, pulse Pulse) []Signal {
	signals := []Signal{
		{"button", label, pulse},
	}

	for count := 0; count < len(signals); count++ {
		signal := signals[count]

		_, hasDest := c.modules[signal.To]
		if !hasDest {
			continue
		}

		wire := c.modules[signal.To].GetConnection(signal.From)
		out := wire.Send(signal.Pulse)

		if out == nil {
			continue
		}

		for _, next := range c.routes[signal.To] {
			signals = append(signals, Signal{
				signal.To,
				next,
				*out,
			})
		}
	}

	return signals
}
