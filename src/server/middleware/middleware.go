package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(middlewares ...Middleware) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			h = middlewares[i](h)
		}

		return h
	}
}

// WithMiddlewares apply the middlewares to the handler.
// The middlewares are executed in the order that they are applied
func WithMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	return Chain(middlewares...)(handler)
}

func New(fn func(http.ResponseWriter, *http.Request, http.Handler)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fn(w, r, next)
		})
	}
}

func BeforeRequest(hook func(*http.Request) error) func(http.Handler) http.Handler {
	return New(func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		next.ServeHTTP(w, r)
	})
}

func AfterResponse(hook func(http.ResponseWriter, *http.Request, int) error) func(http.Handler) http.Handler {
	return New(func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		next.ServeHTTP(w, r)
	})
}