package httpx

import (
	"bytes"
	"maps"
	"net/http"
	"time"
)

type FetchOptions struct {
	Headers map[string]string
	Timeout time.Duration
}

// Fetch performs an HTTP request with the specified method, URL, and options.
// This always returns a Response object, even if there was an error during
// the request. You can check for errors using the IsError() method on the
// Response.
//
// The response body is loaded on-demand when you call Bytes, Text, or Json.
//
// By default, Fetch uses a timeout of 60 seconds. You can override this by
// providing a FetchOptions with a different Timeout value. Setting Timeout to
// negative means no timeout.
func Fetch(method, url string, opts ...FetchOptions) *FetchResponse {
	return FetchWithBody(method, url, nil, opts...)
}

// FetchWithBody performs an HTTP request with the specified method, URL, body, and options.
// This always returns a Response object, even if there was an error during
// the request. You can check for errors using the IsError() method on the
// Response.
//
// The response body is loaded on-demand when you call Bytes, Text, or Json.
//
// By default, FetchWithBody uses a timeout of 60 seconds. You can override this by
// providing a RequestOption with a different Timeout value. Setting Timeout to
// negative means no timeout.
func FetchWithBody(method, url string, body []byte, opts ...FetchOptions) *FetchResponse {
	options := FetchOptions{
		Headers: make(map[string]string),
		Timeout: 60 * time.Second, // Default timeout
	}

	for _, opt := range opts {
		if opt.Headers != nil {
			maps.Copy(options.Headers, opt.Headers)
		}
		if opt.Timeout != 0 {
			options.Timeout = max(0, opt.Timeout)
		}
	}

	client := &http.Client{Timeout: options.Timeout}
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return newResponse(req, nil, err)
	}

	for k, v := range options.Headers {
		req.Header.Set(k, v)
	}

	response, err := client.Do(req)
	return newResponse(req, response, err)
}

// Get is a shortcut for performing a GET using Fetch.
func Get(url string, opts ...FetchOptions) *FetchResponse {
	return Fetch(http.MethodGet, url, opts...)
}

// Post is a shortcut for performing a POST using FetchWithBody.
func Post(url string, body []byte, opts ...FetchOptions) *FetchResponse {
	return FetchWithBody(http.MethodPost, url, body, opts...)
}

// Put is a shortcut for performing a PUT using FetchWithBody.
func Put(url string, body []byte, opts ...FetchOptions) *FetchResponse {
	return FetchWithBody(http.MethodPut, url, body, opts...)
}

// Delete is a shortcut for performing a DELETE using Fetch.
func Delete(url string, opts ...FetchOptions) *FetchResponse {
	return Fetch(http.MethodDelete, url, opts...)
}

// Patch is a shortcut for performing a PATCH using FetchWithBody.
func Patch(url string, body []byte, opts ...FetchOptions) *FetchResponse {
	return FetchWithBody(http.MethodPatch, url, body, opts...)
}

// Head is a shortcut for performing a HEAD using Fetch.
func Head(url string, opts ...FetchOptions) *FetchResponse {
	return Fetch(http.MethodHead, url, opts...)
}

// Options is a shortcut for performing an OPTIONS using Fetch.
func Options(url string, opts ...FetchOptions) *FetchResponse {
	return Fetch(http.MethodOptions, url, opts...)
}

// Trace is a shortcut for performing a TRACE using Fetch.
func Trace(url string, opts ...FetchOptions) *FetchResponse {
	return Fetch(http.MethodTrace, url, opts...)
}

// Connect is a shortcut for performing a CONNECT using Fetch.
func Connect(url string, opts ...FetchOptions) *FetchResponse {
	return Fetch(http.MethodConnect, url, opts...)
}

func SSE(url string, opts ...FetchOptions) *FetchResponse {
	opts = append(opts, FetchOptions{
		Headers: map[string]string{
			"Accept": "text/event-stream",
		},
	})
	return Fetch(http.MethodGet, url, opts...)
}
