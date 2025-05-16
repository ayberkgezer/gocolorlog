package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ayberkgezer/gocolorlog/internal/color"
)

type stdLogger struct {
	l *log.Logger
}

func NewStdLogger() *stdLogger {
	return &stdLogger{l: log.New(os.Stdout, "", 0)}
}

func (s *stdLogger) Info(msg string) {
	s.Context("INFO", "App", "%s", msg)
}

func (s *stdLogger) Infof(format string, args ...any) {
	s.Context("INFO", "App", format, args...)
}

func (s *stdLogger) Warn(msg string) {
	s.Context("WARN", "App", "%s", msg)
}

func (s *stdLogger) Warnf(format string, args ...any) {
	s.Context("WARN", "App", format, args...)
}

func (s *stdLogger) Error(msg string) {
	s.Context("ERROR", "App", "%s", msg)
}

func (s *stdLogger) Errorf(format string, args ...any) {
	s.Context("ERROR", "App", format, args...)
}

func (s *stdLogger) HTTP(status int, method, path string, latency time.Duration, err error) {
	level := "HTTP"
	context := fmt.Sprintf("%d", status)
	msg := fmt.Sprintf("%s %s %d %dms - %s", method, path, status, latency.Milliseconds(), latency)
	if err != nil {
		msg += fmt.Sprintf(" handler error: %v", err)
	}
	// Renkli context için özel fonksiyon
	s.ContextWithColor(level, context, statusColor(status), "%s", msg)
}

func (s *stdLogger) Context(level, context, msg string, args ...any) {
	pid := os.Getpid()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	col := colorForLevel(level)
	formatted := msg
	if len(args) > 0 {
		formatted = fmt.Sprintf(msg, args...)
	}
	s.l.Printf("%s[Log]%s %d - %s  %s%-5s%s [%s] %s",
		color.Cyan, color.Reset, pid, timestamp,
		col, level, color.Reset, context, formatted)
}

func (s *stdLogger) ContextWithColor(level, context, contextColor, msg string, args ...any) {
	pid := os.Getpid()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	col := colorForLevel(level)
	formatted := msg
	if len(args) > 0 {
		formatted = fmt.Sprintf(msg, args...)
	}
	s.l.Printf("%s[Log]%s %d - %s  %s%-5s%s [%s%s%s] %s",
		color.Cyan, color.Reset, pid, timestamp,
		col, level, color.Reset, contextColor, context, color.Reset, formatted)
}

func colorForLevel(level string) string {
	switch level {
	case "INFO":
		return color.BrightGreen
	case "WARN":
		return color.BrightYellow
	case "ERROR":
		return color.BrightRed
	case "HTTP":
		return color.BrightCyan
	default:
		return color.White
	}
}

func statusColor(status int) string {
	switch {
	case status >= 500:
		return color.BrightRed
	case status >= 400:
		return color.BrightYellow
	default:
		return color.BrightGreen
	}
}
