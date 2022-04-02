package p6y

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type duration struct {
	years   int
	months  int
	days    int
	weeks   int
	hours   int
	minutes int
	seconds int
}

func (d duration) Years() int {
	return d.years
}

func (d duration) Months() int {
	return d.months
}
func (d duration) Days() int {
	return d.days
}
func (d duration) Weeks() int {
	return d.weeks
}
func (d duration) Hours() int {
	return d.hours
}
func (d duration) Minutes() int {
	return d.minutes
}

func (d duration) Seconds() int {
	return d.seconds
}

func NewDuration(s string) (duration, error) {
	e := errors.New("failed to parse input string")
	var d duration

	if len(s) < 2 || s[0] != 'P' {
		return d, e
	}

	var rest string

	weeks, rest, err := extrct(s[1:], "W")
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
		years, rest, err = extrct(rest, "Y")
		if err != nil {
			return d, err
		}

		months, rest, err = extrct(rest, "M")
		if err != nil {
			return d, err
		}

		days, rest, err = extrct(rest, "D")
		if err != nil {
			return d, err
		}
	}

	if tc != "" {
		rest = tc
		var err error
		hours, rest, err = extrct(rest, "H")
		if err != nil {
			return d, err
		}

		minutes, rest, err = extrct(rest, "M")
		if err != nil {
			return d, err
		}

		seconds, rest, err = extrct(rest, "S")
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

func extrct(s, t string) (int, string, error) {
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
