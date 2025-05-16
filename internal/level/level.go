// Package level defines log levels for colorlog.
// colorlog için log seviyelerini tanımlar.
package level

// Level represents the log level type.
// Log seviyesi tipini temsil eder.
type Level string

const (
	Info  Level = "INFO"  // INFO
	Warn  Level = "WARN"  // WARN
	Error Level = "ERROR" // ERROR
	HTTP  Level = "HTTP"  // HTTP
)
