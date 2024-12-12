package logger

import (
	"context"
	"io"
	"log/slog"
)

type DiscardLogger struct {
	log *Log
}

func NewDiscardLogger() *DiscardLogger {
	return &DiscardLogger{log: &Log{slog.New(slog.NewTextHandler(io.Discard, nil))}}
}

func (l *DiscardLogger) Handler() slog.Handler {
	return l.log.Handler()
}

func (l *DiscardLogger) With(args ...any) *Log {
	return &Log{Logger: l.log.Logger.With(args...)}
}

func (l *DiscardLogger) WithGroup(name string) *Log {
	return &Log{Logger: l.log.Logger.WithGroup(name)}
}

func (l *DiscardLogger) Enabled(ctx context.Context, level slog.Level) bool {
	return l.log.Enabled(ctx, level)
}

func (l *DiscardLogger) Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	l.log.Log(ctx, level, msg, args...)
}

func (l *DiscardLogger) LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	l.log.LogAttrs(ctx, level, msg, attrs...)
}

func (l *DiscardLogger) Debugf(_ context.Context, _ string, _ ...any) {
}

func (l *DiscardLogger) Infof(_ context.Context, _ string, _ ...any) {
}

func (l *DiscardLogger) Warnf(_ context.Context, _ string, _ ...any) {
}

func (l *DiscardLogger) Errorf(_ context.Context, _ string, _ ...any) {
}
