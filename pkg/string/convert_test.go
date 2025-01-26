package string

import "testing"

func TestStringToInt32(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		defaultValue []int32
		expected     int32
	}{
		{
			name:     "valid number",
			input:    "123",
			expected: 123,
		},
		{
			name:         "valid number with default",
			input:        "456",
			defaultValue: []int32{999},
			expected:     456,
		},
		{
			name:         "invalid number uses default",
			input:        "abc",
			defaultValue: []int32{999},
			expected:     999,
		},
		{
			name:     "invalid number uses zero default",
			input:    "abc",
			expected: 0,
		},
		{
			name:     "empty string uses zero default",
			input:    "",
			expected: 0,
		},
		{
			name:     "max int32",
			input:    "2147483647",
			expected: 2147483647,
		},
		{
			name:     "min int32",
			input:    "-2147483648",
			expected: -2147483648,
		},
		{
			name:     "overflow uses default",
			input:    "2147483648",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringToInt32(tt.input, tt.defaultValue...)
			if got != tt.expected {
				t.Errorf("StringToInt32() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
