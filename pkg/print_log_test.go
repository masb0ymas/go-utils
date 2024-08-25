package pkg

import (
	"fmt"
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
		options  []Options
		expected string
	}{
		{
			name:     "Default options",
			title:    "Test Title",
			message:  "Test Message",
			options:  nil,
			expected: PrintLog("Test Title", "Test Message"),
		},
		{
			name:    "Success label",
			title:   "Success Title",
			message: "Success Message",
			options: []Options{
				{Label: "success"},
			},
			expected: PrintLog("Success Title", "Success Message"),
		},
		{
			name:    "Warning label",
			title:   "Warning Title",
			message: "Warning Message",
			options: []Options{
				{Label: "warning"},
			},
			expected: PrintLog("Warning Title", "Warning Message"),
		},
		{
			name:    "Error label",
			title:   "Error Title",
			message: "Error Message",
			options: []Options{
				{Label: "error"},
			},
			expected: PrintLog("Error Title", "Error Message"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := PrintLog(tc.title, tc.message, tc.options...)
			fmt.Println("RESULT: ", result)
			fmt.Println("EXPECTED: ", tc.expected)
			fmt.Printf("\n")

			if result != tc.expected {
				t.Errorf("PrintLog() = %v, want %v", result, tc.expected)
			}
		})
	}
}

// Helper function to create a pointer to a string
func stringPtr(s string) *string {
	return &s
}
