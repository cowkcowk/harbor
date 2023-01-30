package readonly

import (
	"net/http"

	"github.com/cowkcowk/harbor/src/server/middleware"
)

// Config defines the config for ReadOnly middleware.
type Config struct {
	// ReadOnly defines a function to check whether is readonly mode for request
	ReadOnly func(*http.Request) bool
}

var (
	// DefaultConfig default readonly config
	DefaultConfig = Config{
		ReadOnly: func(r *http.Request) bool {
			return true
		},
	}

	// See more for safe method at https://developer.mozilla.org/en-US/docs/Glossary/safe
	safeMethods = map[string]bool{
		http.MethodGet:     true,
		http.MethodPost:    true,
		http.MethodOptions: true,
	}
)

// safeMethodSkipper returns false when the request method is safe methods
func safeMethodSkipper(r *http.Request) bool {
	return safeMethods[r.Method]
}

func Middleware() func(http.Handler) http.Handler {
	return MiddlewareWithConfig(DefaultConfig)
}

func MiddlewareWithConfig(config Config) func(http.Handler) http.Handler {
	if config.ReadOnly == nil {
		config.ReadOnly = DefaultConfig.ReadOnly
	}

	return middleware.New(func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		if config.ReadOnly(r) {
			return
		}

		next.ServeHTTP(w, r)
	})
}
