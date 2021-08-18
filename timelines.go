package tomorrow

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetTemp is a basic endpoint helper for retreiving temp intervals for a lat/lng combo
func (c *Client) GetTemp(lat, lng float64, opts ...Option) (Timeline, error) {
	r := defaultRequest
	r.Units = c.units
	r.Lat = lat
	r.Lng = lng

	for _, opt := range opts {
		if err := opt(&r); err != nil {
			return Timeline{}, err
		}
	}

	resp, err := c.call(http.MethodGet, "/timelines", nil, map[string]string{
		"units":     r.Units,
		"fields":    strings.Join(r.Fields, ","),
		"timesteps": r.Timesteps,
		"location":  fmt.Sprintf("%f,%f", r.Lat, r.Lng),
	})

	if err != nil {
		return Timeline{}, err
	}

	if len(resp.Data.Timelines) < 1 {
		return Timeline{}, errors.New("get.temp: Timelines not found")
	}

	return resp.Data.Timelines[0], nil
}
