package logx

import "log/slog"

// Level is the severity of a log record. It retypes slog.Level so it can
// carry package-specific methods while staying numerically compatible with
// the standard library.
type Level int

const (
	LevelDebug Level = Level(slog.LevelDebug)
	LevelInfo  Level = Level(slog.LevelInfo)
	LevelWarn  Level = Level(slog.LevelWarn)
	LevelError Level = Level(slog.LevelError)
)

// DefaultLevel is the level used by a Logger when none is provided.
const DefaultLevel = LevelInfo

// ParseLevel parses a level name such as "debug", "INFO", or "warn+2" into
// a Level.
func ParseLevel(s string) (Level, error) {
	var l slog.Level
	if err := l.UnmarshalText([]byte(s)); err != nil {
		return 0, err
	}
	return Level(l), nil
}

// ForceParseLevel is like ParseLevel but ignores the error, returning the
// zero Level (LevelInfo) if s cannot be parsed.
func ForceParseLevel(s string) Level {
	l, _ := ParseLevel(s)
	return l
}

// String returns the level's name, e.g. "INFO" or "WARN+2".
func (l Level) String() string {
	return l.slog().String()
}

// slog converts l to its slog.Level equivalent.
func (l Level) slog() slog.Level {
	return slog.Level(l)
}
