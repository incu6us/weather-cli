package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	slogmulti "github.com/samber/slog-multi"
)

type Formatter uint8

const (
	UnknownFormatter Formatter = iota
	TextFormatter
	JSONFormatter
)

func FormatterFromStaring(s string) (Formatter, error) {
	s = strings.ToLower(s)
	switch s {
	case "text":
		return TextFormatter, nil
	case "json":
		return JSONFormatter, nil
	default:
		return UnknownFormatter, fmt.Errorf("unknown formatter: %s", s)
	}
}

type Logger interface {
	Handler() slog.Handler
	With(args ...any) *Log
	WithGroup(name string) *Log
	Enabled(ctx context.Context, level slog.Level) bool
	Log(ctx context.Context, level slog.Level, msg string, args ...any)
	LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr)
	Debugf(ctx context.Context, msg string, args ...any)
	Infof(ctx context.Context, msg string, args ...any)
	Warnf(ctx context.Context, msg string, args ...any)
	Errorf(ctx context.Context, msg string, args ...any)
}

type Log struct {
	*slog.Logger
}

func NewLog(handlers ...slog.Handler) *Log {
	return &Log{
		Logger: slog.New(slogmulti.Fanout(handlers...)),
	}
}

func parseLogLevel(s string) slog.Level {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	default:
		return slog.LevelError
	}
}

func NewHandler(logLevel string, formatter Formatter) slog.Handler {
	lvl := new(slog.LevelVar)
	lvl.Set(parseLogLevel(logLevel))

	handlerOptions := &slog.HandlerOptions{
		Level: lvl,
	}

	var handler slog.Handler
	switch formatter {
	case TextFormatter:
		handler = slog.NewTextHandler(os.Stdout, handlerOptions)
	case JSONFormatter:
		handler = slog.NewJSONHandler(os.Stdout, handlerOptions)
	}
	return handler
}

func (l *Log) With(args ...any) *Log {
	return &Log{Logger: l.Logger.With(args...)}
}

func (l *Log) WithGroup(name string) *Log {
	return &Log{Logger: l.Logger.WithGroup(name)}
}

func (l *Log) Debugf(ctx context.Context, msg string, args ...any) {
	msgParams, slogAttrs := l.paramsAndAttrs(args)
	l.Logger.DebugContext(ctx, fmt.Sprintf(msg, msgParams...), slogAttrs...)
}

func (l *Log) Infof(ctx context.Context, msg string, args ...any) {
	msgParams, slogAttrs := l.paramsAndAttrs(args)
	l.Logger.InfoContext(ctx, fmt.Sprintf(msg, msgParams...), slogAttrs...)
}

func (l *Log) Warnf(ctx context.Context, msg string, args ...any) {
	msgParams, slogAttrs := l.paramsAndAttrs(args)
	l.Logger.WarnContext(ctx, fmt.Sprintf(msg, msgParams...), slogAttrs...)
}

func (l *Log) Errorf(ctx context.Context, msg string, args ...any) {
	msgParams, slogAttrs := l.paramsAndAttrs(args)
	l.Logger.ErrorContext(ctx, fmt.Sprintf(msg, msgParams...), slogAttrs...)
}

func (l *Log) paramsAndAttrs(args []any) ([]any, []any) {
	msgParams := make([]any, 0, len(args))
	slogAttrs := make([]any, 0, len(args))
	for _, arg := range args {
		if attr, ok := arg.(slog.Attr); ok {
			slogAttrs = append(slogAttrs, attr)
			continue
		}
		msgParams = append(msgParams, arg)
	}
	return msgParams, slogAttrs
}
