package machine

import "testing"

func TestConjunction(t *testing.T) {

	t.Run("Input", func(t *testing.T) {
		var conj Module = &Conjunction{}

		conj.Connect("A")
		conj.Connect("B")

		aInput := conj.GetConnection("A")
		bInput := conj.GetConnection("B")

		out := aInput.Send(HIGH)

		if out == nil || *out != HIGH {
			t.Errorf("got %v expected %v for initial pulse", *out, HIGH)
		}

		out = bInput.Send(HIGH)

		if out == nil || *out != LOW {
			t.Errorf("got %v expected %v for 2nd pulse", *out, LOW)
		}

		out = bInput.Send(LOW)

		if out == nil || *out != HIGH {
			t.Errorf("got %v expected %v for 3rd pulse", *out, HIGH)
		}
	})
}
