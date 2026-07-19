package logx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"time"
)

// Entry is the rendered data for a single log record, passed to a Formatter.
type Entry struct {
	Time    time.Time
	Level   Level
	Prefix  string
	Message string
	Attrs   []Attr
	Caller  string
	Stack   string
}

// Formatter renders an Entry into the bytes written to a Logger's output.
// The returned bytes should not include a trailing newline; the Logger adds
// one.
type Formatter func(e *Entry) []byte

const (
	ansiReset  = "\033[0m"
	ansiGray   = "\033[90m"
	ansiBlue   = "\033[34m"
	ansiYellow = "\033[33m"
	ansiRed    = "\033[31m"
)

// JSONFormatter renders an Entry as a single-line JSON object.
func JSONFormatter(e *Entry) []byte {
	var buf bytes.Buffer
	first := true

	field := func(key string, val any) {
		if !first {
			buf.WriteByte(',')
		}
		first = false
		writeJSONString(&buf, key)
		buf.WriteByte(':')
		b, _ := json.Marshal(val)
		buf.Write(b)
	}

	buf.WriteByte('{')
	if !e.Time.IsZero() {
		field("time", e.Time.Format(time.RFC3339Nano))
	}
	field("level", e.Level.String())
	if e.Prefix != "" {
		field("prefix", e.Prefix)
	}
	field("msg", e.Message)
	if e.Caller != "" {
		field("caller", e.Caller)
	}
	for _, a := range e.Attrs {
		if !first {
			buf.WriteByte(',')
		}
		first = false
		writeJSONAttr(&buf, a)
	}
	if e.Stack != "" {
		field("stacktrace", e.Stack)
	}
	buf.WriteByte('}')
	return buf.Bytes()
}

// TextFormatter renders an Entry as a single logfmt-style line.
func TextFormatter(e *Entry) []byte {
	return buildText(e, false)
}

// ColorTextFormatter renders an Entry like TextFormatter, but colors the
// level according to its severity using ANSI escape codes.
func ColorTextFormatter(e *Entry) []byte {
	return buildText(e, true)
}

// buildText renders e as a logfmt-style line, optionally coloring the level.
func buildText(e *Entry, colorize bool) []byte {
	var buf bytes.Buffer

	if !e.Time.IsZero() {
		buf.WriteString(e.Time.Format(time.RFC3339))
		buf.WriteByte(' ')
	}

	if colorize {
		buf.WriteString(levelColor(e.Level))
		buf.WriteString(e.Level.String())
		buf.WriteString(ansiReset)
	} else {
		buf.WriteString(e.Level.String())
	}

	if e.Prefix != "" {
		buf.WriteByte(' ')
		buf.WriteString(logfmtValue(e.Prefix))
	}
	buf.WriteByte(' ')
	buf.WriteString(logfmtValue(e.Message))

	appendTextAttrs(&buf, "", e.Attrs)

	if e.Caller != "" {
		buf.WriteString(" caller=")
		buf.WriteString(e.Caller)
	}
	if e.Stack != "" {
		buf.WriteByte('\n')
		buf.WriteString(e.Stack)
	}
	return buf.Bytes()
}

// levelColor returns the ANSI color code for l's severity.
func levelColor(l Level) string {
	switch {
	case l >= LevelError:
		return ansiRed
	case l >= LevelWarn:
		return ansiYellow
	case l >= LevelInfo:
		return ansiBlue
	default:
		return ansiGray
	}
}

// appendTextAttrs writes attrs as " key=value" pairs to buf, flattening
// nested groups into dot-joined keys.
func appendTextAttrs(buf *bytes.Buffer, prefix string, attrs []Attr) {
	for _, a := range attrs {
		v := a.Value.Resolve()
		key := a.Key
		if prefix != "" {
			key = prefix + "." + key
		}
		if v.Kind() == slog.KindGroup {
			appendTextAttrs(buf, key, v.Group())
			continue
		}
		buf.WriteByte(' ')
		buf.WriteString(key)
		buf.WriteByte('=')
		buf.WriteString(logfmtValue(fmt.Sprint(v.Any())))
	}
}

// logfmtValue quotes v if it contains spaces, quotes, or control characters.
func logfmtValue(v string) string {
	if v == "" {
		return `""`
	}
	for _, r := range v {
		if r <= ' ' || r == '"' || r == '=' {
			return strconv.Quote(v)
		}
	}
	return v
}

// writeJSONAttr writes a as a `"key":value` JSON member, expanding nested
// groups into JSON objects.
func writeJSONAttr(buf *bytes.Buffer, a Attr) {
	writeJSONString(buf, a.Key)
	buf.WriteByte(':')
	writeJSONValue(buf, a.Value)
}

// writeJSONValue writes v as a JSON value, expanding groups recursively.
func writeJSONValue(buf *bytes.Buffer, v slog.Value) {
	v = v.Resolve()
	if v.Kind() != slog.KindGroup {
		b, _ := json.Marshal(v.Any())
		buf.Write(b)
		return
	}
	buf.WriteByte('{')
	for i, ga := range v.Group() {
		if i > 0 {
			buf.WriteByte(',')
		}
		writeJSONAttr(buf, ga)
	}
	buf.WriteByte('}')
}

// writeJSONString writes s as a quoted JSON string.
func writeJSONString(buf *bytes.Buffer, s string) {
	b, _ := json.Marshal(s)
	buf.Write(b)
}
