package logx

import (
	"io"
	"sync"
)

type Level int

type Logger struct {
	mu     sync.Mutex
	level  Level
	prefix string
	writer io.Writer
	// attrs []Attr
	// formatter    Formatter
	// timeProvider TimeProvider
	// rateLimiter  timex.RateLimiter
}
