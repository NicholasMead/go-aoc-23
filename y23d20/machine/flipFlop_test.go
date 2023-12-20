package machine

import "testing"

func TestFlipFlop(t *testing.T) {
	t.Run("On High Do Nothing", func(t *testing.T) {
		var flipFlop Module = &FlipFlop{}

		out := flipFlop.GetConnection("A").Send(HIGH)

		if out != nil {
			t.Errorf("Got %v expected %v", out, nil)
		}
	})

	t.Run("On Low Send Alternate Signal", func(t *testing.T) {
		var flipFlop Module = &FlipFlop{}
		input := flipFlop.GetConnection("A")

		out := input.Send(LOW)

		if out == nil || *out != HIGH {
			t.Errorf("Got %v expected %v", *out, HIGH)
		}

		out = input.Send(LOW)

		if out == nil || *out != LOW {
			t.Errorf("Got %v expected %v", *out, LOW)
		}
	})
}
