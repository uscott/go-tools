package tm

import (
	"strings"
	"time"
)

const oneMillion = 1000 * 1000

// Format0 returns the time argument formatted in a certain way
func Format0(t time.Time) string {
	s := strings.ReplaceAll(t.Format(time.RFC3339), "Z", "")
	return strings.ReplaceAll(s, "T", " ")
}

// Format1 returns the format returned by Format0 with ':' removed
func Format1(t time.Time) string {
	return strings.ReplaceAll(Format0(t), ":", "")
}

func GetTmStmp() float64 { // Timestamp in milliseconds
	return float64(time.Now().UnixNano()) / oneMillion
}

func GetTmStmpUtc() float64 { // UTC timestamp in milliseconds
	return float64(time.Now().UTC().UnixNano()) / oneMillion
}

// UTC returns the current time in UTC
func UTC() time.Time {
	return time.Now().UTC()
}
