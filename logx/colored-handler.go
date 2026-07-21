package logx

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"sync"
	"time"
)

type ColoredHandler struct {
	mu                 sync.RWMutex
	writer             io.Writer
	levelFormatter     func(Level) string
	timeFormatter      func(time.Time, Level) string
	callerFormatter    func(*Source, Level) string
	messageFormatter   func(string, Level) string
	attributeFormatter func(Attr, []string, Level) string
	attrs              []slog.Attr
	groups             []string
}

func NewColoredHandler() *ColoredHandler {
	return &ColoredHandler{
		mu:                 sync.RWMutex{},
		writer:             os.Stdout,
		levelFormatter:     ColoredLevelFormatter,
		timeFormatter:      PMTimeFormatter,
		callerFormatter:    CallerFormatter,
		messageFormatter:   func(msg string, _ Level) string { return msg },
		attributeFormatter: AttrFormatter,
		attrs:              []slog.Attr{},
		groups:             []string{},
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

func (h *ColoredHandler) WithLevelFormatter(formatter func(Level) string) *ColoredHandler {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.levelFormatter = formatter
	return h
}

func (h *ColoredHandler) WithTimeFormatter(formatter func(time.Time, Level) string) *ColoredHandler {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.timeFormatter = formatter
	return h
}

func (h *ColoredHandler) WithCallerFormatter(formatter func(*Source, Level) string) *ColoredHandler {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.callerFormatter = formatter
	return h
}

func (h *ColoredHandler) WithMessageFormatter(formatter func(string, Level) string) *ColoredHandler {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.messageFormatter = formatter
	return h
}

func (h *ColoredHandler) WithAttributeFormatter(formatter func(Attr, []string, Level) string) *ColoredHandler {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.attributeFormatter = formatter
	return h
}

func (h *ColoredHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

func (h *ColoredHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return &ColoredHandler{
		writer:             h.writer,
		levelFormatter:     h.levelFormatter,
		timeFormatter:      h.timeFormatter,
		callerFormatter:    h.callerFormatter,
		messageFormatter:   h.messageFormatter,
		attributeFormatter: h.attributeFormatter,
		attrs:              append(h.attrs, attrs...),
		groups:             h.groups,
	}
}

func (h *ColoredHandler) WithGroup(name string) slog.Handler {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return &ColoredHandler{
		writer:             h.writer,
		levelFormatter:     h.levelFormatter,
		timeFormatter:      h.timeFormatter,
		callerFormatter:    h.callerFormatter,
		messageFormatter:   h.messageFormatter,
		attributeFormatter: h.attributeFormatter,
		attrs:              h.attrs,
		groups:             append(h.groups, name),
	}
}

func (h *ColoredHandler) collectAttrs(record Record, groups []string, result *[]string) {
	record.Attrs(func(a slog.Attr) bool {
		if a.Value.Kind() == slog.KindGroup {
			newGroups := append(groups, a.Key)
			h.walkAttrs(record, a.Value.Group(), newGroups, result)
		} else {
			*result = append(*result, h.attributeFormatter(a, groups, record.Level))
		}
		return true
	})
}

func (h *ColoredHandler) walkAttrs(record Record, attrs []Attr, groups []string, result *[]string) {
	for _, a := range attrs {
		if a.Value.Kind() == slog.KindGroup {
			newGroups := append(groups, a.Key)
			h.walkAttrs(record, a.Value.Group(), newGroups, result)
		} else {
			*result = append(*result, h.attributeFormatter(a, groups, record.Level))
		}
	}
}

func (h *ColoredHandler) Handle(_ context.Context, record slog.Record) error {
	h.mu.RLock()
	writer := h.writer
	groups := h.groups
	h.mu.RUnlock()

	timestamp := ""
	if h.timeFormatter != nil && !record.Time.IsZero() {
		timestamp = h.timeFormatter(record.Time, record.Level)
	}

	level := ""
	if h.levelFormatter != nil {
		level = h.levelFormatter(record.Level)
	}

	caller := ""
	source := record.Source()
	if h.callerFormatter != nil && source != nil {
		caller = h.callerFormatter(source, record.Level)
	}

	msg := ""
	if h.messageFormatter != nil {
		msg = h.messageFormatter(record.Message, record.Level)
	}

	attrs := []string{}
	if h.attributeFormatter != nil {
		h.collectAttrs(record, groups, &attrs)
	}

	builder := strings.Builder{}
	if timestamp != "" {
		builder.WriteString(timestamp)
		builder.WriteString(" ")
	}

	if level != "" {
		builder.WriteString(level)
		builder.WriteString(" ")
	}

	if caller != "" {
		builder.WriteString(caller)
		builder.WriteString(" ")
	}

	if msg != "" {
		builder.WriteString(msg)
		builder.WriteString(" ")
	}

	for _, attr := range attrs {
		builder.WriteString(attr)
		builder.WriteString(" ")
	}

	fmt.Fprintln(writer, builder.String())
	return nil
}
