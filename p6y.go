package p6y

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

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

	if len(s) < 2 || s[0] != 'P' {
		return d, e
	}

	var rest string

	weeks, rest, err := x(s[1:], "W")
	if err != nil || (err == nil && len(rest) > 0 && rest != s[1:]) {
		return d, e
	} else if rest != s[1:] {
		d.weeks = weeks
		return d, nil
	}

	cs := strings.Split(s, "T")
	var dc, tc string
	if len(cs) == 2 {
		dc, tc = cs[0][1:], cs[1]
	} else if len(cs) == 1 {
		if cs[0][0] == 'P' {
			dc = cs[0][1:]
		} else {
			tc = cs[0]
		}
	} else {
		return d, e
	}

	var years, months, days, hours, minutes, seconds int

	if dc != "" {
		rest = dc
		var err error
		years, rest, err = x(rest, "Y")
		if err != nil {
			return d, err
		}

		months, rest, err = x(rest, "M")
		if err != nil {
			return d, err
		}

		days, rest, err = x(rest, "D")
		if err != nil {
			return d, err
		}
	}

	if tc != "" {
		rest = tc
		var err error
		hours, rest, err = x(rest, "H")
		if err != nil {
			return d, err
		}

		minutes, rest, err = x(rest, "M")
		if err != nil {
			return d, err
		}

		seconds, rest, err = x(rest, "S")
		if err != nil {
			return d, err
		}
	}

	if len(rest) > 0 {
		return d, e
	}

	d.years = years
	d.months = months
	d.days = days
	d.weeks = weeks
	d.hours = hours
	d.minutes = minutes
	d.seconds = seconds

	return d, nil
}

func x(s, t string) (int, string, error) {
	var tval int

	tpos := strings.Index(s, t)
	if tpos == 0 {
		return 0, "", errors.New(fmt.Sprintf("'%s' token without value", t))
	} else if tpos > 0 {
		tpart := s[0:tpos]
		var err error
		tval, err = strconv.Atoi(tpart)
		if err != nil || tval < 0 {
			return 0, "", errors.New(fmt.Sprintf("negative '%s' token value", t))
		}

		s = s[tpos+1:]
	}

	return tval, s, nil
}
