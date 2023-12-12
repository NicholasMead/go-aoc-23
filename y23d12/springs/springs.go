package springs

import (
	"slices"
	"strings"
)

type historyFinder struct {
	target Record
	memo   map[History][]Record
}

func NewHistoryFinder(target Record) *historyFinder {
	return &historyFinder{
		target: target,
		memo: map[History][]Record{
			"": {{0}},
		},
	}
}

func (finder *historyFinder) CountValid(template History) int {
	if records, memo := finder.memo[template]; memo {
		return len(records)
	}

	for h := 0; h < len(template); h++ {
		prefixRecords := finder.memo[template[0:h]]
		records := []Record{}

		for _, prefix := range prefixRecords {
			record := append(Record{}, prefix...)
			n := len(record) - 1
			if n == -1 {
				panic("invalid prefix")
			}

			switch Result(template[h]) {
			case Ok:
				if record[n] != 0 {
					record = append(record, 0)
				}
				records = append(records, record)

			case Failed:
				record[n]++
				records = append(records, record)

			case Unknown:
				okRecord := append(Record{}, record...)
				if okRecord[n] != 0 {
					okRecord = append(okRecord, 0)
				}

				failRecord := append(Record{}, record...)
				failRecord[n]++

				records = append(records, okRecord, failRecord)
			}
		}

		records = slices.DeleteFunc(records, func(r Record) bool {
			return !isPartial(finder.target, r)
		})

		finder.memo[template[0:h+1]] = records
	}

	count := 0
	for _, candidate := range finder.memo[template] {
		if candidate[len(candidate)-1] == 0 {
			candidate = candidate[:len(candidate)-1]
		}

		if slices.Equal(candidate, finder.target) {
			count++
		}
	}
	return count
}

func isPartial(target, candidate Record) bool {
	if candidate[len(candidate)-1] == 0 {
		candidate = candidate[:len(candidate)-1]
	}

	if len(candidate) > len(target) {
		return false
	}

	for i := range candidate {
		if candidate[i] > target[i] {
			return false
		}
	}
	return true
}

func CountValid(history History, expected Record) (count int) {
	queue := []History{history}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if strings.ContainsRune(string(next), rune(Unknown)) {
			var (
				ok     = History(strings.Replace(string(next), string(Unknown), string(Ok), 1))
				failed = History(strings.Replace(string(next), string(Unknown), string(Failed), 1))
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
		unfoldedHistory += History(Unknown) + history
		unFoldedExpected = append(unFoldedExpected, expected...)
	}

	return unfoldedHistory, unFoldedExpected
}
