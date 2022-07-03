package internal

import "time"

const (
	Hourly = iota
	Daily
	Monthly
	Yearly
)

type PeriodicTask struct {
	TaskPeriod      int
	InvocationPoint time.Time
	Timezone        time.Location
}

func (pt *PeriodicTask) Next() *PeriodicTask {
	switch pt.TaskPeriod {
	case Hourly:
		return &PeriodicTask{
			TaskPeriod:      pt.TaskPeriod,
			InvocationPoint: pt.InvocationPoint.Add(1 * time.Hour),
			Timezone:        pt.Timezone,
		}
	case Daily:
		return &PeriodicTask{
			TaskPeriod:      pt.TaskPeriod,
			InvocationPoint: time.Date(pt.InvocationPoint.Year(), pt.InvocationPoint.Month(), pt.InvocationPoint.Day()+1, pt.InvocationPoint.Hour(), pt.InvocationPoint.Minute(), pt.InvocationPoint.Second(), pt.InvocationPoint.Nanosecond(), pt.InvocationPoint.Location()),
			Timezone:        pt.Timezone,
		}
	case Monthly:
		return &PeriodicTask{
			TaskPeriod:      pt.TaskPeriod,
			InvocationPoint: time.Date(pt.InvocationPoint.Year(), pt.InvocationPoint.Month()+1, pt.InvocationPoint.Day(), pt.InvocationPoint.Hour(), pt.InvocationPoint.Minute(), pt.InvocationPoint.Second(), pt.InvocationPoint.Nanosecond(), pt.InvocationPoint.Location()),
			Timezone:        pt.Timezone,
		}
	case Yearly:
		return &PeriodicTask{
			TaskPeriod:      pt.TaskPeriod,
			InvocationPoint: time.Date(pt.InvocationPoint.Year()+1, pt.InvocationPoint.Month(), pt.InvocationPoint.Day(), pt.InvocationPoint.Hour(), pt.InvocationPoint.Minute(), pt.InvocationPoint.Second(), pt.InvocationPoint.Nanosecond(), pt.InvocationPoint.Location()),
			Timezone:        pt.Timezone,
		}
	default:
		return pt
	}
}
