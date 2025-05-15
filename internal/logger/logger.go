// Package logger implements the default colored logger for colorlog.
// colorlog için varsayılan renkli logger implementasyonunu içerir.
package logger

import (
	"log"
	"time"

	"github.com/ayberkgezer/gocolorlog/internal/color"
	"github.com/ayberkgezer/gocolorlog/internal/level"
	"github.com/ayberkgezer/gocolorlog/internal/options"
)

// Logger is the default logger type for colored output.
// Renkli çıktı için varsayılan logger tipidir.
type Logger struct {
	out *log.Logger
}

// New creates a new Logger instance with optional configuration.
// Opsiyonel ayarlarla yeni bir Logger örneği oluşturur.
func New(opts ...options.Option) *Logger {
	cfg := options.DefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	writer := cfg.Writer
	if writer == nil {
		writer = log.Default().Writer()
	}
	return &Logger{
		out: log.New(writer, "", 0),
	}
}

// Log prints a colored log entry based on HTTP status code and error.
// HTTP status koduna ve hataya göre renkli log çıktısı üretir.
func (l *Logger) Log(status int, method, path string, latency time.Duration, err error) {
	clr, lvl := pickColorAndLevel(status)
	l.out.Printf(
		"%s[%s]%s %s %s → %d (%s)",
		clr, lvl, color.Reset,
		method, path, status, latency,
	)
	if err != nil {
		l.out.Printf(
			"%s[%s]%s handler error: %v",
			color.Red, level.Error, color.Reset, err,
		)
	}
}

// pickColorAndLevel returns the color and log level for a given HTTP status code.
// Verilen HTTP status kodu için renk ve log seviyesini döner.
func pickColorAndLevel(status int) (string, level.Level) {
	switch {
	case status >= 500:
		return color.Red, level.Error
	case status >= 400:
		return color.Yellow, level.Warn
	default:
		return color.Green, level.Info
	}
}
