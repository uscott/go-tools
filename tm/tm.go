package tm

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const oneMillion = 1000 * 1000

var TimeZero = time.Date(1970, 1, 1, 0, 0, 0, 0, time.FixedZone("utc", 0))

// Format0 returns the time argument formatted first in RFC3339
// then with "Z" removed, then "T" replaced with white space
// So YYYY-MM-DDThh:mm:ddZ => YYYY-MM-DD hh:mm:dd
func Format0(t time.Time) string {
	s := strings.ReplaceAll(t.Format(time.RFC3339), "Z", "")
	return strings.ReplaceAll(s, "T", " ")
}

// Format1 returns the format returned by Format0 with ':' removed
// YYYY-MM-DDThh:mm:ddZ => YYYY-MM-DD hhmmdd
func Format1(t time.Time) string {
	return strings.ReplaceAll(Format0(t), ":", "")
}

// Format2 returns the date as YYYY-MM-DD
// YYYY-MM-DDThh:mm:ddZ => YYYY-MM-DD
func Format2(t time.Time) string {
	return t.Format(time.RFC3339)[:10]
}

func GetTimeStampMS() float64 { // Timestamp in milliseconds
	return float64(time.Now().UnixNano()) / oneMillion
}

func GetTimeStampMSUTC() float64 { // UTC timestamp in milliseconds
	return float64(time.Now().UTC().UnixNano()) / oneMillion
}

func Milliseconds(t time.Time) int {
	n := t.Nanosecond()
	return int(math.Round(float64(n)/oneMillion)) % 1000
}

func Microseconds(t time.Time) int {
	n := t.Nanosecond()
	return int(math.Round(float64(n)/1000)) % oneMillion
}

func ParseYMD(s string) (time.Time, error) {

	zero := time.Time{}

	if len(s) < 8 {
		return zero, fmt.Errorf("String too short: %s", s)
	}

	b := []byte(s)
	if b[7] == '-' || b[7] == ' ' {
		b = append(b[:7], b[8:]...)
	}
	if b[4] == '-' || b[4] == ' ' {
		b = append(b[:4], b[5:]...)
	}

	y, err := strconv.Atoi(string(b[:4]))
	if err != nil {
		return zero, err
	}

	m, err := strconv.Atoi(string(b[4:6]))
	if err != nil {
		return zero, err
	}

	d, err := strconv.Atoi(string(b[6:8]))
	if err != nil {
		return zero, err
	}

	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC), nil
}

func ParseYMDHMS(s string) (time.Time, error) {

	zero := time.Time{}

	if len(s) < 14 {
		return zero, fmt.Errorf("String too short: %s", s)
	}

	b := []byte(s)
	for _, i := range []int{16, 13, 10, 7, 4} {
		if len(b) > i+1 && (b[i] == ':' || b[i] == ' ' || b[i] == '-') {
			b = append(b[:i], b[i+1:]...)
		}
	}

	y, err := strconv.Atoi(string(b[:4]))
	if err != nil {
		return zero, err
	}

	m, err := strconv.Atoi(string(b[4:6]))
	if err != nil {
		return zero, err
	}

	d, err := strconv.Atoi(string(b[6:8]))
	if err != nil {
		return zero, err
	}

	hh, err := strconv.Atoi(string(b[8:10]))
	if err != nil {
		return zero, err
	}

	mm, err := strconv.Atoi(string(b[10:12]))
	if err != nil {
		return zero, err
	}

	dd, err := strconv.Atoi(string(b[12:14]))
	if err != nil {
		return zero, err
	}

	return time.Date(y, time.Month(m), d, hh, mm, dd, 0, time.UTC), nil
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
