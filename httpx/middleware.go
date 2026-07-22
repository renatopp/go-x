package httpx

import "net/http"

// Middleware wraps an http.Handler with additional behavior, producing a new
// http.Handler that delegates to it. Because Middleware is defined as
// func(http.Handler) http.Handler, any middleware written for the standard
// net/http ecosystem can be used directly as a Middleware, and vice versa.
type Middleware func(http.Handler) http.Handler

// Chain composes mws into a single Middleware. When the result is applied to
// a final handler, the middlewares run in the order they are given: the
// first middleware is outermost and runs first, calling into the next one
// until the final handler is reached.
func Chain(mws ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		h := final
		for i := len(mws) - 1; i >= 0; i-- {
			h = mws[i](h)
		}
		return h
	}
}

// AsMiddleware adapts an http.Handler that doesn't know about "next" into a
// Middleware: h runs first against the request and response, then control
// passes to the next handler in the chain, if any.
func AsMiddleware(h http.Handler) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			h.ServeHTTP(w, req)
			if next != nil {
				next.ServeHTTP(w, req)
			}
		})
	}
}
