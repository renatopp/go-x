package logx

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/renatopp/go-x/fmtx"
)

// Handler is the simplest implementation of slog.Handler. It writes each
// record as a single plain text line to the underlying writer.
type Handler struct {
	w     io.Writer
	level slog.Level
}

// NewHandler creates a new Handler that writes to w, handling records at
// level and above.
func NewHandler(w io.Writer, level slog.Level) *Handler {
	return &Handler{w: w, level: level}
}

// Enabled reports whether the handler handles records at the given level.
func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

// Handle writes the record to the underlying writer.
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	line := fmtx.Sprint("%s %s", fmtx.Blue(fmtx.Sprint("%s", r.Level)), fmtx.BrightBlue(r.Message))
	r.Attrs(func(a slog.Attr) bool {
		line += fmtx.Dim(fmtx.Sprint("\n     %s: %v", a.Key, a.Value))
		return true
	})
	_, err := fmt.Fprintln(h.w, line)
	return err
}

// WithAttrs returns the handler unchanged, since it does not support
// persistent attributes.
func (h *Handler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

// WithGroup returns the handler unchanged, since it does not support
// groups.
func (h *Handler) WithGroup(_ string) slog.Handler {
	return h
}
