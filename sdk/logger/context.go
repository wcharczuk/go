package logger

import (
	"context"
	"time"
)

type skipKey struct{}

// WithSkip returns a context with a new value.
func WithSkip(ctx context.Context, skip bool) context.Context {
	return context.WithValue(ctx, skipKey{}, skip)
}

// GetSkip returns if we should skip a message on a given context.
func GetSkip(ctx context.Context) bool {
	if value := ctx.Value(skipKey{}); value != nil {
		if typed, ok := value.(bool); ok {
			return typed
		}
	}
	return false
}

type timestampKey struct{}

// WithTimestamp returns a new context with a given timestamp value.
func WithTimestamp(ctx context.Context, ts time.Time) context.Context {
	return context.WithValue(ctx, timestampKey{}, ts)
}

// GetTimestamp gets a timestampoff a context.
func GetTimestamp(ctx context.Context) time.Time {
	if raw := ctx.Value(timestampKey{}); raw != nil {
		if typed, ok := raw.(time.Time); ok {
			return typed
		}
	}
	return time.Time{}
}

type labelsKey struct{}

// WithLabels returns a new context with a given additional labels.
func WithLabels(ctx context.Context, labels map[string]string) context.Context {
	return context.WithValue(ctx, labelsKey{}, labels)
}

// WithLabel returns a new context with a given additional label.
func WithLabel(ctx context.Context, key, value string) context.Context {
	return context.WithValue(ctx, labelsKey{}, map[string]string{
		key: value,
	})
}

// GetLabels gets labels off a context.
func GetLabels(ctx context.Context) map[string]string {
	if raw := ctx.Value(labelsKey{}); raw != nil {
		if typed, ok := raw.(map[string]string); ok {
			return typed
		}
	}
	return nil
}
