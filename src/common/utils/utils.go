package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// TrimLower ...
func TrimLower(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}

// ParseProjectIDOrName parses value to ID(int64) or name(string)
func ParseProjectIDOrName(value interface{}) (int64, string, error) {
	if value == nil {
		return 0, "", errors.New("harborIDOrName is nil")
	}

	var id int64
	var name string
	switch v := value.(type) {
	case int, int64:
		id = reflect.ValueOf(v).Int()
	case string:
		name = value.(string)
	default:
		return 0, "", fmt.Errorf("unsupported type")
	}
	return id, name, nil
}
