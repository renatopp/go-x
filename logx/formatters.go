package logx

import (
	"fmt"
	"strings"
	"time"
)

func SystemTimeFormatter(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.000")
}

func ShortTimeFormatter(t time.Time) string {
	return t.Format("15:04:05")
}

func PMTimeFormatter(t time.Time) string {
	return t.Format("03:04PM")
}

func ColoredLevelFormatter(level Level) string {
	switch level {
	case LevelDebug:
		return "\033[1;38;5;63mDEBUG\033[0m"
	case LevelInfo:
		return "\033[1;38;5;86mINFO \033[0m"
	case LevelWarn:
		return "\033[1;38;5;192mWARN \033[0m"
	case LevelError:
		return "\033[1;38;5;204mERROR\033[0m"
	case LevelFatal:
		return "\033[1;38;5;134mFATAL\033[0m"
	default:
		return "\033[1;38;5;245mUNKN \033[0m"
	}
}

func CallerFormatter(source *Source) string {
	return fmt.Sprintf("\033[2mat %s@%s:%d\033[0m", source.File, source.Function, source.Line)
}

func AttrFormatter(attr Attr, groups []string) string {
	key := attr.Key
	if len(groups) > 0 {
		key = strings.Join(groups, ".") + "." + key
	}
	return fmt.Sprintf("\033[2m%s=%s\033[0m", key, attr.Value)
}

func CommonOutputFormatter(parts OutputParts) string {
	builder := strings.Builder{}
	if parts.Time != "" {
		builder.WriteString(parts.Time)
		builder.WriteString(" ")
	}

	if parts.Level != "" {
		builder.WriteString(parts.Level)
		builder.WriteString(" ")
	}

	if parts.Caller != "" {
		builder.WriteString(parts.Caller)
		builder.WriteString(" ")
	}

	if parts.Message != "" {
		builder.WriteString(parts.Message)
		builder.WriteString(" ")
	}

	for _, attr := range parts.Attrs {
		builder.WriteString(attr)
		builder.WriteString(" ")
	}
	return builder.String()
}
