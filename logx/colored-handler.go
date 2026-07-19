package logx

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
)

type LevelStyle struct {
	LevelFormatter  func(level slog.Level) string
	TimeFormatter   func(time time.Time) string
	CallerFormatter func(source *Source) string
	MsgFormatter    func(msg string) string
	AttrFormatter   func(attr slog.Attr) string
	OutputFormatter func(time, level, caller, msg string, attrs []string) string
}

type ColoredHandler struct {
	writer       io.Writer
	levelStyles  map[Level]*LevelStyle
	defaultStyle *LevelStyle
	attrs        []slog.Attr
	groups       []string
}

func NewColoredHandler() *ColoredHandler {
	return &ColoredHandler{
		writer:       os.Stdout,
		levelStyles:  make(map[Level]*LevelStyle),
		defaultStyle: &LevelStyle{},
		attrs:        []slog.Attr{},
		groups:       []string{},
	}
}

func (h *ColoredHandler) WithWriter(writer io.Writer) *ColoredHandler {
	if writer == nil {
		h.writer = nil
	} else {
		h.writer = writer
	}
	return h
}

func (h *ColoredHandler) WithLevelStyle(level Level, style *LevelStyle) *ColoredHandler {
	if style == nil {
		delete(h.levelStyles, level)
	} else {
		h.levelStyles[level] = style
	}
	return h
}

func (h *ColoredHandler) WithDefaultStyle(style *LevelStyle) *ColoredHandler {
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
	return &ColoredHandler{
		levelStyles:  h.levelStyles,
		defaultStyle: h.defaultStyle,
		attrs:        append(h.attrs, attrs...),
		groups:       h.groups,
	}
}

func (h *ColoredHandler) WithGroup(name string) slog.Handler {
	return &ColoredHandler{
		levelStyles:  h.levelStyles,
		defaultStyle: h.defaultStyle,
		attrs:        h.attrs,
		groups:       append(h.groups, name),
	}
}

func (h *ColoredHandler) Handle(_ context.Context, record slog.Record) error {
	style := h.defaultStyle
	if s, ok := h.levelStyles[Level(record.Level)]; ok {
		style = s
	}

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
	if style.CallerFormatter != nil {
		caller = style.CallerFormatter(record.Source())
	}

	msg := ""
	if style.MsgFormatter != nil {
		msg = style.MsgFormatter(record.Message)
	}

	attrs := []string{}
	if style.AttrFormatter != nil {
		record.Attrs(func(a slog.Attr) bool {
			attrs = append(attrs, style.AttrFormatter(a))
			return true
		})
	}

	output := ""
	if style.OutputFormatter != nil {
		output = style.OutputFormatter(timestamp, level, caller, msg, attrs)
	}

	fmt.Fprintln(h.writer, output)
	return nil
}
