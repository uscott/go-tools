package tm

import (
	"math"
	"strings"
	"time"
)

const oneMillion = 1000 * 1000

// Format0 returns the time argument formatted first in RFC3339
// then with "Z" removed, then "T" replaced with white space
func Format0(t time.Time) string {
	s := strings.ReplaceAll(t.Format(time.RFC3339), "Z", "")
	return strings.ReplaceAll(s, "T", " ")
}

// Format1 returns the format returned by Format0 with ':' removed
func Format1(t time.Time) string {
	return strings.ReplaceAll(Format0(t), ":", "")
}

// Format2 returns the date as YYYY-MM-DD
func Format2(t time.Time) string {
	return t.Format(time.RFC3339)[:10]
}

func GetTmStmp() float64 { // Timestamp in milliseconds
	return float64(time.Now().UnixNano()) / oneMillion
}

func GetTmStmpUtc() float64 { // UTC timestamp in milliseconds
	return float64(time.Now().UTC().UnixNano()) / oneMillion
}

func Milliseconds(t time.Time) int {
	n := t.Nanosecond()
	return int(math.Round(float64(n) / 1e6))
}

func Microseconds(t time.Time) int {
	n := t.Nanosecond()
	return int(math.Round(float64(n) / 1000))
}
func OnTheSecond(t time.Time) bool {
	return Milliseconds(t) == 0
}

func OnTheMinute(t time.Time) bool {
	return t.Second() == 0
}

func OnTheHour(t time.Time) bool {
	return t.Second() == 0 && t.Minute() == 0
}

// UTC returns the current time in UTC
func UTC() time.Time {
	return time.Now().UTC()
}
