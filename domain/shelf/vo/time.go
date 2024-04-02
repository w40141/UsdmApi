// Package vo is defined for value object.
package vo

import "time"

// FromStringToTime reads an ID from its string representation
func FromStringToTime(s string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}, err
	}
	return t, err
}
