package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// daysInMonth is the number of days for non-leap years in each calendar month starting at 1
var daysInMonth = [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func daysIn(m time.Month, year int) int {
	if m == time.February && isLeap(year) {
		return 29
	}
	return daysInMonth[m]
}

func parse(inp []byte) (time.Time, error) {
	format := "YYYYMMDDThhmmssZ"
	if len(inp) != len(format) {
		return time.Time{}, fmt.Errorf("length doesn't match the format length")
	}
	allowedChars := []string{"Y", "M", "D", "T", "h", "m", "s", "Z"}
	itemLen := make(map[string]int)
	for _, substring := range allowedChars {
		itemLen[substring] = strings.Count(format, substring)
	}
	var (
		Y uint64
		M uint64
		d uint64
		h uint64
		m uint64
		s uint64
	)

	// Always assume UTC by default
	var loc = time.UTC

	i := 0

	for _, char := range allowedChars {
		v := itemLen[char]
		switch char {
		case "Y":
			{
				Y, _ = strconv.ParseUint(string(inp[i:i+v]), 10, 64)
				break
			}
		case "M":
			{
				M, _ = strconv.ParseUint(string(inp[i:i+v]), 10, 64)
				break
			}
		case "D":
			{
				d, _ = strconv.ParseUint(string(inp[i:i+v]), 10, 64)
				break
			}
		case "h":
			{
				h, _ = strconv.ParseUint(string(inp[i:i+v]), 10, 64)
				break
			}
		case "m":
			{
				m, _ = strconv.ParseUint(string(inp[i:i+v]), 10, 64)
				break
			}
		case "s":
			{
				s, _ = strconv.ParseUint(string(inp[i:i+v]), 10, 64)
				break
			}
		case "T":
		case "Z":
		default:
			break
		}
		fmt.Println(i, char, v)
		i = i + v
	}
	fmt.Println(itemLen)

	switch {
	case M < 1 || M > 12: // Month 1-12
		return time.Time{}, &RangeError{
			Value:   string(inp),
			Element: "month",
			Given:   int(M),
			Min:     1,
			Max:     12,
		}
	case d < 1 || int(d) > daysIn(time.Month(M), int(Y)): // Day 1-daysIn(month, year)
		return time.Time{}, &RangeError{
			Value:   string(inp),
			Element: "day",
			Given:   int(d),
			Min:     1,
			Max:     daysIn(time.Month(M), int(Y)),
		}
	case h > 23: // Hour 0-23
		return time.Time{}, &RangeError{
			Value:   string(inp),
			Element: "hour",
			Given:   int(h),
			Min:     0,
			Max:     23,
		}
	case m > 59: // Minute 0-59
		return time.Time{}, &RangeError{
			Value:   string(inp),
			Element: "minute",
			Given:   int(m),
			Min:     0,
			Max:     59,
		}
	case s > 59: // Second 0-59
		return time.Time{}, &RangeError{
			Value:   string(inp),
			Element: "second",
			Given:   int(s),
			Min:     0,
			Max:     59,
		}
	}

	return time.Date(int(Y), time.Month(M), int(d), int(h), int(m), int(s), 0, loc), nil

}

func ParseString(s string) (time.Time, error) {
	return parse([]byte(s))
}
