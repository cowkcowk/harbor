package log

import (
	"net/http"

	"github.com/cowkcowk/harbor/src/lib/log"
	"github.com/cowkcowk/harbor/src/server/middleware"
)

// Middleware middleware which add logger to context
func Middleware() func(http.Handler) http.Handler {
	return middleware.New(func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		rid := r.Header.Get("X-Request-ID")
		if rid != "" {
			logger := log.G(r.Context())
			logger.Debugf("attach request id %s to the logger for the request %s %s", rid, r.Method, r.URL.Path)

			ctx := log.WithLogger(r.Context(), logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}