package tomorrow

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	// Imperial constant
	Imperial = "imperial"
	// Metric constant
	Metric = "metric"
)

// Client is the Tomorrow API client struct
type Client struct {
	apiKey string
	units  string
}

// New is the generic method to create a Tomorrow API client with a provided unit
func New(units string, apiKey ...string) (*Client, error) {
	c := Client{
		apiKey: os.Getenv(EnvAPIKey),
	}
	if len(apiKey) > 0 {
		if apiKey[0] != "" {
			c.apiKey = apiKey[0]
		}
	}

	switch units {
	case Imperial, Metric:
		c.units = units
	default:
		return nil, fmt.Errorf(ErrInvalidUnits.Error(), units)
	}

	return &c, nil
}

// NewImperial is a helper to use a Tomorrow API client returning imperial (Fahrenheit) values
func NewImperial(apiKey ...string) (*Client, error) {
	return New(Imperial, apiKey...)
}

// NewMetric is a helper to use a Tomorrow API client returning metrics values
func NewMetric(apiKey ...string) (*Client, error) {
	return New(Metric, apiKey...)
}

// Response is the top level object for all Tomorrow API responses
type Response struct {
	Data Data `json:"data,omitempty"`
}

// Data is the common structure for most responses
type Data struct {
	Timelines []Timeline `json:"timelines,omitempty"`
}

// Timeline struct
type Timeline struct {
	Timestep  string     `json:"timestep,omitempty"`
	StartTime time.Time  `json:"startTime,omitempty"`
	EndTime   time.Time  `json:"endTime,omitempty"`
	Intervals []Interval `json:"intervals,omitempty"`
}

// Interval struct
type Interval struct {
	StartTime time.Time `json:"startTime"`
	Values    Values    `json:"values,omitempty"`
}

// Timestep struct
type Timestep struct {
	StartTime time.Time `json:"startTime,omitempty" yaml:""`
	Values    Values    `json:"values,omitempty" yaml:""`
}

// Values struct
type Values struct {
	Temperature            float64 `json:"temperature,omitempty" yaml:"temperature"`
	TemperatureApparent    float64 `json:"temperatureApparent,omitempty"`
	Humidty                float64 `json:"humidity,omitempty" yaml:"humidty"`
	WeatherCode            int     `json:"weatherCode,omitempty" yaml:"weather_code"`
	PrecipitationIntensity int     `json:"precipitationIntensity,omitempty"`
	PrecipitationType      int     `json:"precipitationType,omitempty"`
	WindSpeed              float64 `json:"windSpeed,omitempty"`
	WindGust               float64 `json:"windGust,omitempty"`
	WindDirection          float64 `json:"windDirection,omitempty"`
	CloudCover             float64 `json:"cloudCover,omitempty"`
	CloudBase              float64 `json:"cloudBase,omitempty"`
	CloudCeiling           float64 `json:"cloudCeiling,omitempty"`
}

// GetTemp is a basic endpoint helper for retreiving temp intervals for a lat/lng combo
func (c *Client) GetTemp(lat, lng float64) (Timeline, error) {

	resp, err := c.call(http.MethodGet, "/timelines", nil, map[string]string{
		"units":     c.units,
		"fields":    strings.Join(fields, ","),
		"timesteps": "1h",
		"location":  fmt.Sprintf("%f,%f", lng, lat),
	})

	if err != nil {
		return Timeline{}, err
	}

	if len(resp.Data.Timelines) < 1 {
		return Timeline{}, errors.New("get.temp: Timelines not found")
	}

	return resp.Data.Timelines[0], nil
}
