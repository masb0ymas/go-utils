package pkg

import (
	"testing"

	"github.com/fatih/color"
)

func TestPrintLog(t *testing.T) {
	// Disable color for testing
	color.NoColor = true

	tests := []struct {
		name     string
		title    string
		message  string
		options  []PrintOptions
		expected string
	}{
		{
			name:     "Default options",
			title:    "Test Title",
			message:  "Test Message",
			options:  nil,
			expected: Println("Test Title", "Test Message"),
		},
		{
			name:    "Success label",
			title:   "Success Title",
			message: "Success Message",
			options: []PrintOptions{
				{
					Label:   stringPtr("success"),
					TagText: stringPtr("success"),
				},
			},
			expected: Println("Success Title", "Success Message", PrintOptions{Label: "success", TagText: "success"}),
		},
		{
			name:    "Warning label",
			title:   "Warning Title",
			message: "Warning Message",
			options: []PrintOptions{
				{
					Label:   stringPtr("warning"),
					TagText: stringPtr("warning"),
				},
			},
			expected: Println("Warning Title", "Warning Message", PrintOptions{Label: "warning", TagText: "warning"}),
		},
		{
			name:    "Error label",
			title:   "Error Title",
			message: "Error Message",
			options: []PrintOptions{
				{
					Label:   stringPtr("error"),
					TagText: stringPtr("error"),
				},
			},
			expected: Println("Error Title", "Error Message", PrintOptions{Label: "error", TagText: "error"}),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Println(tc.title, tc.message, tc.options...)

			if result != tc.expected {
				t.Errorf("\nPrintLog() = %v\nexpected = %v", result, tc.expected)
			}
		})
	}
}

// Helper function to create a pointer to a string
func stringPtr(s string) string {
	return s
}
