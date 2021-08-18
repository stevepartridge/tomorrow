package tomorrow

import "time"

// Response is the top level object for all Tomorrow API responses
type Response struct {
	Data Data `json:"data,omitempty"`
}

// Data is the common structure for most responses
type Data struct {
	Timelines []Timeline `json:"timelines,omitempty"`
	Events    []Event    `json:"events"`
}

// Timeline struct
type Timeline struct {
	Timestep  string     `json:"timestep,omitempty"`
	StartTime time.Time  `json:"startTime,omitempty"`
	EndTime   time.Time  `json:"endTime,omitempty"`
	Intervals []Interval `json:"intervals,omitempty"`
}

type Event struct {
	Insight    string    `json:"insight"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	Severity   string    `json:"severity,omitempty"`
	Certainty  string    `json:"certainty,omitempty"`
	Urgency    string    `json:"urgency,omitempty"`
	// Eventvalues struct {
	// 	Origin         string  `json:"origin"`
	// 	Title          string  `json:"title"`
	// 	Radiativepower float64 `json:"radiativePower"`
	// 	Location       struct {
	// 		Type        string    `json:"type"`
	// 		Coordinates []float64 `json:"coordinates"`
	// 	} `json:"location"`
	// 	Distance  float64 `json:"distance"`
	// 	Direction float64 `json:"direction"`
	// } `json:"eventValues,omitempty"`
	// Eventvalues struct {
	// 	Precipitationintensity int `json:"precipitationIntensity"`
	// 	Mepindex               int `json:"mepIndex"`
	// } `json:"eventValues,omitempty"`
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
	PrecipitationIntensity float64 `json:"precipitationIntensity,omitempty"`
	PrecipitationType      int     `json:"precipitationType,omitempty"`
	WindSpeed              float64 `json:"windSpeed,omitempty"`
	WindGust               float64 `json:"windGust,omitempty"`
	WindDirection          float64 `json:"windDirection,omitempty"`
	CloudCover             float64 `json:"cloudCover,omitempty"`
	CloudBase              float64 `json:"cloudBase,omitempty"`
	CloudCeiling           float64 `json:"cloudCeiling,omitempty"`
}
