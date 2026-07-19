package logx

import (
	"context"
)

type Logger struct {
	handler Handler
}

func (l *Logger) Handler() Handler
func (l *Logger) WithAttributes(attrs ...any) *Logger
func (l *Logger) WithGroup(name string) *Logger

func (l *Logger) Log(level Level, msg string, kvargs ...any)
func (l *Logger) Debug(msg string, kvargs ...any)
func (l *Logger) Info(msg string, kvargs ...any)
func (l *Logger) Warn(msg string, kvargs ...any)
func (l *Logger) Error(msg string, kvargs ...any)
func (l *Logger) Fatal(msg string, kvargs ...any)

func (l *Logger) Logf(level Level, msg string, args ...any)
func (l *Logger) Debugf(msg string, args ...any)
func (l *Logger) Infof(msg string, args ...any)
func (l *Logger) Warnf(msg string, args ...any)
func (l *Logger) Errorf(msg string, args ...any)
func (l *Logger) Fatalf(msg string, args ...any)

func (l *Logger) Logc(ctx context.Context, level Level, msg string, kvargs ...any)
func (l *Logger) Debugc(ctx context.Context, msg string, kvargs ...any)
func (l *Logger) Infoc(ctx context.Context, msg string, kvargs ...any)
func (l *Logger) Warnc(ctx context.Context, msg string, kvargs ...any)
func (l *Logger) Errorc(ctx context.Context, msg string, kvargs ...any)
func (l *Logger) Fatalc(ctx context.Context, msg string, kvargs ...any)

func (l *Logger) log(ctx context.Context, level Level, msg string, kvargs ...any)
