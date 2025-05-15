// Package colorlog provides a simple, colored and status-based logging interface for Go.
// HTTP status koduna göre renkli ve seviyeli loglama arayüzü sunar.
package gocolorlog

import (
	"time"

	"github.com/ayberkgezer/gocolorlog/internal/logger"
)

// Logger is the common interface for any structure that can log with color and status.
// Renkli ve status‐bazlı log atabilen her türlü yapı için ortak arayüzdür.
type Logger interface {
	// Log prints a log entry with status, method, path, latency and error.
	// Status, method, path, gecikme ve hata ile log kaydı oluşturur.
	Log(status int, method, path string, latency time.Duration, err error)
}

// NewLogger returns the default implementation (console colored logger).
// Varsayılan implementasyonu (konsola renkli log) döner.
func NewLogger() Logger {
	return logger.New()
}
