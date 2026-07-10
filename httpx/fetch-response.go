package httpx

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

// FetchResponse represents the result of an HTTP request. It contains the response body,
// status code, headers, and any error that occurred during the request.
type FetchResponse struct {
	err              error
	body             []byte
	bodyLoadErr      error
	bodyLoadOnce     sync.Once
	httpRequest      *http.Request
	httpResponse     *http.Response
	status           string // e.g. "200 OK"
	statusCode       int    // e.g. 200
	protocol         string // e.g. "HTTP/1.0"
	protocolMajor    int    // e.g. 1
	protocolMinor    int    // e.g. 0
	headers          http.Header
	contentLength    int64
	transferEncoding []string
	tls              *tls.ConnectionState
}

// newResponse creates a new Response object from the given http.Request,
// http.Response, and error. Internal use only.
func newResponse(req *http.Request, resp *http.Response, err error) *FetchResponse {
	if err != nil {
		return &FetchResponse{err: err, body: []byte{}}
	}

	return &FetchResponse{
		httpRequest:      req,
		httpResponse:     resp,
		status:           resp.Status,
		statusCode:       resp.StatusCode,
		protocol:         resp.Proto,
		protocolMajor:    resp.ProtoMajor,
		protocolMinor:    resp.ProtoMinor,
		headers:          resp.Header,
		contentLength:    resp.ContentLength,
		transferEncoding: resp.TransferEncoding,
		tls:              resp.TLS,
		body:             []byte{},
	}
}

func (r *FetchResponse) HttpRequest() *http.Request   { return r.httpRequest }
func (r *FetchResponse) HttpResponse() *http.Response { return r.httpResponse }
func (r *FetchResponse) Status() string               { return r.status }
func (r *FetchResponse) StatusCode() int              { return r.statusCode }
func (r *FetchResponse) Protocol() string             { return r.protocol }
func (r *FetchResponse) ProtocolMajor() int           { return r.protocolMajor }
func (r *FetchResponse) ProtocolMinor() int           { return r.protocolMinor }
func (r *FetchResponse) Headers() http.Header         { return r.headers }
func (r *FetchResponse) ContentLength() int64         { return r.contentLength }
func (r *FetchResponse) TransferEncoding() []string   { return r.transferEncoding }
func (r *FetchResponse) TLS() *tls.ConnectionState    { return r.tls }

// IsError returns true if there was an error during the HTTP request. You can
// get the error using the Error() method. Notice that a non-2xx status code
// does not count as an error.
func (r *FetchResponse) IsError() bool { return r.err != nil }

// Error returns the error that occurred during the HTTP request, or nil if
// there was no error. Notice that a non-2xx status code does not count as an
// error.
func (r *FetchResponse) Error() error { return r.err }

// Cookies returns the cookies sent by the server in the response. If there was
// an error during the request, it returns an empty slice.
func (r *FetchResponse) Cookies() []*http.Cookie {
	if r.IsError() {
		return []*http.Cookie{}
	}
	return r.httpResponse.Cookies()
}

// Header returns the value of the specified header key. If the header is not
// present, it returns an empty string. If there was an error during the request,
// it also returns an empty string.
func (r *FetchResponse) Header(key string) string {
	return r.headers.Get(key)
}

// Stream returns the response body as an io.ReadCloser. If there was an error during the request,
// it returns nil.
func (r *FetchResponse) Stream() io.ReadCloser {
	return r.httpResponse.Body
}

func (r *FetchResponse) Scanner() *bufio.Scanner {
	return bufio.NewScanner(r.httpResponse.Body)
}

// Text returns the response body as a string. If there was an error during the request,
// it returns an empty string.
func (r *FetchResponse) Text() string { return string(r.Bytes()) }

// Bytes returns the response body as a byte slice. If there was an error during the request,
// it returns an empty byte slice.
func (r *FetchResponse) Bytes() []byte {
	body, _ := r.loadBody()
	return body
}

// Json unmarshals the response body into the provided variable.
func (r *FetchResponse) Json(v any) error {
	body, err := r.loadBody()
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

// IsSuccess returns true if the HTTP status code is in the 2xx range, indicating a successful response.
func (r *FetchResponse) IsSuccess() bool { return r.statusCode >= 200 && r.statusCode < 300 }

// Is2xx is an alias for IsSuccess, returning true if the HTTP status code is in the 2xx range.
func (r *FetchResponse) Is2xx() bool { return r.IsSuccess() }

// Is3xx returns true if the HTTP status code is in the 3xx range, indicating a redirection response.
func (r *FetchResponse) Is3xx() bool { return r.statusCode >= 300 && r.statusCode < 400 }

// Is4xx returns true if the HTTP status code is in the 4xx range, indicating a client error response.
func (r *FetchResponse) Is4xx() bool { return r.statusCode >= 400 && r.statusCode < 500 }

// Is5xx returns true if the HTTP status code is in the 5xx range, indicating a server error response.
func (r *FetchResponse) Is5xx() bool { return r.statusCode >= 500 && r.statusCode < 600 }

func (r *FetchResponse) loadBody() ([]byte, error) {
	r.bodyLoadOnce.Do(func() {
		if r.httpResponse == nil || r.httpResponse.Body == nil {
			return
		}

		defer r.httpResponse.Body.Close()
		body, err := io.ReadAll(r.httpResponse.Body)
		if err != nil {
			r.bodyLoadErr = err
			r.body = []byte{}
			return
		}

		r.body = body
	})

	return r.body, r.bodyLoadErr
}
