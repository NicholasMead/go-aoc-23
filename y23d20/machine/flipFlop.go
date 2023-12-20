package machine

import "fmt"

type FlipFlop struct {
	on bool
}

func (f *FlipFlop) Connect(string) {}

func (f *FlipFlop) GetConnection(string) Wire {
	return flipFlopInput{f}
}

func (FlipFlop) Type() string {
	return "FlipFlop"
}

type flipFlopInput struct {
	flipFlop *FlipFlop
}

func (input flipFlopInput) Send(pulse Pulse) *Pulse {
	switch pulse {
	case HIGH:
		return nil

	case LOW:
		input.flipFlop.on = !input.flipFlop.on

		if input.flipFlop.on {
			return &HIGH
		} else {
			return &LOW
		}

	default:
		err := fmt.Errorf("unknown Pulse: %v", pulse)
		panic(err)
	}
}
