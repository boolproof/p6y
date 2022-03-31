package p6y

import "errors"

type Duration struct {
	years   int
	months  int
	days    int
	weeks   int
	hours   int
	minutes int
	seconds int
}

func (d Duration) Years() int {
	return d.years
}

func (d Duration) Months() int {
	return d.months
}
func (d Duration) Days() int {
	return d.days
}
func (d Duration) Weeks() int {
	return d.weeks
}
func (d Duration) Hours() int {
	return d.hours
}
func (d Duration) Minutes() int {
	return d.minutes
}

func (d Duration) Seconds() int {
	return d.seconds
}

func NewDuration(s string) (Duration, error) {
	e := errors.New("failed to parse input string")
	var d Duration

	if len(s) < 0 || s[0] != 'P' {
		return d, e
	}

	return d, nil
}
