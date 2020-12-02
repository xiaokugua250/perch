package tools

import "time"

type TimeSlice []time.Time

func (timeS TimeSlice) Len() int {
	return len(timeS)
}

// Define compare
func (timeS TimeSlice) Less(i, j int) bool {
	return timeS[i].Before(timeS[j])
}

// Define swap over an array
func (timeS TimeSlice) Swap(i, j int) {
	timeS[i], timeS[j] = timeS[j], timeS[i]
}
