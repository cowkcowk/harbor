package utils

import "strings"

// TrimLower ...
func TrimLower(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}
