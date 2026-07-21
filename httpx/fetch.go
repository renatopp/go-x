package httpx

// import (
// 	"bytes"
// 	"context"
// 	"maps"
// 	"net/http"
// 	"time"
// )

// type FetchOptions struct {
// 	Headers       map[string]string
// 	Timeout       time.Duration
// 	Context       context.Context
// 	HttpClient    *http.Client
// 	Jar           http.CookieJar
// 	CheckRedirect func(req *http.Request, via []*http.Request) error
// }

// // Fetch performs an HTTP request with the specified method, URL, and options.
// // This always returns a Response object, even if there was an error during
// // the request. You can check for errors using the IsError() method on the
// // Response.
// //
// // The response body is loaded on-demand when you call Bytes, Text, or Json.
// //
// // By default, Fetch uses a timeout of 60 seconds. You can override this by
// // providing a FetchOptions with a different Timeout value. Setting Timeout to
// // negative means no timeout.
// func Fetch(method, url string, opts ...FetchOptions) *FetchResponse {
// 	return FetchWithBody(method, url, nil, opts...)
// }

// // FetchWithBody performs an HTTP request with the specified method, URL, body, and options.
// // This always returns a Response object, even if there was an error during
// // the request. You can check for errors using the IsError() method on the
// // Response.
// //
// // The response body is loaded on-demand when you call Bytes, Text, or Json.
// //
// // By default, FetchWithBody uses a timeout of 60 seconds. You can override this by
// // providing a RequestOption with a different Timeout value. Setting Timeout to
// // negative means no timeout.
// func FetchWithBody(method, url string, body []byte, opts ...FetchOptions) *FetchResponse {
// 	options := FetchOptions{
// 		Headers: make(map[string]string),
// 		Timeout: 60 * time.Second, // Default timeout
// 	}

// 	for _, opt := range opts {
// 		if opt.Headers != nil {
// 			maps.Copy(options.Headers, opt.Headers)
// 		}
// 		if opt.Timeout != 0 {
// 			options.Timeout = max(0, opt.Timeout)
// 		}
// 		if opt.Context != nil {
// 			options.Context = opt.Context
// 		}
// 		if opt.HttpClient != nil {
// 			options.HttpClient = opt.HttpClient
// 		}
// 		if opt.Jar != nil {
// 			options.Jar = opt.Jar
// 		}
// 		if opt.CheckRedirect != nil {
// 			options.CheckRedirect = opt.CheckRedirect
// 		}
// 	}

// 	ctx := options.Context
// 	if ctx == nil {
// 		ctx = context.Background()
// 	}
// 	if options.Timeout > 0 {
// 		var cancel context.CancelFunc
// 		ctx, cancel = context.WithTimeout(ctx, options.Timeout)
// 		defer cancel()
// 	}

// 	client := options.HttpClient
// 	if client == nil {
// 		client = &http.Client{Transport: transport, Jar: options.Jar, CheckRedirect: options.CheckRedirect}
// 	}

// 	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
// 	if err != nil {
// 		return newResponse(req, nil, err)
// 	}

// 	for k, v := range options.Headers {
// 		req.Header.Set(k, v)
// 	}

// 	response, err := client.Do(req)
// 	return newResponse(req, response, err)
// }

// // Get is a shortcut for performing a GET using DefaultFetcher.
// func Get(url string, opts ...FetchOptions) *FetchResponse {
// 	return DefaultFetcher.Get(url, wrapFetchOptions(opts)...)
// }

// // Post is a shortcut for performing a POST using DefaultFetcher.
// func Post(url string, body []byte, opts ...FetchOptions) *FetchResponse {
// 	return DefaultFetcher.Post(url, body, wrapFetchOptions(opts)...)
// }

// // Put is a shortcut for performing a PUT using DefaultFetcher.
// func Put(url string, body []byte, opts ...FetchOptions) *FetchResponse {
// 	return DefaultFetcher.Put(url, body, wrapFetchOptions(opts)...)
// }

// // Delete is a shortcut for performing a DELETE using DefaultFetcher.
// func Delete(url string, opts ...FetchOptions) *FetchResponse {
// 	return DefaultFetcher.Delete(url, wrapFetchOptions(opts)...)
// }

// // Patch is a shortcut for performing a PATCH using DefaultFetcher.
// func Patch(url string, body []byte, opts ...FetchOptions) *FetchResponse {
// 	return DefaultFetcher.Patch(url, body, wrapFetchOptions(opts)...)
// }

// // Head is a shortcut for performing a HEAD using DefaultFetcher.
// func Head(url string, opts ...FetchOptions) *FetchResponse {
// 	return DefaultFetcher.Head(url, wrapFetchOptions(opts)...)
// }

// // Options is a shortcut for performing an OPTIONS using DefaultFetcher.
// func Options(url string, opts ...FetchOptions) *FetchResponse {
// 	return DefaultFetcher.Options(url, wrapFetchOptions(opts)...)
// }

// // Trace is a shortcut for performing a TRACE using DefaultFetcher.
// func Trace(url string, opts ...FetchOptions) *FetchResponse {
// 	return DefaultFetcher.Trace(url, wrapFetchOptions(opts)...)
// }

// // Connect is a shortcut for performing a CONNECT using DefaultFetcher.
// func Connect(url string, opts ...FetchOptions) *FetchResponse {
// 	return DefaultFetcher.Connect(url, wrapFetchOptions(opts)...)
// }

// // SSE is a shortcut for performing a GET with the Accept header set to
// // text/event-stream using DefaultFetcher.
// func SSE(url string, opts ...FetchOptions) *FetchResponse {
// 	opts = append(opts, FetchOptions{
// 		Headers: map[string]string{
// 			"Accept": "text/event-stream",
// 		},
// 	})
// 	return Get(url, opts...)
// }

// // wrapFetchOptions converts plain FetchOptions into FetcherOptions so they
// // can be passed through to a Fetcher.
// func wrapFetchOptions(opts []FetchOptions) []FetcherOptions {
// 	wrapped := make([]FetcherOptions, len(opts))
// 	for i, opt := range opts {
// 		wrapped[i] = FetcherOptions{FetchOptions: opt}
// 	}
// 	return wrapped
// }
