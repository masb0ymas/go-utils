package validate

import (
	"testing"
)

func TestIsValidDate(t *testing.T) {
	tests := []struct {
		name     string
		date     *string
		expected bool
	}{
		{
			name:     "valid date",
			date:     strPtr("2023-05-20T15:04:05Z"),
			expected: true,
		},
		{
			name:     "empty date",
			date:     strPtr(""),
			expected: false,
		},
		{
			name:     "zero date",
			date:     strPtr("0001-01-01T00:00:00Z"),
			expected: false,
		},
		{
			name:     "nil date",
			date:     nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidDate(tt.date)
			if result != tt.expected {
				t.Errorf("IsValidDate() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}
