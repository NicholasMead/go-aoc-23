package springs

import "strings"

type Result rune

var (
	Ok      Result = '.'
	Failed  Result = '#'
	Unknown Result = '?'
)

type History string

func (h History) Normalise() History {
	history := string(h)

	history = strings.Trim(history, ".")

	for strings.Contains(history, "..") {
		history = strings.ReplaceAll(history, "..", ".")
	}

	return History(history)
}

func (h *History) Add(r Result) {
	*h = *h + History(r)
}

func (h History) String() string {
	return string(h.Normalise())
}
