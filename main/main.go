package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/ayberkgezer/gocolorlog"
)

func main() {
	// Varsayılan renkli logger ile kullanım:
	gocolorlog.Info("Starting application")
	gocolorlog.Infof("Listening on port %d", 8080)
	gocolorlog.Warn("Cache miss for key user:123")
	gocolorlog.Warnf("Slow query: %s", "SELECT * FROM users")
	gocolorlog.Error("Failed to connect to DB")
	gocolorlog.Errorf("Failed to open file: %s", "config.yaml")

	gocolorlog.HTTP(200, "GET", "/api/test", 120*time.Millisecond, nil)
	gocolorlog.HTTP(404, "POST", "/api/notfound", 80*time.Millisecond, nil)
	gocolorlog.HTTP(500, "DELETE", "/api/error", 200*time.Millisecond, errors.New("internal server error"))

	gocolorlog.Info("Database connection: Successful")
	gocolorlog.ContextLevel("INFO", "Bootstrap", "Application is running on: %s", "http://localhost:3000")
	gocolorlog.ContextLevel("INFO", "Bootstrap", "Environment: %s", "development")
	gocolorlog.HTTP(201, "POST", "/ai-chat", 6*time.Millisecond, nil)
	gocolorlog.ContextLevel("WARN", "Bootstrap", "Cache miss for key %s", "user:123")
	gocolorlog.ContextLevel("ERROR", "Bootstrap", "Failed to connect to DB: %v", fmt.Errorf("timeout"))
}
