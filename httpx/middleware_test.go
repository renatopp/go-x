package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// recordingHandler appends its name to order each time it's invoked.
type recordingHandler struct {
	name  string
	order *[]string
}

func (h recordingHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	*h.order = append(*h.order, h.name)
}

// recordingMiddleware appends "<name>:before" then calls next, then appends
// "<name>:after", letting tests assert nesting order.
func recordingMiddleware(name string, order *[]string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			*order = append(*order, name+":before")
			if next != nil {
				next.ServeHTTP(res, req)
			}
			*order = append(*order, name+":after")
		})
	}
}

func TestChainRunsMiddlewaresInDeclarationOrder(t *testing.T) {
	var order []string
	final := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		order = append(order, "final")
	})

	handler := Chain(recordingMiddleware("a", &order), recordingMiddleware("b", &order))(final)
	handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))

	want := []string{"a:before", "b:before", "final", "b:after", "a:after"}
	if !equalSlices(order, want) {
		t.Errorf("expected order %v, got %v", want, order)
	}
}

func TestChainWithNoMiddlewaresReturnsFinalHandlerUnchanged(t *testing.T) {
	var called bool
	final := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		called = true
	})

	handler := Chain()(final)
	handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))

	if !called {
		t.Error("expected final handler to run")
	}
}

func TestAsMiddlewareRunsHandlerBeforeNext(t *testing.T) {
	var order []string
	self := recordingHandler{name: "self", order: &order}
	next := recordingHandler{name: "next", order: &order}

	handler := AsMiddleware(self)(next)
	handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))

	if got := []string{"self", "next"}; !equalSlices(order, got) {
		t.Errorf("expected order %v, got %v", got, order)
	}
}

func TestAsMiddlewareWithNilNext(t *testing.T) {
	var order []string
	self := recordingHandler{name: "self", order: &order}

	handler := AsMiddleware(self)(nil)
	handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))

	if got := []string{"self"}; !equalSlices(order, got) {
		t.Errorf("expected order %v, got %v", got, order)
	}
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

