// Package options provides configuration options for the logger.
// Logger için yapılandırma seçeneklerini sunar.
package options

import "io"

// Option is a function that modifies the logger configuration.
// Logger yapılandırmasını değiştiren fonksiyon tipidir.
type Option func(cfg *config)

// config holds logger configuration.
// Logger yapılandırmasını tutar.
type config struct {
	Writer io.Writer
}

// DefaultConfig returns the default logger configuration.
// Varsayılan logger yapılandırmasını döner.
func DefaultConfig() *config {
	return &config{Writer: nil}
}

// WithWriter sets a custom writer for the logger.
// Logger için özel bir writer ayarlamanızı sağlar.
func WithWriter(w io.Writer) Option {
	return func(cfg *config) {
		cfg.Writer = w
	}
}
