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
	var d, tmp duration

	if len(s) < 2 || s[0] != 'P' {
		return d, e
	}

	var rest string
	var err error
	d.weeks, rest, err = extrct(s[1:], "W")
	if err != nil || (err == nil && len(rest) > 0 && rest != s[1:]) {
		return tmp, e
	} else if rest != s[1:] {
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
		return tmp, e
	}

	if dc != "" {
		rest = dc
		var err error
		d.years, rest, err = extrct(rest, "Y")
		if err != nil {
			return tmp, err
		}

		d.months, rest, err = extrct(rest, "M")
		if err != nil {
			return tmp, err
		}

		d.days, rest, err = extrct(rest, "D")
		if err != nil {
			return tmp, err
		}
	}

	if tc != "" {
		rest = tc
		var err error
		d.hours, rest, err = extrct(rest, "H")
		if err != nil {
			return tmp, err
		}

		d.minutes, rest, err = extrct(rest, "M")
		if err != nil {
			return tmp, err
		}

		d.seconds, rest, err = extrct(rest, "S")
		if err != nil {
			return tmp, err
		}
	}

	if len(rest) > 0 {
		return tmp, e
	}

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
