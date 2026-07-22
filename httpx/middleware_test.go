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

func TestAsMiddlewareRunsHandlerBeforeNext(t *testing.T) {
	var order []string
	self := recordingHandler{name: "self", order: &order}
	next := recordingHandler{name: "next", order: &order}

	mw := AsMiddleware(self)
	handler := mw(next)
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

func TestChainRunsHandlersInOrder(t *testing.T) {
	var order []string
	h1 := recordingHandler{name: "h1", order: &order}
	h2 := recordingHandler{name: "h2", order: &order}
	h3 := recordingHandler{name: "h3", order: &order}

	handler := Chain(h1, h2, h3)
	handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))

	if got := []string{"h1", "h2", "h3"}; !equalSlices(order, got) {
		t.Errorf("expected order %v, got %v", got, order)
	}
}

func TestChainWithSingleHandler(t *testing.T) {
	var order []string
	h1 := recordingHandler{name: "h1", order: &order}

	handler := Chain(h1)
	handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))

	if got := []string{"h1"}; !equalSlices(order, got) {
		t.Errorf("expected order %v, got %v", got, order)
	}
}

func TestChainWithNoHandlersDoesNotPanic(t *testing.T) {
	handler := Chain()
	handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))
}

func TestChainSharesRequestAndResponse(t *testing.T) {
	h1 := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		req.Header.Set("X-Seen", "h1")
		res.Header().Set("X-From", "h1")
	})
	h2 := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Header.Get("X-Seen") != "h1" {
			t.Errorf("expected h2 to see request mutated by h1")
		}
		res.Header().Set("X-From", res.Header().Get("X-From")+",h2")
	})

	res := httptest.NewRecorder()
	Chain(h1, h2).ServeHTTP(res, httptest.NewRequest(http.MethodGet, "/", nil))

	if got := res.Header().Get("X-From"); got != "h1,h2" {
		t.Errorf("expected shared response header %q, got %q", "h1,h2", got)
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
