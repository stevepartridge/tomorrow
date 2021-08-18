package tomorrow

import (
	"fmt"
	"os"
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

	RateLimitDay      int
	RateLimitHour     int
	RateRemainingDay  int
	RateRemainingHour int
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
