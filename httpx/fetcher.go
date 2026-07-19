package httpx

import (
	"context"
	"errors"
	"maps"
	"math/rand"
	"net/http"
	"time"

	"github.com/renatopp/go-x/timex"
)

var transport = &http.Transport{
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 20,
	IdleConnTimeout:     90 * time.Second,
}

var DefaultClient = &http.Client{
	Transport: transport,
}
var DefaultFetcher = NewFetcher()

// ErrCircuitOpen is returned when a request is rejected because the
// Fetcher's circuit breaker is open.
var ErrCircuitOpen = errors.New("httpx: circuit breaker is open")

// FetcherOptions configures a Fetcher, either as its defaults (NewFetcher)
// or as a per-request override that complements or replaces those defaults.
type FetcherOptions struct {
	FetchOptions

	MaxRetries  int
	BackoffBase time.Duration
	BackoffMax  time.Duration
	Jitter      bool

	// Concurrency limits the number of in-flight requests for a Fetcher. It
	// is only applied when constructing a Fetcher (NewFetcher); per-request
	// overrides are ignored since the limit is shared, coordinated state.
	Concurrency int

	RateLimiter timex.RateLimiter
	Breaker     *CircuitBreaker
}

// Fetcher is a configurable, reusable HTTP client with retries, exponential
// backoff, concurrency control, rate limiting, and circuit breaking on top
// of its default headers, timeout, and client behavior.
type Fetcher struct {
	FetcherOptions

	sem chan struct{}
}

// NewFetcher creates a Fetcher with sane defaults, customized by opts.
func NewFetcher(opts ...FetcherOptions) *Fetcher {
	options := FetcherOptions{
		FetchOptions: FetchOptions{
			Headers:    make(map[string]string),
			Timeout:    60 * time.Second,
			HttpClient: DefaultClient,
		},
		BackoffBase: 100 * time.Millisecond,
		BackoffMax:  10 * time.Second,
	}

	for _, opt := range opts {
		options = options.merge(opt)
	}

	f := &Fetcher{FetcherOptions: options}
	if options.Concurrency > 0 {
		f.sem = make(chan struct{}, options.Concurrency)
	}
	return f
}

// Do performs an HTTP request applying the Fetcher's defaults merged with
// any per-request opts, including retries, rate limiting, concurrency
// control, and circuit breaking. It always returns a Response, even if
// there was an error - check IsError()/Error() on the result.
func (f *Fetcher) Do(method, url string, body []byte, opts ...FetcherOptions) *FetchResponse {
	cfg := f.FetcherOptions
	for _, opt := range opts {
		cfg = cfg.merge(opt)
	}

	if cfg.Breaker != nil && !cfg.Breaker.Allow() {
		return newResponse(nil, nil, ErrCircuitOpen)
	}

	if cfg.RateLimiter != nil {
		ctx := cfg.Context
		if ctx == nil {
			ctx = context.Background()
		}
		if err := cfg.RateLimiter.Wait(ctx); err != nil {
			return newResponse(nil, nil, err)
		}
	}

	if f.sem != nil {
		f.sem <- struct{}{}
		defer func() { <-f.sem }()
	}

	var resp *FetchResponse
	for attempt := 0; ; attempt++ {
		resp = FetchWithBody(method, url, body, cfg.FetchOptions)
		ok := !resp.IsError() && resp.StatusCode() < 500

		if cfg.Breaker != nil {
			if ok {
				cfg.Breaker.Success()
			} else {
				cfg.Breaker.Failure()
			}
		}

		if ok || attempt >= cfg.MaxRetries {
			return resp
		}
		time.Sleep(cfg.backoff(attempt))
	}
}

// Fetch is a shortcut for performing a request with no body using Do.
func (f *Fetcher) Fetch(method, url string, opts ...FetcherOptions) *FetchResponse {
	return f.Do(method, url, nil, opts...)
}

// Get is a shortcut for performing a GET using Do.
func (f *Fetcher) Get(url string, opts ...FetcherOptions) *FetchResponse {
	return f.Do(http.MethodGet, url, nil, opts...)
}

// Post is a shortcut for performing a POST using Do.
func (f *Fetcher) Post(url string, body []byte, opts ...FetcherOptions) *FetchResponse {
	return f.Do(http.MethodPost, url, body, opts...)
}

// Put is a shortcut for performing a PUT using Do.
func (f *Fetcher) Put(url string, body []byte, opts ...FetcherOptions) *FetchResponse {
	return f.Do(http.MethodPut, url, body, opts...)
}

// Delete is a shortcut for performing a DELETE using Do.
func (f *Fetcher) Delete(url string, opts ...FetcherOptions) *FetchResponse {
	return f.Do(http.MethodDelete, url, nil, opts...)
}

// Patch is a shortcut for performing a PATCH using Do.
func (f *Fetcher) Patch(url string, body []byte, opts ...FetcherOptions) *FetchResponse {
	return f.Do(http.MethodPatch, url, body, opts...)
}

// Head is a shortcut for performing a HEAD using Do.
func (f *Fetcher) Head(url string, opts ...FetcherOptions) *FetchResponse {
	return f.Do(http.MethodHead, url, nil, opts...)
}

// Options is a shortcut for performing an OPTIONS using Do.
func (f *Fetcher) Options(url string, opts ...FetcherOptions) *FetchResponse {
	return f.Do(http.MethodOptions, url, nil, opts...)
}

// Trace is a shortcut for performing a TRACE using Do.
func (f *Fetcher) Trace(url string, opts ...FetcherOptions) *FetchResponse {
	return f.Do(http.MethodTrace, url, nil, opts...)
}

// Connect is a shortcut for performing a CONNECT using Do.
func (f *Fetcher) Connect(url string, opts ...FetcherOptions) *FetchResponse {
	return f.Do(http.MethodConnect, url, nil, opts...)
}

// merge returns a copy of o with the non-zero/non-nil fields of over applied
// on top. Headers are merged key by key rather than replaced wholesale.
func (o FetcherOptions) merge(over FetcherOptions) FetcherOptions {
	if over.Headers != nil {
		headers := make(map[string]string, len(o.Headers)+len(over.Headers))
		maps.Copy(headers, o.Headers)
		maps.Copy(headers, over.Headers)
		o.Headers = headers
	}
	if over.Timeout != 0 {
		o.Timeout = over.Timeout
	}
	if over.Context != nil {
		o.Context = over.Context
	}
	if over.HttpClient != nil {
		o.HttpClient = over.HttpClient
	}
	if over.Jar != nil {
		o.Jar = over.Jar
	}
	if over.CheckRedirect != nil {
		o.CheckRedirect = over.CheckRedirect
	}
	if over.MaxRetries != 0 {
		o.MaxRetries = over.MaxRetries
	}
	if over.BackoffBase != 0 {
		o.BackoffBase = over.BackoffBase
	}
	if over.BackoffMax != 0 {
		o.BackoffMax = over.BackoffMax
	}
	if over.Jitter {
		o.Jitter = true
	}
	if over.Concurrency != 0 {
		o.Concurrency = over.Concurrency
	}
	if over.RateLimiter != nil {
		o.RateLimiter = over.RateLimiter
	}
	if over.Breaker != nil {
		o.Breaker = over.Breaker
	}
	return o
}

// backoff computes the delay before the next retry attempt: exponential
// growth capped at BackoffMax, with optional full jitter.
func (o FetcherOptions) backoff(attempt int) time.Duration {
	d := o.BackoffBase << attempt
	if d <= 0 || d > o.BackoffMax {
		d = o.BackoffMax
	}
	if o.Jitter && d > 0 {
		d = time.Duration(rand.Int63n(int64(d)))
	}
	return d
}
