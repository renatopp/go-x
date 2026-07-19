package logx

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/renatopp/go-x/timex"
)

// Options configures a Logger. Zero-value fields fall back to sane
// defaults: Level defaults to LevelInfo, Writer to os.Stderr, Formatter to
// TextFormatter, and TimeProvider to time.Now. EnableTime, EnableCaller,
// and EnableStack default to false (opt-in).
type Options struct {
	Level        Level
	Prefix       string
	Attrs        []Attr
	Writer       io.Writer
	Formatter    Formatter
	TimeProvider TimeProvider
	EnableTime   bool
	EnableCaller bool
	EnableStack  bool
	Sampler      *Sampler
	RateLimiter  timex.RateLimiter
}

// Logger is a slog.Handler with configurable level, attributes, prefix,
// timestamp/caller/stacktrace capture, output formatting, sampling, and
// rate limiting. It is safe for concurrent use.
type Logger struct {
	prefix       string
	attrs        []Attr
	writer       io.Writer
	formatter    Formatter
	timeProvider TimeProvider
	enableTime   bool
	enableCaller bool
	enableStack  bool
	sampler      *Sampler
	rateLimiter  timex.RateLimiter

	level  atomic.Int64
	groups []string
	mu     *sync.Mutex
}

// NewLogger creates a Logger from opts, filling in defaults for any
// unconfigured field.
func NewLogger(opts Options) *Logger {
	if opts.Writer == nil {
		opts.Writer = os.Stderr
	}
	if opts.Formatter == nil {
		opts.Formatter = TextFormatter
	}
	if opts.TimeProvider == nil {
		opts.TimeProvider = defaultTimeProvider
	}

	l := &Logger{
		prefix:       opts.Prefix,
		attrs:        append([]Attr{}, opts.Attrs...),
		writer:       opts.Writer,
		formatter:    opts.Formatter,
		timeProvider: opts.TimeProvider,
		enableTime:   opts.EnableTime,
		enableCaller: opts.EnableCaller,
		enableStack:  opts.EnableStack,
		sampler:      opts.Sampler,
		rateLimiter:  opts.RateLimiter,
		mu:           &sync.Mutex{},
	}
	l.level.Store(int64(opts.Level))
	return l
}

// Level returns the Logger's current minimum level.
func (l *Logger) Level() Level {
	return Level(l.level.Load())
}

// SetLevel changes the Logger's minimum level.
func (l *Logger) SetLevel(level Level) {
	l.level.Store(int64(level))
}

// With returns a Logger with args added as attributes, following the same
// key-value/Attr argument convention as slog.Logger.With.
func (l *Logger) With(args ...any) *Logger {
	attrs := argsToAttrs(args)
	if len(attrs) == 0 {
		return l
	}
	return l.WithAttrs(attrs).(*Logger)
}

// Group returns a Logger whose future attributes are nested under name.
func (l *Logger) Group(name string) *Logger {
	return l.WithGroup(name).(*Logger)
}

// Logger returns a *slog.Logger backed by this Logger, for callers that
// want the full log/slog API (e.g. LogAttrs, context-aware helpers).
func (l *Logger) Logger() *slog.Logger {
	return slog.New(l)
}

// Debug logs at LevelDebug.
func (l *Logger) Debug(msg string, args ...any) { l.log(context.Background(), LevelDebug, msg, args) }

// Info logs at LevelInfo.
func (l *Logger) Info(msg string, args ...any) { l.log(context.Background(), LevelInfo, msg, args) }

// Warn logs at LevelWarn.
func (l *Logger) Warn(msg string, args ...any) { l.log(context.Background(), LevelWarn, msg, args) }

// Error logs at LevelError.
func (l *Logger) Error(msg string, args ...any) { l.log(context.Background(), LevelError, msg, args) }

// Debugf logs at LevelDebug using fmt.Sprintf.
func (l *Logger) Debugf(format string, args ...any) {
	l.log(context.Background(), LevelDebug, fmt.Sprintf(format, args...), nil)
}

// Infof logs at LevelInfo using fmt.Sprintf.
func (l *Logger) Infof(format string, args ...any) {
	l.log(context.Background(), LevelInfo, fmt.Sprintf(format, args...), nil)
}

// Warnf logs at LevelWarn using fmt.Sprintf.
func (l *Logger) Warnf(format string, args ...any) {
	l.log(context.Background(), LevelWarn, fmt.Sprintf(format, args...), nil)
}

// Errorf logs at LevelError using fmt.Sprintf.
func (l *Logger) Errorf(format string, args ...any) {
	l.log(context.Background(), LevelError, fmt.Sprintf(format, args...), nil)
}

// Print logs at LevelInfo without any attributes.
func (l *Logger) Print(msg string) { l.log(context.Background(), LevelInfo, msg, nil) }

// PrintContext logs at LevelInfo with ctx without any attributes.
func (l *Logger) PrintContext(ctx context.Context, msg string) {
	l.log(ctx, LevelInfo, msg, nil)
}

// Printf logs at LevelInfo using fmt.Sprintf without any attributes.
func (l *Logger) Printf(format string, args ...any) {
	l.log(context.Background(), LevelInfo, fmt.Sprintf(format, args...), nil)
}

// DebugContext logs at LevelDebug with ctx.
func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelDebug, msg, args)
}

// InfoContext logs at LevelInfo with ctx.
func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelInfo, msg, args)
}

// WarnContext logs at LevelWarn with ctx.
func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelWarn, msg, args)
}

// ErrorContext logs at LevelError with ctx.
func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelError, msg, args)
}

// Enabled reports whether level is at or above the Logger's current level.
// It implements slog.Handler.
func (l *Logger) Enabled(_ context.Context, level slog.Level) bool {
	return level >= slog.Level(l.level.Load())
}

// Handle renders r and writes it to the Logger's Writer, applying sampling
// and rate limiting first. It implements slog.Handler.
func (l *Logger) Handle(_ context.Context, r slog.Record) error {
	if l.sampler != nil && !l.sampler.Allow() {
		return nil
	}
	if l.rateLimiter != nil && !l.rateLimiter.Consume(1) {
		return nil
	}

	e := &Entry{
		Level:   Level(r.Level),
		Prefix:  l.prefix,
		Message: r.Message,
	}
	if l.enableTime {
		e.Time = r.Time
		if e.Time.IsZero() {
			e.Time = l.timeProvider()
		}
	}
	if l.enableCaller && r.PC != 0 {
		e.Caller = callerString(r.PC)
	}
	if l.enableStack {
		e.Stack = string(debug.Stack())
	}

	recordAttrs := make([]Attr, 0, r.NumAttrs())
	r.Attrs(func(a Attr) bool {
		recordAttrs = append(recordAttrs, a)
		return true
	})
	e.Attrs = append(append([]Attr{}, l.attrs...), wrapGroups(l.groups, recordAttrs)...)

	line := append(l.formatter(e), '\n')

	l.mu.Lock()
	defer l.mu.Unlock()
	_, err := l.writer.Write(line)
	return err
}

// WithAttrs returns a Logger with attrs added, nested under any currently
// open groups. It implements slog.Handler.
func (l *Logger) WithAttrs(attrs []Attr) slog.Handler {
	if len(attrs) == 0 {
		return l
	}
	l2 := l.clone()
	l2.attrs = append(l2.attrs, wrapGroups(l.groups, attrs)...)
	return l2
}

// WithGroup returns a Logger whose future attributes are nested under name.
// It implements slog.Handler.
func (l *Logger) WithGroup(name string) slog.Handler {
	if name == "" {
		return l
	}
	l2 := l.clone()
	l2.groups = append(append([]string{}, l.groups...), name)
	return l2
}

// clone returns a copy of l that shares its Writer and mutex, so writes
// from derived Loggers (via WithAttrs/WithGroup) stay serialized.
func (l *Logger) clone() *Logger {
	l2 := &Logger{
		prefix:       l.prefix,
		attrs:        append([]Attr{}, l.attrs...),
		writer:       l.writer,
		formatter:    l.formatter,
		timeProvider: l.timeProvider,
		enableTime:   l.enableTime,
		enableCaller: l.enableCaller,
		enableStack:  l.enableStack,
		sampler:      l.sampler,
		rateLimiter:  l.rateLimiter,
		groups:       l.groups,
		mu:           l.mu,
	}
	l2.level.Store(l.level.Load())
	return l2
}

// log builds and handles a record for msg at level, capturing the caller
// site when EnableCaller is set.
func (l *Logger) log(ctx context.Context, level Level, msg string, args []any) {
	if !l.Enabled(ctx, level.slog()) {
		return
	}

	var pc uintptr
	if l.enableCaller {
		var pcs [1]uintptr
		runtime.Callers(3, pcs[:]) // skip [Callers, log, Debug/Info/...]
		pc = pcs[0]
	}

	r := slog.NewRecord(l.timeProvider(), level.slog(), msg, pc)
	r.AddAttrs(argsToAttrs(args)...)
	_ = l.Handle(ctx, r)
}

// callerString formats pc as "file:line".
func callerString(pc uintptr) string {
	if pc == 0 {
		return ""
	}
	frame, _ := runtime.CallersFrames([]uintptr{pc}).Next()
	if frame.File == "" {
		return ""
	}
	return frame.File + ":" + strconv.Itoa(frame.Line)
}

// wrapGroups nests attrs under groups, innermost group first, returning a
// single-element slice holding the outermost group Attr. If groups is
// empty, attrs is returned unchanged.
func wrapGroups(groups []string, attrs []Attr) []Attr {
	if len(groups) == 0 || len(attrs) == 0 {
		return attrs
	}
	args := make([]any, len(attrs))
	for i, a := range attrs {
		args[i] = a
	}
	wrapped := slog.Group(groups[len(groups)-1], args...)
	for i := len(groups) - 2; i >= 0; i-- {
		wrapped = slog.Group(groups[i], wrapped)
	}
	return []Attr{wrapped}
}

// argsToAttrs parses args into Attrs following slog's convention: Attr
// values are taken as-is, and other values are paired with the following
// argument as a key, with a "!BADKEY" fallback for unpaired trailing args.
func argsToAttrs(args []any) []Attr {
	var attrs []Attr
	for len(args) > 0 {
		switch a := args[0].(type) {
		case Attr:
			attrs = append(attrs, a)
			args = args[1:]
		case string:
			if len(args) == 1 {
				attrs = append(attrs, Any("!BADKEY", a))
				args = nil
			} else {
				attrs = append(attrs, Any(a, args[1]))
				args = args[2:]
			}
		default:
			attrs = append(attrs, Any("!BADKEY", a))
			args = args[1:]
		}
	}
	return attrs
}
