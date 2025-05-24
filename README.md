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
	gocolorlog.Debug("Starting debug mode")
	gocolorlog.Debugf("Debug info: %s", "application initialization")
	gocolorlog.Info("Starting application")
	gocolorlog.Infof("Listening on port %d", 8080)
	gocolorlog.Warn("Cache miss for key user:123")
 	gocolorlog.Warnf("Slow query: %s", "SELECT * FROM users")
	gocolorlog.Error("Failed to connect to DB")
	gocolorlog.Errorf("Failed to open file: %s", "config.yaml")

	gocolorlog.HTTP(200, "GET", "/api/test", 120*time.Millisecond,"ip-adress", "",nil)
	gocolorlog.HTTP(404, "POST", "/api/notfound", 80*time.Millisecond,"ip-adress","Requsetid", nil)
	gocolorlog.HTTP(500, "DELETE", "/api/error", 200*time.Millisecond,"ip-adress", "Requsetid",fmt.Errorf("internal server error"))

	gocolorlog.ContextLevel("INFO", "Bootstrap", "Application is running on: %s", "http://localhost:3000")

	gocolorlog.Fatal("Critical error system shutdown")
	gocolorlog.Fatalf("Critical error: %v", fmt.Errorf("system failure"))
}
```

**Sample Output:**
```
[Log] 16639 - 2025-05-25 00:44:02  DEBUG [App] Starting debug mode
[Log] 16639 - 2025-05-25 00:44:02  DEBUG [App] Debug info: application initialization
[Log] 12345 - 2025-05-15 12:34:56  [INFO ] [App] Starting application
[Log] 12345 - 2025-05-15 12:34:56  [INFO ] [App] Listening on port 8080
[Log] 12345 - 2025-05-15 12:34:56  [WARN ] [App] Cache miss for key user:123
[Log] 12345 - 2025-05-15 12:34:56  [WARN ] [App] Slow query: SELECT * FROM users
[Log] 12345 - 2025-05-15 12:34:56  [ERROR] [App] Failed to connect to DB
[Log] 12345 - 2025-05-15 12:34:56  [ERROR] [App] Failed to open file: config.yaml
[Log] 62388 - 2025-05-16 21:29:01  HTTP  [200] GET | /api/test | 120ms - 120ms | 192.168.1.1
[Log] 62388 - 2025-05-16 21:29:01  HTTP  [404] POST | /api/notfound | 80ms - 80ms | RequestID: 71g261g61 |192.168.1.1 |
[Log] 62388 - 2025-05-16 21:29:01  HTTP  [500] DELETE | /api/error | 200ms - 200ms | 192.168.1.1 | RequestID: 71g261g61 | [Error]: internal server error
[Log] 12345 - 2025-05-15 12:34:56  [INFO ] [Bootstrap] Application is running on: http://localhost:3000

[Log] 16639 - 2025-05-25 00:44:02  FATAL [App] Critical error occurred - shutting down
```
*(Colors will be visible in a terminal that supports ANSI colors.)*

---

## API

### Logger Interface

```go
type Logger interface {
	Debug(msg string)
	Debugf(format string, args ...any)
    Info(msg string)
    Infof(format string, args ...any)
    Warn(msg string)
    Warnf(format string, args ...any)
    Error(msg string)
    Errorf(format string, args ...any)
    HTTP(status int, method, path string, latency time.Duration, ip string, requestID string, err error)
    ContextLevel(level, context, msg string, args ...any)
	Fatal(msg string)
	Fatalf(format string, args ...any)
}
```

#### Method Descriptions

- **Debug / Debugf**: Debug mod log
- **Info / Infof**: Log informational messages.
- **Warn / Warnf**: Log warnings.
- **Error / Errorf**: Log errors.
- **HTTP**: Log HTTP requests with status, method, path, latency, and optional error.
- **ContextLevel**: Log with a custom level and context.
- **Fatal / Fatalf**: Log critical errors and exit the application.

---

### Using the Default Logger

```go
gocolorlog.Info("Message")
gocolorlog.Warnf("Warning: %s", "details")
gocolorlog.HTTP(404, "POST", "/api/notfound", 80*time.Millisecond,"ip-adress", "",nil)
gocolorlog.ContextLevel("ERROR", "DB", "Connection error: %v", err)
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
