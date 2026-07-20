package logx

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"
	"time"
)

type OutputParts struct {
	Time    string
	Level   string
	Caller  string
	Message string
	Attrs   []string
}

type LevelStyle struct {
	LevelFormatter  func(level slog.Level) string
	TimeFormatter   func(time time.Time) string
	CallerFormatter func(source *Source) string
	MsgFormatter    func(msg string) string
	AttrFormatter   func(attr Attr, groups []string) string
	OutputFormatter func(parts OutputParts) string
}

type ColoredHandler struct {
	mu           sync.RWMutex
	writer       io.Writer
	levelStyles  map[Level]*LevelStyle
	defaultStyle *LevelStyle
	attrs        []slog.Attr
	groups       []string
}

func NewColoredHandler() *ColoredHandler {
	return &ColoredHandler{
		writer:      os.Stdout,
		levelStyles: make(map[Level]*LevelStyle),
		defaultStyle: &LevelStyle{
			LevelFormatter:  ColoredLevelFormatter,
			TimeFormatter:   PMTimeFormatter,
			CallerFormatter: CallerFormatter,
			MsgFormatter:    func(msg string) string { return msg },
			AttrFormatter:   AttrFormatter,
			OutputFormatter: CommonOutputFormatter,
		},
		attrs:  []slog.Attr{},
		groups: []string{},
	}
}

func (h *ColoredHandler) WithWriter(writer io.Writer) *ColoredHandler {
	h.mu.Lock()
	defer h.mu.Unlock()
	if writer == nil {
		h.writer = nil
	} else {
		h.writer = writer
	}
	return h
}

func (h *ColoredHandler) WithLevelStyle(level Level, style *LevelStyle) *ColoredHandler {
	h.mu.Lock()
	defer h.mu.Unlock()
	if style == nil {
		delete(h.levelStyles, level)
	} else {
		h.levelStyles[level] = style
	}
	return h
}

func (h *ColoredHandler) WithDefaultStyle(style *LevelStyle) *ColoredHandler {
	h.mu.Lock()
	defer h.mu.Unlock()
	if style == nil {
		h.defaultStyle = nil
	} else {
		h.defaultStyle = style
	}
	return h
}

func (h *ColoredHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

func (h *ColoredHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return &ColoredHandler{
		levelStyles:  h.levelStyles,
		defaultStyle: h.defaultStyle,
		attrs:        append(h.attrs, attrs...),
		groups:       h.groups,
	}
}

func (h *ColoredHandler) WithGroup(name string) slog.Handler {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return &ColoredHandler{
		levelStyles:  h.levelStyles,
		defaultStyle: h.defaultStyle,
		attrs:        h.attrs,
		groups:       append(h.groups, name),
	}
}

func (h *ColoredHandler) collectAttrs(record slog.Record, groups []string, formatter func(Attr, []string) string, result *[]string) {
	record.Attrs(func(a slog.Attr) bool {
		if a.Value.Kind() == slog.KindGroup {
			newGroups := append(groups, a.Key)
			h.walkAttrs(a.Value.Group(), newGroups, formatter, result)
		} else {
			*result = append(*result, formatter(a, groups))
		}
		return true
	})
}

func (h *ColoredHandler) walkAttrs(attrs []slog.Attr, groups []string, formatter func(Attr, []string) string, result *[]string) {
	for _, a := range attrs {
		if a.Value.Kind() == slog.KindGroup {
			newGroups := append(groups, a.Key)
			h.walkAttrs(a.Value.Group(), newGroups, formatter, result)
		} else {
			*result = append(*result, formatter(a, groups))
		}
	}
}

func (h *ColoredHandler) Handle(_ context.Context, record slog.Record) error {
	h.mu.RLock()
	style := h.defaultStyle
	if s, ok := h.levelStyles[Level(record.Level)]; ok {
		style = s
	}
	writer := h.writer
	groups := h.groups
	h.mu.RUnlock()

	if style == nil {
		return nil
	}

	timestamp := ""
	if style.TimeFormatter != nil {
		timestamp = style.TimeFormatter(record.Time)
	}

	level := ""
	if style.LevelFormatter != nil {
		level = style.LevelFormatter(record.Level)
	}

	caller := ""
	source := record.Source()
	if style.CallerFormatter != nil && source != nil {
		caller = style.CallerFormatter(source)
	}

	msg := ""
	if style.MsgFormatter != nil {
		msg = style.MsgFormatter(record.Message)
	}

	attrs := []string{}
	if style.AttrFormatter != nil {
		h.collectAttrs(record, groups, style.AttrFormatter, &attrs)
	}

	output := ""
	if style.OutputFormatter != nil {
		output = style.OutputFormatter(OutputParts{
			Time:    timestamp,
			Level:   level,
			Caller:  caller,
			Message: msg,
			Attrs:   attrs,
		})
	}

	fmt.Fprintln(writer, output)
	return nil
}
