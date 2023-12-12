package springs

import "fmt"

type Counter struct {
	History History
	Record  Record
	memo    map[string]int
}

func (counter Counter) Count() int {
	if counter.memo == nil {
		counter.memo = make(map[string]int)
	}

	history := counter.History.Normalise()
	record := counter.Record.Normalise()

	for r := 0; r <= len(record); r++ {
		for h := 0; h <= len(history); h++ {
			history := counter.History[:h]
			target := counter.Record[:r].DrawHistory()
			record := counter.Record[:r]

			count := counter.computeCount(history, record)

			fmt.Println(history, " ", target, " ", count, " ")
		}
	}

	return counter.memo[counter.key()]
}

func (counter Counter) computeCount(history History, record Record) (c int) {
	history = history.Normalise()
	h := len(history)
	r := len(record)

	// check memo
	k := key(history, record)
	if count, found := counter.memo[k]; found {
		return count
	} else {
		//set memo on return
		defer func() {
			counter.memo[k] = c
		}()
	}

	if h == 0 && r == 0 {
		return 1
	} else if h == 0 || r == 0 {
		return 0
	}

	okCount := func() int {
		if r == 0 {
			return 1
		}
		if record[r-1] > 0 {
			return 0
		}
		return counter.computeCount(history[:h-1], record[:r-1])
	}

	failCount := func() int {
		if r == 0 {
			return 0
		}
		if record[r-1] == 0 {
			return 0
		}
		next := append(Record{}, record...)
		next[r-1] -= 1
		next = next.Normalise()
		return counter.computeCount(history[:h-1], next)
	}

	switch Result(history[h-1]) {
	case Ok:
		return okCount()

	case Failed:
		return failCount()

	case Unknown:
		var (
			ok   = okCount()
			fail = failCount()
		)
		return ok + fail

	default:
		panic(history)
	}
}

func (c Counter) key() string {
	return key(c.History, c.Record)
}

func key(h History, r Record) string {
	return fmt.Sprint(h, " ", r)
}
