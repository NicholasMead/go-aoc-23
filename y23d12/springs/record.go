package springs

import (
	"fmt"
	"strconv"
	"strings"
)

type Record []int

func RecordFromString(s string) Record {
	parts := strings.Split(s, ",")
	record := make(Record, len(parts))

	for i, r := range parts {
		record[i], _ = strconv.Atoi(r)
	}

	return record
}

func (record Record) AddOk() Record {
	n := len(record) - 1
	copy := record.Copy()

	if record[n] != 0 {
		return append(copy, 0)
	} else {
		return copy
	}
}

func (record Record) AddFail() Record {
	n := len(record) - 1
	copy := record.Copy()
	copy[n]++
	return copy
}

func (record Record) String() string {
	s := []string{}
	for _, r := range record {
		s = append(s, fmt.Sprint(r))
	}
	return strings.Join(s, ",")
}

func (record Record) Copy() Record {
	return append(Record{}, record...)
}

func (record Record) Normalise() Record {
	n := len(record)
	if record[n-1] == 0 {
		return record[:n-1]
	}
	return record
}
