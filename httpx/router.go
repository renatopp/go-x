package httpx

package web

import "net/http"

// Router is the root HTTP router. It implements http.Handler and can be passed
// directly to http.Server. Middlewares are applied in declaration order.
type Router struct {
	mux *http.ServeMux
	mws []Middleware
}

func NewRouter() *Router {
	return &Router{mux: http.NewServeMux()}
}

func (r *Router) WithMiddleware(mws ...Middleware) *Router {
	r.mws = append(r.mws, mws...)
	return r
}

func (r *Router) WithHandle(pattern string, h ...http.Handler) *Router {
	r.mux.Handle(pattern, Chain(h...)))
	return r
}

func (r *Router) WithHandleFunc(pattern string, h ...http.HandlerFunc) *Router {
	return r.WithHandle(pattern, Chain(h...))
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

// NewGroup creates a route group that inherits the router's middlewares.
// Additional middlewares can be added per group without affecting the router.
func (r *Router) NewGroup() *RouterGroup {
	return &RouterGroup{mux: r.mux, mws: sliceCopy(r.mws)}
}

// RouterGroup is a logical grouping of routes that share a middleware chain.
// It does not introduce path prefixes — full patterns must always be declared.
type RouterGroup struct {
	mux *http.ServeMux
	mws []Middleware
}

func (g *RouterGroup) WithMiddleware(mws ...Middleware) *RouterGroup {
	g.mws = append(g.mws, mws...)
	return g
}

func (g *RouterGroup) WithHandle(pattern string, h http.Handler, mws ...Middleware) *RouterGroup {
	g.mux.Handle(pattern, ChainMiddlewares(append(g.mws, mws...), h))
	return g
}

func (g *RouterGroup) WithHandleFunc(pattern string, h http.HandlerFunc, mws ...Middleware) *RouterGroup {
	return g.WithHandle(pattern, h, mws...)
}

// NewGroup creates a child group that inherits this group's middleware chain.
func (g *RouterGroup) NewGroup() *RouterGroup {
	return &RouterGroup{mux: g.mux, mws: sliceCopy(g.mws)}
}

func sliceCopy[T any](s []T) []T {
	c := make([]T, len(s))
	copy(c, s)
	return c
}
