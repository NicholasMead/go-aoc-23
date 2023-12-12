package springs

import (
	"slices"
	"strings"
)

type History = string

type Result = rune

var (
	Ok      Result = '.'
	Failed  Result = '#'
	Unknown Result = '?'
)

type Record []int

func CountValid(history History, expected Record) (count int) {
	queue := []History{history}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if strings.ContainsRune(next, Unknown) {
			var (
				ok     = History(strings.Replace(next, string(Unknown), string(Ok), 1))
				failed = History(strings.Replace(next, string(Unknown), string(Failed), 1))
			)
			queue = append(queue, ok, failed)
			continue
		}

		if IsValid(next, expected) {
			count++
		}
	}

	return count
}

func IsValid(history History, expected Record) bool {
	actual := Record{0}
	a := 0

	for _, result := range history {
		switch Result(result) {
		case Ok:
			if actual[a] != 0 {
				actual = append(actual, 0)
				a++
			}

		case Failed:
			actual[a]++

		default:
			panic(string(result))
		}
	}

	if actual[a] == 0 {
		actual = actual[0:a]
	}

	return slices.Equal(actual, expected)
}

func Unfold(history History, expected Record) (History, Record) {
	unfoldedHistory := history
	unFoldedExpected := append(Record{}, expected...)

	for i := 0; i < 5; i++ {
		unfoldedHistory += History(string(Unknown) + history)
		unFoldedExpected = append(unFoldedExpected, expected...)
	}

	return unfoldedHistory, unFoldedExpected
}
