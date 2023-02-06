package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"strings"
)

// NodeIDContextKey is used to keep node ID in the system context
type NodeIDContextKey string

const (
	NodeID NodeIDContextKey = "node_id"
)

// MakeIdentifier creates uuid for job.
func MakeIdentifier() string {
	b := make([]byte, 12)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}

// IsEmptyStr check if the specified str is empty (len ==0) after triming prefix and suffix spaces.
func IsEmptyStr(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// ReadEnv return the value of env variable.
func ReadEnv(key string) string {
	return os.Getenv(key)
}

// FileExists check if the specified exists.
func FileExists(file string) bool {
	if !IsEmptyStr(file) {
		_, err := os.Stat(file)
		if err == nil {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}

		return true
	}

	return false
}

func DirExists(path string) bool {
	if IsEmptyStr(path) {
		return false
	}

	f, err := os.Stat(path)
	if err != nil {
		return false
	}

	return f.IsDir()
}