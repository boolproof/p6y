package p6y

import (
	"testing"
)

type testCase struct {
	title   string
	s       string
	err     bool
	years   int
	months  int
	days    int
	weeks   int
	hours   int
	minutes int
	seconds int
}

func TestNewDuration(t *testing.T) {
	testCases := []testCase{
		{"ERR No P", "1Y", true, 0, 0, 0, 0, 0, 0, 0},
		{"OK 1Y", "P1Y", false, 1, 0, 0, 0, 0, 0, 0},
		{"OK 1M", "P1M", false, 0, 1, 0, 0, 0, 0, 0},
		{"OK 1D", "P1D", false, 0, 0, 1, 0, 0, 0, 0},
		// {"OK 1W", "P1W", false, 0, 0, 0, 1, 0, 0, 0},
		{"OK T1H", "PT1H", false, 0, 0, 0, 0, 1, 0, 0},
		{"OK T1M", "PT1M", false, 0, 0, 0, 0, 0, 1, 0},
		{"OK T1S", "PT1S", false, 0, 0, 0, 0, 0, 0, 1},
		{"OK full", "P6Y1M2DT15H4M5S", false, 6, 1, 2, 0, 15, 4, 5},
	}

	for _, tc := range testCases {
		d, e := NewDuration(tc.s)
		if tc.err {
			if e == nil {
				t.Errorf("[%s] error should not be nil", tc.title)
			}
		} else {
			if e != nil {
				t.Errorf("[%s] error should be nil, got %s", tc.title, e.Error())
			}
		}
		if d.Years() != tc.years {
			t.Errorf("[%s] expected years: %d, got: %d", tc.title, tc.years, d.Years())
		}
		if d.Months() != tc.months {
			t.Errorf("[%s] expected months: %d, got: %d", tc.title, tc.months, d.Months())
		}
		if d.Days() != tc.days {
			t.Errorf("[%s] expected days: %d, got: %d", tc.title, tc.days, d.Days())
		}
		if d.Weeks() != tc.weeks {
			t.Errorf("[%s] expected weeks: %d, got: %d", tc.title, tc.weeks, d.Weeks())
		}
		if d.Hours() != tc.hours {
			t.Errorf("[%s] expected hours: %d, got: %d", tc.title, tc.hours, d.Hours())
		}
		if d.Minutes() != tc.minutes {
			t.Errorf("[%s] expected minutes: %d, got: %d", tc.title, tc.minutes, d.Minutes())
		}
		if d.Seconds() != tc.seconds {
			t.Errorf("[%s] expected seconds: %d, got: %d", tc.title, tc.seconds, d.Seconds())
		}
	}
}
