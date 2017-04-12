package date

import (
	"fmt"
	"time"
)

type Date struct {
	tm time.Time
}

const timeFmt = "2006-01-02"

func Today() Date {
	return FromTime(time.Now())
}

func TodayIn(loc *time.Location) Date {
	return FromTime(time.Now().In(loc))
}

func Parse(s string) (Date, error) {
	tm, err := time.ParseInLocation(timeFmt, s, time.UTC)
	if err != nil {
		return Date{}, err
	} else {
		return Date{tm}, err
	}
}

func MustParse(s string) Date {
	day, err := Parse(s)
	if err != nil {
		panic(fmt.Sprintf("Invalid date format: %#v", s))
	} else {
		return day
	}
}

func Make(year int, month time.Month, day int) Date {
	return Date{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

func FromTime(tm time.Time) Date {
	y, m, d := tm.Date()
	return Date{time.Date(y, m, d, 0, 0, 0, 0, time.UTC)}
}

func (d Date) Date() (year int, month time.Month, day int) {
	return d.tm.Date()
}

func (d Date) In(loc *time.Location) time.Time {
	if loc == time.UTC {
		return d.tm
	}
	yr, mo, dy := d.tm.Date()
	return time.Date(yr, mo, dy, 0, 0, 0, 0, loc)
}

func (d Date) InUTC() time.Time {
	return d.tm
}

func (d Date) String() string {
	return d.tm.Format(timeFmt)
}

func (d Date) StringOr(zero string) string {
	if d.IsZero() {
		return zero
	} else {
		return d.String()
	}
}

func (d Date) IsZero() bool {
	return d.tm.IsZero()
}

func (d Date) Equal(u Date) bool {
	return d.tm.Equal(u.tm)
}

func (d Date) Before(u Date) bool {
	return d.tm.Before(u.tm)
}

func (d Date) After(u Date) bool {
	return d.tm.After(u.tm)
}

func (d Date) AddDays(days int) Date {
	return Date{d.tm.AddDate(0, 0, days)}
}

func (d Date) Add(years, months, days int) Date {
	return Date{d.tm.AddDate(years, months, days)}
}

func (d Date) Next() Date {
	return d.AddDays(1)
}

func (d Date) Prev() Date {
	return d.AddDays(-1)
}

// satisfies flag.Value interface
func (d *Date) Set(s string) error {
	u, err := Parse(s)
	if err != nil {
		return err
	}
	*d = u
	return nil
}

func Min(a, b Date) Date {
	if a.Before(b) {
		return a
	} else {
		return b
	}
}

func Max(a, b Date) Date {
	if a.After(b) {
		return a
	} else {
		return b
	}
}
