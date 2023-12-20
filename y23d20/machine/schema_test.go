package machine

import (
	"slices"
	"testing"
)

func TestSchema(t *testing.T) {
	t.Run("broadcaster", func(t *testing.T) {
		line := "broadcaster -> a, b, c"
		expect := Schema{"broadcaster", Broadcast{}, []string{"a", "b", "c"}}

		schema := ParseSchema(line)

		if !isMatch(schema, expect) {
			t.Errorf("got %v, expected %v", schema, expect)
		}
	})

	t.Run("flipFlop", func(t *testing.T) {
		line := "%a -> a, b, c"
		expect := Schema{"a", &FlipFlop{}, []string{"a", "b", "c"}}

		schema := ParseSchema(line)

		if !isMatch(schema, expect) {
			t.Errorf("got %v, expected %v", schema, expect)
		}
	})

	t.Run("Conjunction", func(t *testing.T) {
		line := "&a -> a, b, c"
		expect := Schema{"a", &Conjunction{}, []string{"a", "b", "c"}}

		schema := ParseSchema(line)

		if !isMatch(schema, expect) {
			t.Errorf("got %v, expected %v", schema, expect)
		}
	})
}

func isMatch(a, b Schema) bool {
	if a.InLabel != b.InLabel {
		return false
	}

	if a.Module.Type() != b.Module.Type() {
		return false
	}

	return slices.Equal(a.OutLabels, b.OutLabels)
}
