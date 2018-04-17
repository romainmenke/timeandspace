package timeandspace

import "time"

func (t TimeAndSpace) Time() time.Time {
	if t.TimeX0 == 0 {
		return time.Time{}
	}

	return time.Date(int(t.TimeX0), 0, 0, 0, 0, 0, 0, time.UTC).AddDate(0, 0, int(t.TimeY3))
}

func NewTime(t time.Time) TimeAndSpace {
	_, week := t.ISOWeek()
	return TimeAndSpace{
		TimeX0: uint16(t.Year()),
		TimeY0: quarterForMonth(t.Month()),
		TimeY1: uint16(t.Month()),
		TimeY2: uint16(week),
		TimeY3: uint16(t.YearDay()),
	}
}

func firstMonthForQuarter(quarter uint16) time.Month {
	switch quarter {
	case 1:
		return time.January
	case 2:
		return time.April
	case 3:
		return time.July
	case 4:
		return time.October
	default:
		return time.January
	}
}

func quarterForMonth(month time.Month) uint16 {
	switch month {
	case 1, 2, 3:
		return 1
	case 4, 5, 6:
		return 2
	case 7, 8, 9:
		return 3
	case 10, 11, 12:
		return 4
	default:
		return 1
	}
}
