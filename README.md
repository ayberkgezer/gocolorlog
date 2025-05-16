# gocolorlog

**gocolorlog** is a simple, extensible, and colored logging package for Go.
It maps HTTP status codes to colored log levels, making your terminal output clear and easy to read.

---

## Features

- **Colored output** based on HTTP status code:
  - Green for 2xx (Success)
  - Yellow for 4xx (Client Error)
  - Red for 5xx (Server Error)
- **Simple and extensible interface**
- **Console output by default** (stdout)
- **Custom logger support** (implement your own)
- **Minimal dependencies**

---

## Installation

```sh
go get github.com/ayberkgezer/gocolorlog
```

---

## Usage

### Basic Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/ayberkgezer/gocolorlog"
)

func main() {
	gocolorlog.Info("Starting application")
	gocolorlog.Infof("Listening on port %d", 8080)
	gocolorlog.Warn("Cache miss for key user:123")
	gocolorlog.Warnf("Slow query: %s", "SELECT * FROM users")
	gocolorlog.Error("Failed to connect to DB")
	gocolorlog.Errorf("Failed to open file: %s", "config.yaml")

	gocolorlog.HTTP(200, "GET", "/api/test", 120*time.Millisecond, nil)
	gocolorlog.HTTP(404, "POST", "/api/notfound", 80*time.Millisecond, nil)
	gocolorlog.HTTP(500, "DELETE", "/api/error", 200*time.Millisecond, fmt.Errorf("internal server error"))

	gocolorlog.Context("INFO", "Bootstrap", "Application is running on: %s", "http://localhost:3000")
}
```

**Sample Output:**
```
[Log] 12345 - 2025-05-15 12:34:56  [INFO ] [App] Starting application
[Log] 12345 - 2025-05-15 12:34:56  [INFO ] [App] Listening on port 8080
[Log] 12345 - 2025-05-15 12:34:56  [WARN ] [App] Cache miss for key user:123
[Log] 12345 - 2025-05-15 12:34:56  [WARN ] [App] Slow query: SELECT * FROM users
[Log] 12345 - 2025-05-15 12:34:56  [ERROR] [App] Failed to connect to DB
[Log] 12345 - 2025-05-15 12:34:56  [ERROR] [App] Failed to open file: config.yaml
[Log] 12345 - 2025-05-15 12:34:56  [HTTP ] [200] GET /api/test 200 120ms - 120ms
[Log] 12345 - 2025-05-15 12:34:56  [HTTP ] [404] POST /api/notfound 404 80ms - 80ms
[Log] 12345 - 2025-05-15 12:34:56  [HTTP ] [500] DELETE /api/error 500 200ms - 200ms handler error: internal server error
[Log] 12345 - 2025-05-15 12:34:56  [INFO ] [Bootstrap] Application is running on: http://localhost:3000
```
*(Colors will be visible in a terminal that supports ANSI colors.)*

---

## API

### Logger Interface

```go
type Logger interface {
    Info(msg string)
    Infof(format string, args ...any)
    Warn(msg string)
    Warnf(format string, args ...any)
    Error(msg string)
    Errorf(format string, args ...any)
    HTTP(status int, method, path string, latency time.Duration, err error)
    Context(level, context, msg string, args ...any)
}
```

#### Method Descriptions

- **Info / Infof**: Log informational messages.
- **Warn / Warnf**: Log warnings.
- **Error / Errorf**: Log errors.
- **HTTP**: Log HTTP requests with status, method, path, latency, and optional error.
- **Context**: Log with a custom level and context.

---

### Using the Default Logger

```go
gocolorlog.Info("Message")
gocolorlog.Warnf("Warning: %s", "details")
gocolorlog.HTTP(404, "POST", "/api/notfound", 80*time.Millisecond, nil)
gocolorlog.Context("ERROR", "DB", "Connection error: %v", err)
```

---

### Custom Logger

You can implement your own logger by satisfying the `Logger` interface and set it as the default:

```go
type MyLogger struct{}

func (m *MyLogger) Info(msg string) { /* ... */ }
func (m *MyLogger) Infof(format string, args ...any) { /* ... */ }
// ... implement other methods

gocolorlog.SetLogger(&MyLogger{})
```

---

### Redirecting Output

By default, logs are written to stdout.
To write logs to a file or elsewhere, implement your own logger or use the internal options package.

---

## Advanced Features

- **Colored log levels**: INFO (green), WARN (yellow), ERROR (red), HTTP (cyan)
- **PID and timestamp in output**
- **Easily extensible**
- **Customizable formatting and coloring**

---

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).

---

**Contributions and issues are welcome!**

---

**Author:** [Ayberk Gezer](https://github.com/ayberkgezer)
