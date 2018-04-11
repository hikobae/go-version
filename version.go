package version

import (
	"strconv"
	"strings"
)

// Version represents version string
type Version struct {
	str      string
	segments []int
}

// NewVersion returns new Version given a version string.
// If fail to parse string, returns nil.
func NewVersion(s string) (*Version, error) {
	segmentsStr := strings.Split(s, ".")
	segments := make([]int, len(segmentsStr))
	for i, seg := range segmentsStr {
		n, err := strconv.Atoi(seg)
		if err != nil {
			return nil, err
		}
		segments[i] = n
	}
	return &Version{
		str:      s,
		segments: segments,
	}, nil
}

// String returns string of v
func (v *Version) String() string {
	return v.str
}

// Cmp compares v and o and returns:
//
//   -1 if v <  o
//    0 if v == 0
//    1 if v >  0
func (v *Version) Cmp(o *Version) int {
	if v.String() == o.String() {
		return 0
	}

	vLen := len(v.segments)
	oLen := len(o.segments)

	l := min(vLen, oLen)
	for i := 0; i < l; i++ {
		if v.segments[i] > o.segments[i] {
			return 1
		} else if v.segments[i] < o.segments[i] {
			return -1
		}
	}

	if vLen > l {
		if isAllZero(v.segments[l:]) {
			return 0
		}
		return 1
	}
	if oLen > l {
		if isAllZero(o.segments[l:]) {
			return 0
		}
		return -1
	}

	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isAllZero(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

// Equals returns true if v == o
func (v *Version) Equals(o *Version) bool {
	return v.Cmp(o) == 0
}

// LessThan true if v < o
func (v *Version) LessThan(o *Version) bool {
	return v.Cmp(o) < 0
}

// GreaterThan true if v > o
func (v *Version) GreaterThan(o *Version) bool {
	return v.Cmp(o) > 0
}
