package interval

import (
	"cmp"
	"slices"
)

// Interval represents a left closed, right open interval [Lower, Upper).
type Interval struct {
	Lower int64
	Upper int64
}

// Compare compares two intervals.
//
// The lower bounds are compared first. If they are equal, the upper bounds are compared.
func (i Interval) Compare(rhs Interval) int {
	if st := cmp.Compare(i.Lower, rhs.Lower); st != 0 {
		return st
	}
	return cmp.Compare(i.Upper, rhs.Upper)
}

// Shift shifts the interval by the given offset.
func (i Interval) Shift(offset int64) Interval {
	return Interval{i.Lower + offset, i.Upper + offset}
}

// Merge sorts and merges contiguous or overlapping intervals.
func Merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

	slices.SortFunc(intervals, func(l, r Interval) int { return l.Compare(r) })
	var result []Interval
	for _, i := range intervals {
		if len(result) == 0 || result[len(result)-1].Upper < i.Lower {
			result = append(result, i)
		} else if result[len(result)-1].Upper < i.Upper {
			result[len(result)-1].Upper = i.Upper
		}
	}
	return result
}
