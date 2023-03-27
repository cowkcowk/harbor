package requestid

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/goharbor/harbor/src/server/middleware"
)

// HeaderXRequestID X-Request-ID header
const HeaderXRequestID = "X-Request-ID"

func Middleware() func(http.Handler) http.Handler {
	return middleware.New(func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		rid := r.Header.Get(HeaderXRequestID)
		if rid == "" {
			rid = uuid.New().String()
			r.Header.Set(HeaderXRequestID, rid)
		}
		w.Header().Set(HeaderXRequestID, rid)

		next.ServeHTTP(w, r)
	})
}
