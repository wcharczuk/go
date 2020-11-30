package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"

	"go.charczuk.com/sdk/selector"
)

// New creates a new logger.
func New(opts ...Option) *Log {
	l := Log{
		BaseContext: context.Background(),
		Output:      os.Stdout,
		Formatter:   NewTextFormatter(),
		Selector:    nil,
	}
	for _, opt := range opts {
		opt(&l)
	}
	return &l
}

// Option mutates a logger.
type Option func(*Log)

// OptOutput sets the logger output.
func OptOutput(wr io.Writer) Option {
	return func(l *Log) { l.Output = wr }
}

// OptBaseContext sets the logger BaseContext.
func OptBaseContext(ctx context.Context) Option {
	return func(l *Log) { l.BaseContext = ctx }
}

// OptFormatter sets the logger formatter.
func OptFormatter(f Formatter) Option {
	return func(l *Log) { l.Formatter = f }
}

// OptSelector sets the logger selector.
func OptSelector(sel selector.Selector) Option {
	return func(l *Log) { l.Selector = sel }
}

// OptSelectorParsed sets the logger selector.
func OptSelectorParsed(sel string) Option {
	return func(l *Log) { l.Selector = selector.MustParse(sel) }
}

// Log is the mainline implementation of a logger.
type Log struct {
	BaseContext context.Context
	Output      io.Writer
	Formatter   Formatter
	Selector    selector.Selector
}

// Background returns a context.
func (l Log) Background() context.Context {
	if l.BaseContext != nil {
		return l.BaseContext
	}
	return context.Background()
}

// Debug prints a message.
func (l Log) Debug(args ...interface{}) {
	l.EmitContext(l.Background(), l.stringMessage(FlagDebug, args...))
}

// Debugf prints a message.
func (l Log) Debugf(format string, args ...interface{}) {
	l.EmitContext(l.Background(), l.stringMessagef(FlagDebug, format, args...))
}

// Info prints a message.
func (l Log) Info(args ...interface{}) {
	l.EmitContext(l.Background(), l.stringMessage(FlagInfo, args...))
}

// Infof prints a message.
func (l Log) Infof(format string, args ...interface{}) {
	l.EmitContext(l.Background(), l.stringMessagef(FlagInfo, format, args...))
}

// Warning prints a message.
func (l Log) Warning(args ...interface{}) {
	l.EmitContext(l.Background(), l.stringMessage(FlagWarning, args...))
}

// Warningf prints a message.
func (l Log) Warningf(format string, args ...interface{}) {
	l.EmitContext(l.Background(), l.stringMessagef(FlagWarning, format, args...))
}

// Error prints a message.
func (l Log) Error(args ...interface{}) {
	l.EmitContext(l.Background(), l.stringMessage(FlagError, args...))
}

// Errorf prints a message.
func (l Log) Errorf(format string, args ...interface{}) {
	l.EmitContext(l.Background(), l.stringMessagef(FlagError, format, args...))
}

// Fatal prints a message.
func (l Log) Fatal(args ...interface{}) {
	l.EmitContext(l.Background(), l.stringMessage(FlagFatal, args...))
}

// Fatalf prints a message.
func (l Log) Fatalf(format string, args ...interface{}) {
	l.EmitContext(l.Background(), l.stringMessagef(FlagFatal, format, args...))
}

// Labels returns the full message labels.
func (l Log) Labels(labels map[string]string) map[string]string {
	return MergeLabels(
		GetLabels(l.Background()),
		labels,
	)
}

// Caller returns the caller at a given skip frame.
func (l Log) Caller(skip int) (caller runtime.Frame) {
	caller.PC, caller.File, caller.Line, _ = runtime.Caller(skip)
	return
}

// EmitContext emits a given message.
func (l Log) EmitContext(ctx context.Context, msg Message) {
	if GetSkip(ctx) {
		return
	}
	if l.Selector != nil && !l.Selector.Matches(msg.Labels()) {
		return
	}
	l.Formatter.Format(l.Output, msg)
}

func (l Log) stringMessage(flag string, args ...interface{}) StringMessage {
	return StringMessage{
		MessageCaller: l.Caller(3),
		Message:       fmt.Sprint(args...),
		MessageLabels: l.Labels(map[string]string{
			LabelFlag: flag,
		}),
	}
}

func (l Log) stringMessagef(flag, format string, args ...interface{}) StringMessage {
	return StringMessage{
		MessageCaller: l.Caller(3),
		Message:       fmt.Sprintf(format, args...),
		MessageLabels: l.Labels(map[string]string{
			LabelFlag: flag,
		}),
	}
}
