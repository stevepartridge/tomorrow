package tomorrow

import (
	"strings"
	"testing"
	"time"
)

func TestOptions(t *testing.T) {

	t.Run("Fields with empty list", func(r *testing.T) {

		err := Fields()(&Request{})
		if err == nil {
			t.Error("Expected error but saw none")
		}
		if !strings.Contains(err.Error(), ErrFieldsEmpty.Error()) {
			t.Errorf("Expected error to contain %s but saw %s", ErrFieldsEmpty.Error(), err.Error())
		}

	})

	t.Run("Fields with basic list", func(r *testing.T) {

		req := &Request{}
		err := Fields("temperature", "temperatureApparent")(req)
		if err != nil {
			t.Errorf("Unexpected error but saw %s", err.Error())
		}

		if len(req.Fields) != 2 {
			t.Errorf("Expected fields count to be 2 but saw %d", len(req.Fields))
		}

	})

	t.Run("Range with zero start", func(r *testing.T) {

		req := &Request{}
		err := Range(time.Time{}, time.Time{})(req)
		if err == nil {
			t.Error("Expected error but saw none")
		}
		if !strings.Contains(err.Error(), ErrRangeStartIsZero.Error()) {
			t.Errorf("Expected error to contain %s but saw %s", ErrRangeStartIsZero.Error(), err.Error())
		}

	})

	t.Run("Range with zero end", func(r *testing.T) {

		req := &Request{}
		err := Range(time.Now(), time.Time{})(req)
		if err == nil {
			t.Error("Expected error but saw none")
		}
		if !strings.Contains(err.Error(), ErrRangeEndIsZero.Error()) {
			t.Errorf("Expected error to contain %s but saw %s", ErrRangeEndIsZero.Error(), err.Error())
		}

	})

	t.Run("Range with valid start and end", func(r *testing.T) {

		req := &Request{}
		err := Range(time.Now(), time.Now().Add(time.Second*1))(req)
		if err != nil {
			t.Errorf("Unexpected error but saw %s", err.Error())
		}

	})

}
