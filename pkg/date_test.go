package pkg

import (
	"testing"
	"time"
)

func TestTimeIn(t *testing.T) {
	tests := []struct {
		name           string
		country        string
		expectedOffset int
	}{
		{
			name:           "Indonesia",
			country:        "ID",
			expectedOffset: 7 * 60 * 60, // UTC+7
		},
		{
			name:           "Kuala_Lumpur",
			country:        "MY",
			expectedOffset: 8 * 60 * 60, // UTC+8
		},
		{
			name:           "Singapore",
			country:        "SG",
			expectedOffset: 8 * 60 * 60, // UTC+8
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TimeIn(tt.country)

			// Get the offset in seconds
			_, offset := result.Zone()

			if offset != tt.expectedOffset {
				t.Errorf("TimeIn(%s) got timezone offset %d, want %d", tt.country, offset, tt.expectedOffset)
			}

			// Check if the returned time is within a reasonable range (1 second) of the current time
			now := time.Now().In(result.Location())
			diff := now.Sub(result).Abs()
			if diff > time.Second {
				t.Errorf("TimeIn(%s) returned time %v, which is more than 1 second different from current time %v", tt.country, result, now)
			}
		})
	}
}

func TestTimeInInvalidCountry(t *testing.T) {
	// This test will panic due to the log.Fatal in the original function
	// We need to recover from the panic to prevent the test from crashing
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	TimeIn("InvalidCountry")
}
