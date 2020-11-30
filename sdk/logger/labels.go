package logger

// Constants
const (
	LabelFlag   = "flag"
	FlagDebug   = "debug"
	FlagInfo    = "info"
	FlagWarning = "warning"
	FlagError   = "error"
	FlagFatal   = "fatal"
)

// MergeLabels returns a merged variadic set of labels.
func MergeLabels(labels ...map[string]string) map[string]string {
	output := make(map[string]string)
	for _, labelSet := range labels {
		for key, value := range labelSet {
			output[key] = value
		}
	}
	return output
}
