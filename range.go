package date

import (
	"time"
)

type Range struct {
	// Start is the first day that belongs to this range. Can be zero if there's no lower bound.
	Start Date

	// End is one day after the last day that belongs to this range. Can be zero if there's no upper bound.
	End Date
}

func HalfOpenRange(start, end Date) Range {
	return Range{start, end}
}

func ClosedRange(start, end Date) Range {
	if !end.IsZero() {
		end = end.Next()
	}
	return Range{start, end}
}

func (r Range) Contains(tm time.Time) bool {
	if !r.Start.IsZero() {
		if r.Start.tm.After(tm) {
			return false
		}
	}

	if !r.End.IsZero() {
		if !tm.Before(r.End.tm) {
			return false
		}
	}

	return true
}

func (r Range) ContainsDate(d Date) bool {
	return r.Contains(d.tm)
}
