package validate

func IsValidDate(date *string) bool {
	return date != nil && *date != "" && *date != "0001-01-01T00:00:00Z"
}
