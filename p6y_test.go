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
		{"ERR trash in before", "AP6Y1M2DT15H4M5S", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR trash in after", "P6Y1M2DT15H4M5SA", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR non int Y", "P.Y", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR non int M", "P M", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR non int D", "PaD", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR non int W", "PZW", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR non int TH", "PT-H", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR non int TM", "PT'M", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR non int TS", "PT\"S", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR negative Y", "P-6Y", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR negative M", "P-001M", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR negative D", "P-2D", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR negative W", "P-8W", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR negative TH", "PT-15H", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR negative TM", "PT-04M", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR negative TS", "PT-5S", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR combine W with Y", "P1Y1W", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR combine W with M", "P1M1W", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR combine W with D", "P1W1D", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR combine W with TH", "P1WT1H", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR combine W with TM", "P1WT1M", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR combine W with TS", "P1WT1S", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR double Y", "P2Y1Y", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR double M", "P1M2M", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR double D", "P0D2D", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR double W", "P2W0W", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR double TH", "PT1H0H", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR double TM", "PT0M2M", true, 0, 0, 0, 0, 0, 0, 0},
		{"ERR double TS", "PT60S2S", true, 0, 0, 0, 0, 0, 0, 0},

		{"OK 1Y", "P1Y", false, 1, 0, 0, 0, 0, 0, 0},
		{"OK 1M", "P1M", false, 0, 1, 0, 0, 0, 0, 0},
		{"OK 1D", "P1D", false, 0, 0, 1, 0, 0, 0, 0},
		{"OK 1W", "P1W", false, 0, 0, 0, 1, 0, 0, 0},
		{"OK T1H", "PT1H", false, 0, 0, 0, 0, 1, 0, 0},
		{"OK T1M", "PT1M", false, 0, 0, 0, 0, 0, 1, 0},
		{"OK T1S", "PT1S", false, 0, 0, 0, 0, 0, 0, 1},
		{"OK 1Y", "P999Y", false, 999, 0, 0, 0, 0, 0, 0},
		{"OK 1M", "P999M", false, 0, 999, 0, 0, 0, 0, 0},
		{"OK 1D", "P999D", false, 0, 0, 999, 0, 0, 0, 0},
		{"OK 1W", "P999W", false, 0, 0, 0, 999, 0, 0, 0},
		{"OK T1H", "PT999H", false, 0, 0, 0, 0, 999, 0, 0},
		{"OK T1M", "PT999M", false, 0, 0, 0, 0, 0, 999, 0},
		{"OK T1S", "PT999S", false, 0, 0, 0, 0, 0, 0, 999},
		{"OK date only", "P6Y1M2D", false, 6, 1, 2, 0, 0, 0, 0},
		{"OK time only", "PT15H4M5S", false, 0, 0, 0, 0, 15, 4, 5},
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
