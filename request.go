package tomorrow

import "time"

const (
	defaultTimestep = "1h"
)

var (
	defaultRequest = Request{
		Timesteps: defaultTimestep,
		Fields:    fields,
	}
)

type Request struct {
	Timesteps string
	Units     string
	Lat       float64
	Lng       float64
	Start     time.Time
	End       time.Time
	Fields    []string
}
