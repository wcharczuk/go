package logger

var defaultLogger *Log

// Debug prints a debug message to the screen.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}
