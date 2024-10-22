package loggers

import (
	"context"
	"log/slog"
	"strconv"
)

type ContextHandler struct {
	slog.Handler
}

func (ch *ContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return ch.Handler.Enabled(ctx, level)
}

func (ch *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(addRequestId(ctx)...)
	return ch.Handler.Handle(ctx, r)
}

func (ch *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{ch.Handler.WithAttrs(attrs)}
}

func (ch *ContextHandler) WithGroup(name string) slog.Handler {
	return &ContextHandler{ch.Handler.WithGroup(name)}
}

func addRequestId(ctx context.Context) []slog.Attr {
	var as []slog.Attr

	as = append(as, slog.String("trace_id", getStringValue(ctx, "trace_id")))
	as = append(as, slog.String("span_id", getStringValue(ctx, "span_id")))

	group := slog.Group(
		"request",
		slog.String("host", getStringValue(ctx, "host")),
		slog.String("method", getStringValue(ctx, "method")),
	)
	as = append(as, group)
	return as
}

// getStringValue get default value from context
func getStringValue(ctx context.Context, key string) string {
	value := ""
	ctxValue := ctx.Value(key)
	if ctxValue != nil {
		stringValue, ok := ctxValue.(string)
		if !ok {
			return value
		}
		value = stringValue
	}
	return value
}

// getInt64Value get default value from context
func getInt64Value(ctx context.Context, key string) int64 {
	value := int64(0)
	ctxValue := ctx.Value(key)
	if ctxValue != nil {
		int64Value, ok := ctxValue.(int64)
		if !ok {
			return value
		}
		value = int64Value
	}
	return value
}

// getUInt64Value get default value from context
func getUInt64Value(ctx context.Context, key string) uint64 {
	value := uint64(0)
	ctxValue := ctx.Value(key)
	if ctxValue != nil {
		int64Value, ok := ctxValue.(uint64)
		if !ok {
			return value
		}
		value = int64Value
	}
	return value
}

// getUInt64Value get default value from context
func getUInt64ToStringValue(ctx context.Context, key string) string {
	value := uint64(0)
	ctxValue := ctx.Value(key)
	if ctxValue != nil {
		int64Value, ok := ctxValue.(uint64)
		if !ok {
			return "0"
		}
		value = int64Value
	}

	return strconv.FormatUint(value, 10)
}
