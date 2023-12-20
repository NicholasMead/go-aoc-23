package machine

import "testing"

func TestCircuit(t *testing.T) {
	t.Run("Simple Circuit", func(t *testing.T) {
		circuit := NewCircuit([]Schema{
			{"broadcaster", Broadcast{}, []string{"a", "b", "c"}},
			{"a", &FlipFlop{}, []string{"b"}},
			{"b", &FlipFlop{}, []string{"c"}},
			{"c", &FlipFlop{}, []string{"inv"}},
			{"inv", &Conjunction{}, []string{"a"}},
		})

		signals := circuit.Input("broadcaster", LOW)

		if len(signals) != 12 {
			t.Errorf("expected %v signals, got %v: %v", 12, len(signals), signals)
		}
	})

	t.Run("Circuit With Output", func(t *testing.T) {
		circuit := NewCircuit([]Schema{
			{"broadcaster", Broadcast{}, []string{"a"}},
			{"a", &FlipFlop{}, []string{"inv", "con"}},
			{"b", &FlipFlop{}, []string{"con"}},
			{"inv", &Conjunction{}, []string{"b"}},
			{"con", &Conjunction{}, []string{"output"}},
		})

		signals := circuit.Input("broadcaster", LOW)

		if len(signals) != 8 {
			t.Errorf("expected %v signals, got %v: %v", 12, len(signals), signals)
		}
	})
}
