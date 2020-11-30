package assert

import "os"

// OutputFormatFromEnv gets the output format from the env or the default.
func OutputFormatFromEnv() OutputFormat {
	outputFormat := OutputFormatText
	if envOutputFormat := os.Getenv("TEST_OUTPUT_FORMAT"); envOutputFormat != "" {
		outputFormat = OutputFormat(envOutputFormat)
	}
	return outputFormat
}

// OutputFormat is an assertion error output format.
type OutputFormat string

// OutputFormats
const (
	OutputFormatDefault OutputFormat = ""
	OutputFormatText    OutputFormat = "text"
	OutputFormatJSON    OutputFormat = "json"
)

const (
	assertionFailedLabel = "Assertion Failed!"
	locationLabel        = "Assert Location"
	assertionLabel       = "Assertion"
	messageLabel         = "Message"
	expectedLabel        = "Expected"
	actualLabel          = "Actual"
)
