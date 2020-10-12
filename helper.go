package main

import "time"

// Cast a boolean value to a float64 value containing 0 or 1
func boolToFloat64(b bool) float64 {
	if b == true {
		return float64(1)
	}

	return float64(0)
}

// Take an integer and a duration and multiply them
// Example: multiplyDuration(3, time.Seconds)
// to get a duration of 3 seconds
func multiplyDuration(i int, t time.Duration) time.Duration {
	return time.Duration(i * int(t))
}
