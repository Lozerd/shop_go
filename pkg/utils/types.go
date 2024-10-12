package utils

import "strings"

func StringOrDefault(value string, default_if_empty string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return default_if_empty
	}

	return value
}
