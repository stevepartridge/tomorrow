package tomorrow

import (
	"os"
	"testing"
)

const (
	testLat1    = -121.0249798
	testLng1    = 38.8880406
	testBadLat1 = -1.2345
	testBadLng1 = 9.8765
)

func TestNewTomorrow(t *testing.T) {

	t.Run("With bad unit", func(r *testing.T) {

		_, err := New("nonsense unit")
		if err == nil {
			t.Error("Expected error but did not see one")
		}

	})

	t.Run("With Imperial unit", func(r *testing.T) {

		_, err := NewImperial("nonsense key")
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

	})

	t.Run("With Metric unit", func(r *testing.T) {

		_, err := NewMetric("nonsense key")
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

	})

	t.Run("With API Key", func(r *testing.T) {

		err := os.Setenv(EnvAPIKey, os.Getenv("TEST_TOMORROW_API_KEY"))
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		tmw, err := New(Imperial)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		temp, err := tmw.GetTemp(testLat1, testLng1)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		if len(temp.Intervals) == 0 {
			t.Errorf("Expected temperature results but found none")
		}

		err = os.Unsetenv(EnvAPIKey)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	})

	t.Run("With bad Lat Lng", func(r *testing.T) {

		err := os.Setenv(EnvAPIKey, os.Getenv("TEST_TOMORROW_API_KEY"))
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		tmw, err := New(Imperial)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		_, err = tmw.GetTemp(testBadLat1, testBadLng1)
		if err == nil {
			t.Error("Expected error but did not see one")
		}

		err = os.Unsetenv(EnvAPIKey)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	})

}
