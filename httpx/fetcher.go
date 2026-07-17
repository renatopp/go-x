package httpx

import (
	"net/http"
	"time"
)

var transport = &http.Transport{
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 20,
	IdleConnTimeout:     90 * time.Second,
}

var defaultClient = &http.Client{
	Transport: transport,
}

var defaultTimeout = 60 * time.Second
var defaultFetcher = NewFetcher()

func SetDefaultHttpClient(client *http.Client) {
	defaultClient = client
}

func SetDefaultTimeout(timeout time.Duration) {
	defaultTimeout = timeout
}

func SetDefaultFetcher(fetcher *Fetcher) {
	defaultFetcher = fetcher
}

// Fetcher represents a http request
type Fetcher struct {
	Headers    map[string]string
	HttpClient *http.Client
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		Headers: make(map[string]string),
		HttpClient: &http.Client{
			Timeout:   60 * time.Second,
			Transport: transport,
		},
	}
}
