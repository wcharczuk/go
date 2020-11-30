package logger

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"go.charczuk.com/sdk/bufferutil"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

const (
	// Space is the space character.
	Space = " "
	// Newline is the newline character.
	Newline = "\n"
)

const (
	// DefaultTextTimeFormat is the default time format.
	DefaultTextTimeFormat = time.RFC3339Nano
)

// NewTextFormatter returns a new text formatter.
func NewTextFormatter(opts ...TextFormatterOption) *TextFormatter {
	tf := TextFormatter{
		BufferPool: bufferutil.NewPool(128),
	}
	for _, opt := range opts {
		opt(&tf)
	}
	return &tf
}

// TextFormatterOption mutates a text formatter
type TextFormatterOption func(*TextFormatter)

// OptTextFormatterNoColor sets the text formatter not to use color.
func OptTextFormatterNoColor(noColor bool) TextFormatterOption {
	return func(tf *TextFormatter) { tf.NoColor = noColor }
}

// OptTextFormatterNoTimestamp sets the text formatter not to show the timestamp.
func OptTextFormatterNoTimestamp(noTimestamp bool) TextFormatterOption {
	return func(tf *TextFormatter) { tf.NoTimestamp = noTimestamp }
}

// OptTextFormatterNoCaller sets the text formatter not to show the caller.
func OptTextFormatterNoCaller(noCaller bool) TextFormatterOption {
	return func(tf *TextFormatter) { tf.NoCaller = noCaller }
}

// OptTextFormatterNoLabels sets the text formatter not to show labels.
func OptTextFormatterNoLabels(noLabels bool) TextFormatterOption {
	return func(tf *TextFormatter) { tf.NoLabels = noLabels }
}

// OptTextFormatterTimeFormat sets the text formatter timestamp format.
func OptTextFormatterTimeFormat(timeFormat string) TextFormatterOption {
	return func(tf *TextFormatter) { tf.TimeFormat = timeFormat }
}

// TextFormatter is a formatter for text log messages.
type TextFormatter struct {
	NoColor     bool
	NoTimestamp bool
	NoCaller    bool
	NoLabels    bool
	TimeFormat  string
	BufferPool  *bufferutil.Pool
}

// TimeFormatOrDefault returns the time format or a default
func (tf TextFormatter) TimeFormatOrDefault() string {
	if tf.TimeFormat != "" {
		return tf.TimeFormat
	}
	return DefaultTextTimeFormat
}

// Colorize (optionally) applies a color to a string.
func (tf TextFormatter) Colorize(value string, color int) string {
	if tf.NoColor {
		return value
	}
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, value)
}

// FormatCaller formats the caller info.
func (tf TextFormatter) FormatCaller(caller runtime.Frame) string {
	return fmt.Sprintf("%s:%d", filepath.Base(caller.File), caller.Line)
}

// FormatTimestamp returns a new timestamp string.
func (tf TextFormatter) FormatTimestamp(ts time.Time) string {
	value := ts.Format(tf.TimeFormatOrDefault())
	return tf.Colorize(fmt.Sprintf("%-30s", value), gray)
}

// FormatLabels returns the scope labels section of the message as a string.
func (tf TextFormatter) FormatLabels(labels map[string]string) string {
	var keys []string
	for key := range labels {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var values []string
	for _, key := range keys {
		values = append(values, fmt.Sprintf("%s=%s", tf.Colorize(key, blue), labels[key]))
	}
	return strings.Join(values, " ")
}

// Format implements formatter.
func (tf TextFormatter) Format(output io.Writer, msg Message) {
	buffer := tf.BufferPool.Get()
	defer tf.BufferPool.Put(buffer)

	if !tf.NoTimestamp {
		buffer.WriteString(tf.FormatTimestamp(time.Now().UTC()))
		buffer.WriteString(Space)
	}
	if !tf.NoCaller {
		buffer.WriteString(tf.FormatCaller(msg.Caller()))
		buffer.WriteString(Space)
	}
	if !tf.NoLabels {
		labels := msg.Labels()
		if len(labels) > 0 {
			buffer.WriteString("\t")
			buffer.WriteString(tf.FormatLabels(labels))
		}
	}
	buffer.WriteString(Newline)
	_, _ = io.Copy(output, buffer)
	return
}
