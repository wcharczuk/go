package log

import "context"

var defaultLogger *Log

func init() {
	defaultLogger = New()
}

// Debug prints a message to the screen.
func Debug(args ...interface{}) {
	defaultLogger.EmitContext(defaultLogger.Background(), defaultLogger.StringMessage(FlagDebug, args...))
}

// Debugf prints a message to the screen.
func Debugf(format string, args ...interface{}) {
	defaultLogger.EmitContext(defaultLogger.Background(), defaultLogger.StringMessagef(FlagDebug, format, args...))
}

// Info prints a message to the screen.
func Info(args ...interface{}) {
	defaultLogger.EmitContext(defaultLogger.Background(), defaultLogger.StringMessage(FlagInfo, args...))
}

// Infof prints a message to the screen.
func Infof(format string, args ...interface{}) {
	defaultLogger.EmitContext(defaultLogger.Background(), defaultLogger.StringMessagef(FlagInfo, format, args...))
}

// Warning prints a message to the screen.
func Warning(args ...interface{}) {
	defaultLogger.EmitContext(defaultLogger.Background(), defaultLogger.StringMessage(FlagWarning, args...))
}

// Warningf prints a message to the screen.
func Warningf(format string, args ...interface{}) {
	defaultLogger.EmitContext(defaultLogger.Background(), defaultLogger.StringMessagef(FlagWarning, format, args...))
}

// Error prints a message to the screen.
func Error(args ...interface{}) {
	defaultLogger.EmitContext(defaultLogger.Background(), defaultLogger.StringMessage(FlagError, args...))
}

// Errorf prints a message to the screen.
func Errorf(format string, args ...interface{}) {
	defaultLogger.EmitContext(defaultLogger.Background(), defaultLogger.StringMessagef(FlagError, format, args...))
}

// Fatal prints a message to the screen.
func Fatal(args ...interface{}) {
	defaultLogger.EmitContext(defaultLogger.Background(), defaultLogger.StringMessage(FlagFatal, args...))
}

// Fatalf prints a message to the screen.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.EmitContext(defaultLogger.Background(), defaultLogger.StringMessagef(FlagFatal, format, args...))
}

// EmitContext prints a message to the screen.
func EmitContext(ctx context.Context, msg Message) {
	defaultLogger.EmitContext(ctx, msg)
}
