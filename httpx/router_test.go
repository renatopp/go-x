package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterAppliesMiddlewaresToRegisteredRoutes(t *testing.T) {
	var order []string
	r := NewRouter()
	r.WithMiddleware(recordingMiddleware("mw1", &order), recordingMiddleware("mw2", &order))
	r.WithHandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		order = append(order, "handler")
	})

	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))

	want := []string{"mw1:before", "mw2:before", "handler", "mw2:after", "mw1:after"}
	if !equalSlices(order, want) {
		t.Errorf("expected order %v, got %v", want, order)
	}
}

func TestRouterUseDoesNotAffectAlreadyRegisteredRoutes(t *testing.T) {
	var order []string
	r := NewRouter()
	r.WithHandleFunc("/early", func(res http.ResponseWriter, req *http.Request) {
		order = append(order, "early")
	})
	r.WithMiddleware(recordingMiddleware("mw", &order))
	r.WithHandleFunc("/late", func(res http.ResponseWriter, req *http.Request) {
		order = append(order, "late")
	})

	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/early", nil))
	if got := []string{"early"}; !equalSlices(order, got) {
		t.Errorf("expected early route unaffected by later Use, got %v", order)
	}

	order = nil
	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/late", nil))
	if got := []string{"mw:before", "late", "mw:after"}; !equalSlices(order, got) {
		t.Errorf("expected late route to run middleware, got %v", order)
	}
}

func TestRouterWithAddsMiddlewareOnlyToDerivedRouter(t *testing.T) {
	var order []string
	r := NewRouter()
	r.WithMiddleware(recordingMiddleware("base", &order))
	r.WithHandleFunc("/base", func(res http.ResponseWriter, req *http.Request) {
		order = append(order, "base-handler")
	})
	r.WithGroup(recordingMiddleware("extra", &order)).WithHandleFunc("/extra", func(res http.ResponseWriter, req *http.Request) {
		order = append(order, "extra-handler")
	})

	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/base", nil))
	if got := []string{"base:before", "base-handler", "base:after"}; !equalSlices(order, got) {
		t.Errorf("expected base route unaffected by With, got %v", order)
	}

	order = nil
	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/extra", nil))
	want := []string{"base:before", "extra:before", "extra-handler", "extra:after", "base:after"}
	if !equalSlices(order, want) {
		t.Errorf("expected extra route to run base+extra middlewares, got %v", order)
	}
}

func TestRouterImplementsHTTPHandler(t *testing.T) {
	var _ http.Handler = NewRouter()
}
