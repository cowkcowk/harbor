package router

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/beego/beego/v2/server/web"
	beegocontext "github.com/beego/beego/v2/server/web/context"

	"github.com/goharbor/harbor/src/server/middleware"
)

// ContextKeyInput ...
type ContextKeyInput struct{}

// NewRoute creates a new route
func NewRoute() *Route {
	return &Route{}
}

// Route stores the information that matches a request
type Route struct {
	parent      *Route
	methods     []string
	path        string
	middlewares []middleware.Middleware
}

// NewRoute returns a sub route based on the current one
func (r *Route) NewRoute() *Route {
	return &Route{
		parent: r,
	}
}

// Method sets the method that the route matches
func (r *Route) Method(method string) *Route {
	r.methods = append(r.methods, method)
	return r
}

// Path sets the path that the route matches. Path uses the beego router path pattern
func (r *Route) Path(path string) *Route {
	r.path = path
	return r
}

// Middleware sets the middleware that executed when handling the request
func (r *Route) Middleware(middleware middleware.Middleware) *Route {
	r.middlewares = append(r.middlewares, middleware)
	return r
}

func (r *Route) Handler(handler http.Handler) {
	methods := r.methods
	if len(methods) == 0 && r.parent != nil {
		methods = r.parent.methods
	}

	path := r.path
	if r.parent != nil {
		path = filepath.Join(r.parent.path, path)
	}

	var middlewares []middleware.Middleware
	if r.parent != nil {
		middlewares = r.parent.middlewares
	}

	middlewares = append(middlewares, r.middlewares...)
	filterFunc := web.FilterFunc(func(ctx *beegocontext.Context) {
		ctx.Request = ctx.Request.WithContext(
			context.WithValue(ctx.Request.Context(), ContextKeyInput{}, ctx.Input))
		// TODO remove the WithMiddlewares?
		middleware.WithMiddlewares(handler, middlewares...).
			ServeHTTP(ctx.ResponseWriter, ctx.Request)
	})

	if len(methods) == 0 {
		web.Any(path, filterFunc)
		return
	}

	for _, method := range methods {
		switch method {
		case http.MethodGet:
			web.Get(path, filterFunc)
		case http.MethodHead:
			web.Head(path, filterFunc)
		}
	}
}

// HandlerFunc sets the handler function that handles the request
func (r *Route) HandlerFunc(f http.HandlerFunc) {
	r.Handler(f)
}

// Param returns the beego router param by a given key from the context
func Param(ctx context.Context, key string) string {
	if ctx == nil {
		return ""
	}
	input, ok := ctx.Value(ContextKeyInput{}).(*beegocontext.BeegoInput)
	if !ok {
		return ""
	}
	return input.Param(key)
}
