// Package gocolorlog provides a simple, extensible, and colored logging interface for Go applications.
// It supports colored output, HTTP-aware logging, and custom logger implementations.
package gocolorlog

import (
	"time"

	"github.com/ayberkgezer/gocolorlog/internal/level"
	"github.com/ayberkgezer/gocolorlog/internal/logger"
)

// Logger is the extensible logging interface for gocolorlog.
// You can implement this interface to customize logging behavior.
type Logger interface {
	// Debug logs a debug message.
	Debug(msg string)
	// Debugf logs a formatted debug message.
	Debugf(format string, args ...any)
	// Info logs an informational message.
	Info(msg string)
	// Infof logs a formatted informational message.
	Infof(format string, args ...any)
	// Warn logs a warning message.
	Warn(msg string)
	// Warnf logs a formatted warning message.
	Warnf(format string, args ...any)
	// Error logs an error message.
	Error(msg string)
	// Errorf logs a formatted error message.
	Errorf(format string, args ...any)
	// Fatal logs a fatal error message and exits the application.
	Fatal(msg string)
	// Fatalf logs a formatted fatal error message and exits the application.
	Fatalf(format string, args ...any)
	// HTTP logs an HTTP request with status, method, path, latency, and optional error.
	HTTP(status int, method, path string, latency time.Duration, ip string, requestID string, err error)
	// Context logs a message with a custom level and context.
	Context(level.Level, string, string, ...any)
}

var defaultLogger Logger = logger.NewStdLogger()

// SetLogger replaces the default logger with your own implementation.
// Pass nil to keep the current logger.
func SetLogger(l Logger) {
	if l != nil {
		defaultLogger = l
	}
}

// Debug logs a debug message using the default logger.
func Debug(msg string) { defaultLogger.Debug(msg) }

// Debugf logs a formatted debug message using the default logger.
func Debugf(format string, args ...any) { defaultLogger.Debugf(format, args...) }

// Fatal logs a fatal error message using the default logger and exits the application.
func Fatal(msg string) {
	defaultLogger.Fatal(msg)
}

// Fatalf logs a formatted fatal error message using the default logger and exits the application.
func Fatalf(format string, args ...any) {
	defaultLogger.Fatalf(format, args...)
}

// Info logs an informational message using the default logger.
func Info(msg string) { defaultLogger.Info(msg) }

// Infof logs a formatted informational message using the default logger.
func Infof(format string, args ...any) { defaultLogger.Infof(format, args...) }

// Warn logs a warning message using the default logger.
func Warn(msg string) { defaultLogger.Warn(msg) }

// Warnf logs a formatted warning message using the default logger.
func Warnf(format string, args ...any) { defaultLogger.Warnf(format, args...) }

// Error logs an error message using the default logger.
func Error(msg string) { defaultLogger.Error(msg) }

// Errorf logs a formatted error message using the default logger.
func Errorf(format string, args ...any) { defaultLogger.Errorf(format, args...) }

// HTTP logs an HTTP request using the default logger.
func HTTP(status int, method, path string, latency time.Duration, ip string, requestID string, err error) {
	defaultLogger.HTTP(status, method, path, latency, ip, requestID, err)
}

// Context logs a message with a custom level and context using the default logger.
func ContextLevel(level level.Level, context, msg string, args ...any) {
	defaultLogger.Context(level, context, msg, args...)
}
