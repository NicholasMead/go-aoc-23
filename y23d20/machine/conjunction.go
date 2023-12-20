package machine

type Conjunction struct {
	inputs map[string]Pulse
}

func (conjunction *Conjunction) Connect(input string) {
	if conjunction.inputs == nil {
		conjunction.inputs = make(map[string]Pulse)
	}

	conjunction.inputs[input] = false
}

func (conjunction *Conjunction) GetConnection(input string) Wire {
	return conjunctionInput{
		input,
		conjunction,
	}
}

func (Conjunction) Type() string {
	return "Conjunction"
}

type conjunctionInput struct {
	name     string
	junction *Conjunction
}

func (i conjunctionInput) Send(pulse Pulse) *Pulse {
	if i.junction == nil {
		return nil
	}

	i.junction.inputs[i.name] = pulse

	for _, memo := range i.junction.inputs {
		if !memo {
			return &HIGH
		}
	}
	return &LOW
}
