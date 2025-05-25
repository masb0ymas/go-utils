package pkg

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// StringToInt32 converts a string to int32 with a default value if conversion fails
func StringToInt32(str string, defaultValue ...int32) int32 {
	// Set default value to 0 if not provided
	def := int32(0)
	if len(defaultValue) > 0 {
		def = defaultValue[0]
	}

	// Convert string to int64
	val, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return def
	}

	return int32(val)
}

// StringToInt64 converts a string to int64 with a default value if conversion fails
func StringToInt64(str string, defaultValue ...int64) int64 {
	// Set default value to 0 if not provided
	def := int64(0)
	if len(defaultValue) > 0 {
		def = defaultValue[0]
	}

	// Convert string to int64
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return def
	}

	return int64(val)
}

// StringToTime converts string to time.Time
func StringToTime(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, nil
	}

	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, timeStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, errors.New("invalid time format")
}

// IntToString converts an integer to a string
func IntToString(n int) string {
	return strconv.Itoa(n)
}

// JoinStrings joins a slice of strings with a separator
func JoinStrings(strs []string, sep string) string {
	return strings.Join(strs, sep)
}
