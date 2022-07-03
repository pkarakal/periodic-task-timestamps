package internal

import (
	"testing"
	"time"
)

func TestParsingDate(t *testing.T) {
	given := "20060102T030400Z"
	expected := time.Date(2006, 01, 02, 03, 04, 00, 0, time.UTC)
	got, e := ParseString(given)
	if e != nil {
		t.Fatalf("Got error instead of value, %v", e)
	}
	if got != expected {
		t.Fatalf("Dates don't match. Expected: %v, Got: %v", expected, got)
	}
}

func TestShouldReturnError(t *testing.T) {
	given := "20201310T030303Z"
	expected := RangeError{
		Value:   given,
		Element: "month",
		Given:   13,
		Min:     1,
		Max:     12,
	}
	_, got := ParseString(given)
	if got != nil && got.Error() != expected.Error() {
		t.Fatalf("Expected month range error. Expected : %v, Got: %v", expected.Error(), got.Error())
	}
}

func TestShouldConvertDateToStringWithZeroPad(t *testing.T) {
	given := time.Date(2020, 01, 31, 03, 03, 03, 0, time.UTC)
	expected := "20200131T030303Z"
	got := DateToString(given)
	if expected != got {
		t.Fatalf("Expected %s. Got %s", expected, got)
	}
}
func TestShouldConvertDateToStringWithoutZeroPad(t *testing.T) {
	given := time.Date(2020, 10, 31, 13, 14, 15, 0, time.UTC)
	expected := "20201031T131415Z"
	got := DateToString(given)
	if expected != got {
		t.Fatalf("Expected %s. Got %s", expected, got)
	}
}

func TestShouldGetTimeDifference(t *testing.T) {
	givenA := time.Date(2020, 10, 31, 13, 14, 15, 0, time.UTC)
	givenB := time.Date(2021, 11, 12, 15, 16, 17, 0, time.UTC)
	year, month, day, hour, min, sec := Diff(givenA, givenB)
	if year != 1 && month != 0 && day != 12 && hour != 2 && min != 2 && sec != 2 {
		t.Fatalf("Time difference not correct")
	}
}

func TestShouldGetTimeDifferenceLeapYear(t *testing.T) {
	givenA := time.Date(2020, 02, 28, 13, 14, 15, 0, time.UTC)
	givenB := time.Date(2021, 11, 12, 15, 16, 17, 0, time.UTC)
	year, month, day, hour, min, sec := Diff(givenA, givenB)
	if year != 1 && month != 8 && day != 13 && hour != 2 && min != 2 && sec != 2 {
		t.Fatalf("Time difference not correct")
	}
}

func TestShouldGetTimeDifferenceNonLeapYear(t *testing.T) {
	givenA := time.Date(2021, 02, 28, 13, 14, 15, 0, time.UTC)
	givenB := time.Date(2022, 11, 12, 15, 16, 17, 0, time.UTC)
	year, month, day, hour, min, sec := Diff(givenA, givenB)
	if year != 1 && month != 8 && day != 12 && hour != 2 && min != 2 && sec != 2 {
		t.Fatalf("Time difference not correct")
	}
}
