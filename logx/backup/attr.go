package logx

import (
	"log/slog"
	"time"
)

// Attr is a key-value pair attached to a log record. It retypes slog.Attr
// so Logger can build on top of log/slog without exposing it directly.
type Attr = slog.Attr

// String returns an Attr for a string value.
func String(key, value string) Attr { return slog.String(key, value) }

// Int returns an Attr for an int value.
func Int(key string, value int) Attr { return slog.Int(key, value) }

// Int64 returns an Attr for an int64 value.
func Int64(key string, value int64) Attr { return slog.Int64(key, value) }

// Uint64 returns an Attr for a uint64 value.
func Uint64(key string, value uint64) Attr { return slog.Uint64(key, value) }

// Float64 returns an Attr for a float64 value.
func Float64(key string, value float64) Attr { return slog.Float64(key, value) }

// Bool returns an Attr for a bool value.
func Bool(key string, value bool) Attr { return slog.Bool(key, value) }

// Time returns an Attr for a time.Time value.
func Time(key string, value time.Time) Attr { return slog.Time(key, value) }

// Duration returns an Attr for a time.Duration value.
func Duration(key string, value time.Duration) Attr { return slog.Duration(key, value) }

// Any returns an Attr for a value of any type.
func Any(key string, value any) Attr { return slog.Any(key, value) }

// Group returns an Attr that groups the given Attrs under key.
func Group(key string, args ...any) Attr { return slog.Group(key, args...) }
