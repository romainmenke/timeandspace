package timeandspace

import "time"

func (t TimeAndSpace) Time() time.Time {
	if t.TX == 0 {
		return time.Time{}
	}

	switch t.TZ {
	case 0:
		return time.Date(int(t.TX), 0, 0, 0, 0, 0, 0, time.UTC)
	case 1:
		return time.Date(int(t.TX), firstMonthForQuarter(t.TY), 0, 0, 0, 0, 0, time.UTC)
	case 2:
		return time.Date(int(t.TX), time.Month(t.TY), 0, 0, 0, 0, 0, time.UTC)
	case 3:
		return time.Date(int(t.TX), 0, 0, 0, 0, 0, 0, time.UTC).AddDate(0, 0, int(t.TY)*7)
	case 4:
		return time.Date(int(t.TX), 0, 0, 0, 0, 0, 0, time.UTC).AddDate(0, 0, int(t.TY))
	default:
		return time.Date(int(t.TX), 0, 0, 0, 0, 0, 0, time.UTC)
	}

}

func NewTimes(t time.Time) []TimeAndSpace {
	t = t.UTC()

	times := make([]TimeAndSpace, 0, 5)

	var z uint8
	for z = 0; z < uint8(cap(times)); z++ {
		times = append(times, newTime(z, t))
	}

	return times
}

func newTime(z uint8, t time.Time) TimeAndSpace {
	switch z {
	case 0:
		return TimeAndSpace{
			TZ: 0,
			TX: uint16(t.Year()),
		}
	case 1:
		return TimeAndSpace{
			TZ: 1,
			TX: uint16(t.Year()),
			TY: quarterForMonth(t.Month()),
		}
	case 2:
		return TimeAndSpace{
			TZ: 2,
			TX: uint16(t.Year()),
			TY: uint16(t.Month()),
		}
	case 3:
		_, week := t.ISOWeek()
		return TimeAndSpace{
			TZ: 3,
			TX: uint16(t.Year()),
			TY: uint16(week),
		}
	case 4:
		return TimeAndSpace{
			TZ: 4,
			TX: uint16(t.Year()),
			TY: uint16(t.YearDay()),
		}
	default:
		panic("out of bounds zoom")
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
