package main

import "github.com/NicholasMead/go-aoc-23/y23d14/platform"

type SpinCycleHashGenerator func(spinCycleNumber int) string

func NewSpinCycleHashGenerator(p platform.Platform) SpinCycleHashGenerator {
	memo := []string{p.Hash()}

	return func(i int) string {
		for i >= len(memo) {
			p.SpinCycle()
			memo = append(memo, p.Hash())
		}

		return memo[i]
	}
}

func DetectCycle(generator SpinCycleHashGenerator, cap int) (offset int, period int) {
	memo := map[string]int{}

	for i := 0; i < cap; i++ {
		hash := generator(i)

		if index, found := memo[hash]; found {
			return index, i - index
		} else {
			memo[hash] = i
		}
	}

	return 0, 0
}
