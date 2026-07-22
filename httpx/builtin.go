package httpx

import "net/http"

// Re-export commonly used net/http types for convenience.
type (
	Client         = http.Client
	Request        = http.Request
	Response       = http.Response
	ResponseWriter = http.ResponseWriter
	Handler        = http.Handler
	HandlerFunc    = http.HandlerFunc
	Header         = http.Header
	Cookie         = http.Cookie
	CookieJar      = http.CookieJar
	Transport      = http.Transport
	ServeMux       = http.ServeMux
	Server         = http.Server
	RoundTripper   = http.RoundTripper
	Dir            = http.Dir
	File           = http.File
	FileSystem     = http.FileSystem
)

// Re-export commonly used net/http functions for convenience.
var (
	NewRequest            = http.NewRequest
	NewRequestWithContext = http.NewRequestWithContext
	Get                   = http.Get
	Head                  = http.Head
	Post                  = http.Post
	PostForm              = http.PostForm
	NewServeMux           = http.NewServeMux
	ListenAndServe        = http.ListenAndServe
	ListenAndServeTLS     = http.ListenAndServeTLS
	Serve                 = http.Serve
	Error                 = http.Error
	NotFound              = http.NotFound
	Redirect              = http.Redirect
	SetCookie             = http.SetCookie
	FileServer            = http.FileServer
	StripPrefix           = http.StripPrefix
	ServeFile             = http.ServeFile
	DefaultClient         = http.DefaultClient
	DefaultServeMux       = http.DefaultServeMux
)

// Re-export commonly used net/http status code constants.
const (
	StatusOK                  = http.StatusOK
	StatusCreated             = http.StatusCreated
	StatusNoContent           = http.StatusNoContent
	StatusMovedPermanently    = http.StatusMovedPermanently
	StatusFound               = http.StatusFound
	StatusNotModified         = http.StatusNotModified
	StatusBadRequest          = http.StatusBadRequest
	StatusUnauthorized        = http.StatusUnauthorized
	StatusForbidden           = http.StatusForbidden
	StatusNotFound            = http.StatusNotFound
	StatusMethodNotAllowed    = http.StatusMethodNotAllowed
	StatusInternalServerError = http.StatusInternalServerError
	StatusBadGateway          = http.StatusBadGateway
	StatusServiceUnavailable  = http.StatusServiceUnavailable
)
