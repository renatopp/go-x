package httpx

import (
	"net/http"
	"slices"
)

// Router is an HTTP router built on http.ServeMux. It implements
// http.Handler and can be passed directly to http.Server or as a Middleware
// target. Middlewares registered with Use apply, in declaration order, to
// every route registered afterwards on this Router (or a Router derived
// from it via With or Group); they never affect routes already registered.
type Router struct {
	mux *http.ServeMux
	mws []Middleware
}

// NewRouter creates an empty Router.
func NewRouter() *Router {
	return &Router{mux: http.NewServeMux()}
}

// WithMiddleware appends mws to the router's middleware stack.
func (r *Router) WithMiddleware(mws ...Middleware) *Router {
	r.mws = append(r.mws, mws...)
	return r
}

// WithHandle registers h for pattern, wrapped with the router's current
// middleware stack.
func (r *Router) WithHandle(pattern string, h http.Handler) *Router {
	r.mux.Handle(pattern, Chain(r.mws...)(h))
	return r
}

// WithHandleFunc registers h for pattern, wrapped with the router's current
// middleware stack.
func (r *Router) WithHandleFunc(pattern string, h http.HandlerFunc) *Router {
	return r.WithHandle(pattern, h)
}

// WithGroup returns a new Router that shares the same underlying mux but has mws
// appended to its middleware stack. It is useful for applying extra
// middlewares to a handful of routes without affecting the parent Router.
func (r *Router) WithGroup(mws ...Middleware) *Router {
	return &Router{mux: r.mux, mws: append(slices.Clone(r.mws), mws...)}
}

// ServeHTTP implements http.Handler.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
