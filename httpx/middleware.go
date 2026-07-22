package httpx

import "net/http"

// Middleware wraps a http.Handler with additional behavior, producing a new
// http.Handler that delegates to it.
type Middleware func(http.Handler) http.Handler

// AsMiddleware adapts a http.Handler into a Middleware that runs handler
// first and, if given a next handler, runs it afterward against the same
// request and response.
func AsMiddleware(handler http.Handler) Middleware {
	return func(next http.Handler) http.Handler {
		return &middlewareWrapper{self: handler, next: next}
	}
}

// Chain combines middlewares into a single http.Handler that runs each one
// in order against the same request and response.
func Chain(middlewares ...http.Handler) http.Handler {
	if len(middlewares) == 0 {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {})
	}

	result := middlewares[len(middlewares)-1]
	for i := len(middlewares) - 2; i >= 0; i-- {
		result = &middlewareWrapper{self: middlewares[i], next: result}
	}
	return result
}

type middlewareWrapper struct {
	self http.Handler
	next http.Handler
}

func (m *middlewareWrapper) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	m.self.ServeHTTP(res, req)
	if m.next != nil {
		m.next.ServeHTTP(res, req)
	}
}
