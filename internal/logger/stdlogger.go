package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ayberkgezer/gocolorlog/internal/color"
	"github.com/ayberkgezer/gocolorlog/internal/level"
)

type stdLogger struct {
	l *log.Logger
}

func NewStdLogger() *stdLogger {
	return &stdLogger{l: log.New(os.Stdout, "", 0)}
}

func (s *stdLogger) Debug(msg string) {
	s.Context(level.Debug, "App", "%s", msg)
}
func (s *stdLogger) Debugf(format string, args ...any) {
	s.Context(level.Debug, "App", format, args...)
}

func (s *stdLogger) Info(msg string) {
	s.Context(level.Info, "App", "%s", msg)
}

func (s *stdLogger) Infof(format string, args ...any) {
	s.Context(level.Info, "App", format, args...)
}

func (s *stdLogger) Warn(msg string) {
	s.Context(level.Warn, "App", "%s", msg)
}

func (s *stdLogger) Warnf(format string, args ...any) {
	s.Context(level.Warn, "App", format, args...)
}

func (s *stdLogger) Error(msg string) {
	s.Context(level.Error, "App", "%s", msg)
}
func (s *stdLogger) Errorf(format string, args ...any) {
	s.Context(level.Error, "App", format, args...)
}

func (s *stdLogger) Fatal(msg string) {
	s.Context(level.Fatal, "App", "%s", msg)
	os.Exit(1)
}
func (s *stdLogger) Fatalf(format string, args ...any) {
	s.Context(level.Fatal, "App", format, args...)
	os.Exit(1)
}

func (s *stdLogger) HTTP(status int, method, path string, latency time.Duration, ip string, requestID string, err error) {
	lvl := level.HTTP
	context := fmt.Sprintf("%d", status)
	msg := fmt.Sprintf("%s | %s | %dms - %s | %s", method, path, latency.Milliseconds(), latency, ip)
	if requestID != "" {
		msg += fmt.Sprintf(" | RequestID: %s", requestID)
	}
	if err != nil {
		msg += fmt.Sprintf(" | %s[Error]: %v %s", color.Red, err, color.Reset)
	}
	// Renkli context için özel fonksiyon
	s.ContextWithColor(lvl, context, statusColor(status), "%s", msg)
}

func (s *stdLogger) Context(lvl level.Level, context, msg string, args ...any) {
	pid := os.Getpid()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	col := colorForLevel(lvl)
	formatted := msg
	if len(args) > 0 {
		formatted = fmt.Sprintf(msg, args...)
	}
	s.l.Printf("%s[Log]%s %d - %s  %s%-5s%s [%s] %s",
		color.Cyan, color.Reset, pid, timestamp,
		col, lvl, color.Reset, context, formatted)
}

func (s *stdLogger) ContextWithColor(lvl level.Level, context, contextColor, msg string, args ...any) {
	pid := os.Getpid()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	col := colorForLevel(lvl)
	formatted := msg
	if len(args) > 0 {
		formatted = fmt.Sprintf(msg, args...)
	}
	s.l.Printf("%s[Log]%s %d - %s  %s%-5s%s [%s%s%s] %s",
		color.Cyan, color.Reset, pid, timestamp,
		col, lvl, color.Reset, contextColor, context, color.Reset, formatted)
}

func (s *stdLogger) ContextLevelWithRequestID(lvl level.Level, context, requestID, msg string, args ...any) {
	pid := os.Getpid()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	col := colorForLevel(lvl)
	formatted := msg
	if len(args) > 0 {
		formatted = fmt.Sprintf(msg, args...)
	}
	s.l.Printf("%s[Log]%s %d - %s  %s%-5s%s [%s] [%s] %s",
		color.Cyan, color.Reset, pid, timestamp,
		col, lvl, color.Reset, context, requestID, formatted)
}

func colorForLevel(lvl level.Level) string {
	switch lvl {
	case level.Info:
		return color.BrightGreen
	case level.Warn:
		return color.BrightYellow
	case level.Error:
		return color.BrightRed
	case level.HTTP:
		return color.BrightCyan
	case level.Debug:
		return color.BrightBlue
	case level.Fatal:
		return color.BgRed + color.White
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
