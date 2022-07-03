package internal

import (
	"testing"
	"time"
)

var givenDate = time.Date(2020, 12, 13, 14, 15, 16, 0, time.UTC)

func TestHourlyNext(t *testing.T) {
	task := PeriodicTask{
		TaskPeriod:      Hourly,
		InvocationPoint: givenDate,
		Timezone:        *time.UTC,
	}
	next := task.Next()
	expected := time.Date(2020, 12, 13, 15, 15, 16, 0, time.UTC)
	if next.InvocationPoint != expected {
		t.Errorf("Expected %v, Got: %v", expected, next.InvocationPoint)
	}
}

func TestDailyNext(t *testing.T) {
	task := PeriodicTask{
		TaskPeriod:      Daily,
		InvocationPoint: givenDate,
		Timezone:        *time.UTC,
	}
	next := task.Next()
	expected := time.Date(2020, 12, 14, 14, 15, 16, 0, time.UTC)
	if next.InvocationPoint != expected {
		t.Errorf("Expected %v, Got: %v", expected, next.InvocationPoint)
	}
}

func TestMontlyNext(t *testing.T) {
	task := PeriodicTask{
		TaskPeriod:      Monthly,
		InvocationPoint: givenDate,
		Timezone:        *time.UTC,
	}
	next := task.Next()
	expected := time.Date(2021, 01, 13, 14, 15, 16, 0, time.UTC)
	if next.InvocationPoint != expected {
		t.Errorf("Expected %v, Got: %v", expected, next.InvocationPoint)
	}
}

func TestYearlyNext(t *testing.T) {
	task := PeriodicTask{
		TaskPeriod:      Yearly,
		InvocationPoint: givenDate,
		Timezone:        *time.UTC,
	}
	next := task.Next()
	expected := time.Date(2021, 12, 13, 14, 15, 16, 0, time.UTC)
	if next.InvocationPoint != expected {
		t.Errorf("Expected %v, Got: %v", expected, next.InvocationPoint)
	}
}
