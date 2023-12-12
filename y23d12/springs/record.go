package springs

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Record []int

func NewRecordFromString(s string) Record {
	parts := strings.Split(s, ",")
	record := make(Record, len(parts))

	for i, r := range parts {
		record[i], _ = strconv.Atoi(r)
	}

	return record
}

func (record Record) DrawHistory() History {
	history := History("")

	for _, count := range record {
		for i := 0; i < count; i++ {
			history.Add(Failed)
		}
		history.Add(Ok)
	}

	return history.Normalise()
}

func (record Record) String() string {
	s := []string{}
	for _, r := range record {
		s = append(s, fmt.Sprint(r))
	}
	return strings.Join(s, ",")
}

func (r Record) Normalise() Record {
	copy := append(Record{}, r...)
	slices.DeleteFunc(copy, func(i int) bool { return i == 0 })
	return copy
}
