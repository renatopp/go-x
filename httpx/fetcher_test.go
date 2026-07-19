package httpx

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestFetcherRetriesOnServerError(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	f := NewFetcher(FetcherOptions{MaxRetries: 3, BackoffBase: time.Millisecond, BackoffMax: 5 * time.Millisecond})
	resp := f.Get(srv.URL)

	if resp.IsError() {
		t.Fatalf("unexpected error: %v", resp.Error())
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success, got status %d", resp.StatusCode())
	}
	if got := atomic.LoadInt32(&calls); got != 3 {
		t.Fatalf("expected 3 calls, got %d", got)
	}
}

func TestFetcherGivesUpAfterMaxRetries(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	f := NewFetcher(FetcherOptions{MaxRetries: 2, BackoffBase: time.Millisecond, BackoffMax: 5 * time.Millisecond})
	resp := f.Get(srv.URL)

	if resp.StatusCode() != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode())
	}
	if got := atomic.LoadInt32(&calls); got != 3 { // 1 initial + 2 retries
		t.Fatalf("expected 3 calls, got %d", got)
	}
}

func TestFetcherCircuitBreakerOpens(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	breaker := NewCircuitBreaker(2, time.Minute)
	f := NewFetcher(FetcherOptions{Breaker: breaker})

	f.Get(srv.URL) // failure 1
	f.Get(srv.URL) // failure 2, opens breaker

	resp := f.Get(srv.URL)
	if !resp.IsError() || resp.Error() != ErrCircuitOpen {
		t.Fatalf("expected ErrCircuitOpen, got status=%d err=%v", resp.StatusCode(), resp.Error())
	}
}

func TestFetcherRateLimiterThrottles(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	f := NewFetcher(FetcherOptions{RateLimiter: NewRateLimiter(10, 1)}) // 1 burst, 10/s => 100ms between extra tokens

	start := time.Now()
	f.Get(srv.URL)
	f.Get(srv.URL)
	elapsed := time.Since(start)

	if elapsed < 50*time.Millisecond {
		t.Fatalf("expected rate limiter to throttle second request, elapsed=%v", elapsed)
	}
}

func TestFetcherConcurrencyLimit(t *testing.T) {
	var inFlight, maxInFlight int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cur := atomic.AddInt32(&inFlight, 1)
		for {
			max := atomic.LoadInt32(&maxInFlight)
			if cur <= max || atomic.CompareAndSwapInt32(&maxInFlight, max, cur) {
				break
			}
		}
		time.Sleep(20 * time.Millisecond)
		atomic.AddInt32(&inFlight, -1)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	f := NewFetcher(FetcherOptions{Concurrency: 2})

	done := make(chan struct{})
	for range 5 {
		go func() {
			f.Get(srv.URL)
			done <- struct{}{}
		}()
	}
	for range 5 {
		<-done
	}

	if got := atomic.LoadInt32(&maxInFlight); got > 2 {
		t.Fatalf("expected at most 2 in-flight requests, got %d", got)
	}
}

func TestPackageLevelGetUsesDefaultFetcher(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Test", "ok")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	resp := Get(srv.URL)
	if resp.IsError() {
		t.Fatalf("unexpected error: %v", resp.Error())
	}
	if resp.Header("X-Test") != "ok" {
		t.Fatalf("expected header X-Test=ok, got %q", resp.Header("X-Test"))
	}
}

func TestFetcherPerRequestHeaderOverride(t *testing.T) {
	var got string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Get("X-Custom")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	f := NewFetcher(FetcherOptions{FetchOptions: FetchOptions{Headers: map[string]string{"X-Custom": "default"}}})
	f.Get(srv.URL, FetcherOptions{FetchOptions: FetchOptions{Headers: map[string]string{"X-Custom": "override"}}})

	if got != "override" {
		t.Fatalf("expected header override, got %q", got)
	}
}
