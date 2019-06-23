package utils

import "time"

// ElapsedTime returns duration
func ElapsedTime(fn func()) time.Duration {
	start := time.Now()
	fn()
	t := time.Now()
	return t.Sub(start)
}
