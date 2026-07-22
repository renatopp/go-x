package logx

import (
	"strings"
	"testing"
	"time"
)

func TestSystemTimeFormatter(t *testing.T) {
	tm := time.Date(2024, 3, 5, 13, 4, 5, 600_000_000, time.UTC)
	got := SystemTimeFormatter(tm, LevelInfo)
	want := "2024-03-05 13:04:05.600"
	if got != want {
		t.Errorf("SystemTimeFormatter: got %q, want %q", got, want)
	}
}

func TestShortTimeFormatter(t *testing.T) {
	tm := time.Date(2024, 3, 5, 13, 4, 5, 0, time.UTC)
	got := ShortTimeFormatter(tm, LevelInfo)
	want := "13:04:05"
	if got != want {
		t.Errorf("ShortTimeFormatter: got %q, want %q", got, want)
	}
}

func TestPMTimeFormatter(t *testing.T) {
	tm := time.Date(2024, 3, 5, 13, 4, 5, 0, time.UTC)
	got := PMTimeFormatter(tm, LevelInfo)
	want := "01:04PM"
	if got != want {
		t.Errorf("PMTimeFormatter: got %q, want %q", got, want)
	}
}

func TestColoredLevelFormatter(t *testing.T) {
	tests := []struct {
		level Level
		want  string
	}{
		{LevelDebug, "DEBUG"},
		{LevelInfo, "INFO"},
		{LevelWarn, "WARN"},
		{LevelError, "ERROR"},
		{LevelFatal, "FATAL"},
		{Level(999), "UNKN"},
	}
	for _, tt := range tests {
		got := ColoredLevelFormatter(tt.level)
		if !strings.Contains(got, tt.want) {
			t.Errorf("ColoredLevelFormatter(%v): got %q, want to contain %q", tt.level, got, tt.want)
		}
	}
}

func TestCallerFormatter(t *testing.T) {
	src := &Source{File: "f.go", Function: "pkg.Fn", Line: 42}
	got := CallerFormatter(src, LevelInfo)
	for _, want := range []string{"f.go", "pkg.Fn", "42"} {
		if !strings.Contains(got, want) {
			t.Errorf("CallerFormatter: got %q, want to contain %q", got, want)
		}
	}
}

func TestAttrFormatter(t *testing.T) {
	attr := String("key", "value")

	got := AttrFormatter(attr, nil, LevelInfo)
	if !strings.Contains(got, "key=value") {
		t.Errorf("AttrFormatter: got %q, want to contain %q", got, "key=value")
	}

	got = AttrFormatter(attr, []string{"g1", "g2"}, LevelInfo)
	if !strings.Contains(got, "g1.g2.key=value") {
		t.Errorf("AttrFormatter with groups: got %q, want to contain %q", got, "g1.g2.key=value")
	}
}
