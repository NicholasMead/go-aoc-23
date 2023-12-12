package springs

import (
	"slices"
)

func CountValid(history History, target Record) int {
	var prev map[string]int
	current := map[string]int{"": 1}

	for h := 0; h < len(history); h++ {
		prev = current
		current = map[string]int{}

		for prev, count := range prev {
			record := RecordFromString(prev)
			var candidates []Record

			switch Result(history[h]) {
			case Ok:
				candidates = append(candidates, record.AddOk())

			case Failed:
				candidates = append(candidates, record.AddFail())

			case Unknown:
				candidates = append(candidates, record.AddOk())
				candidates = append(candidates, record.AddFail())
			}

			for _, candidate := range candidates {
				if !isPartial(target, candidate) {
					continue
				}
				current[candidate.String()] += count
			}
		}
	}

	total := 0
	for candidate, count := range current {
		record := RecordFromString(candidate).Normalise()
		if slices.Equal(record, target) {
			total += count
		}
	}
	return total
}

func isPartial(target, candidate Record) bool {
	candidate = candidate.Normalise()

	n := len(candidate)
	if n == 0 {
		return true
	}

	if len(candidate) > len(target) {
		return false
	}

	for i := 0; i < n-1; i++ {
		if candidate[i] != target[i] {
			return false
		}
	}

	return candidate[n-1] <= target[n-1]
}

func Unfold(history History, expected Record) (History, Record) {
	unfoldedHistory := history
	unFoldedExpected := expected.Copy()

	for i := 0; i < 4; i++ {
		unfoldedHistory += History(Unknown) + history
		unFoldedExpected = append(unFoldedExpected, expected...)
	}

	return unfoldedHistory, unFoldedExpected
}
