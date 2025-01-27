package pkg

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		item     string
		expected bool
	}{
		{
			name:     "empty slice",
			slice:    []string{},
			item:     "test",
			expected: false,
		},
		{
			name:     "single item found",
			slice:    []string{"test"},
			item:     "test",
			expected: true,
		},
		{
			name:     "single item not found",
			slice:    []string{"test"},
			item:     "other",
			expected: false,
		},
		{
			name:     "multiple items found",
			slice:    []string{"test1", "test2", "test3"},
			item:     "test2",
			expected: true,
		},
		{
			name:     "multiple items not found",
			slice:    []string{"test1", "test2", "test3"},
			item:     "test4",
			expected: false,
		},
		{
			name:     "empty string found",
			slice:    []string{"", "test"},
			item:     "",
			expected: true,
		},
		{
			name:     "empty string not found",
			slice:    []string{"test1", "test2"},
			item:     "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.slice, tt.item)
			if got != tt.expected {
				t.Errorf("Contains() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
