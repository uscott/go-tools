package tm

import (
	"time"
)

const one_million = 1000 * 1000

func GetTmStmp() float64 { // Timestamp in milliseconds
	return float64(time.Now().UnixNano()) / one_million
}

func GetTmStmpUtc() float64 { // UTC timestamp in milliseconds
	return float64(time.Now().UTC().UnixNano()) / one_million
}

// UTC returns the current time in UTC
func UTC() time.Time {
	return time.Now().UTC()
}
