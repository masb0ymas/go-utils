package string

import "strconv"

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
