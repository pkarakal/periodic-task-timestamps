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
