package common

import "time"

func Timer(fn func()) time.Duration {
	start := time.Now()
	fn()
	end := time.Now()
	return end.Sub(start)
}
