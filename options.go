package tomorrow

import (
	"errors"
	"time"
)

var (
	ErrRangeStartIsZero = errors.New("option range: start time is zero")
	ErrRangeEndIsZero   = errors.New("option range: end time is zero")
	ErrFieldsEmpty      = errors.New("option fields: empty list")
)

// Option type
type Option func(r *Request) error

// Fields option
func Fields(fields ...string) func(*Request) error {
	return func(r *Request) error {
		if len(fields) == 0 {
			return ErrFieldsEmpty
		}
		r.Fields = fields
		return nil
	}
}

// Range option
func Range(start, end time.Time) func(*Request) error {
	return func(r *Request) error {
		if start.IsZero() {
			return ErrRangeStartIsZero
		}
		if end.IsZero() {
			return ErrRangeEndIsZero
		}
		r.Start = start
		r.End = end
		return nil
	}
}
