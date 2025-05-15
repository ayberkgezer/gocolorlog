# gocolorlog

**gocolorlog** is a simple, extensible, and colored logging package for Go.
It maps HTTP status codes to colored log levels, making your terminal output clear and easy to read.

## Features

- **Colored output** based on HTTP status code:
  - Green for 2xx (Success)
  - Yellow for 4xx (Client Error)
  - Red for 5xx (Server Error)
- **Simple and extensible interface**
- **Console output by default**, but can be customized
- **Minimal dependencies**

## Installation

```sh
go get github.com/ayberkgezer/gocolorlog
```

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/ayberkgezer/gocolorlog"
)

func main() {
	log := gocolorlog.NewLogger()
	log.Log(200, "GET", "/api/test", 120*time.Millisecond, nil)
	log.Log(404, "POST", "/api/notfound", 80*time.Millisecond, nil)
	log.Log(500, "DELETE", "/api/error", 200*time.Millisecond,
		fmt.Errorf("internal server error"))
}
```

**Output:**
```
[INFO]  GET /api/test → 200 (120ms)
[WARN]  POST /api/notfound → 404 (80ms)
[ERROR] DELETE /api/error → 500 (200ms)
[ERROR] handler error: internal server error
```
*(Colors will be visible in a terminal that supports ANSI colors.)*

## API

### Logger Interface

```go
type Logger interface {
    Log(status int, method, path string, latency time.Duration, err error)
}
```

- `status`: HTTP status code (int)
- `method`: HTTP method (string)
- `path`: Request path (string)
- `latency`: Request duration (time.Duration)
- `err`: Error, if any (error)

### Creating a Logger

```go
log := gocolorlog.NewLogger()
```

You can also customize the output destination (see below).

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).

---

**Contributions and issues are welcome!**

---

If you need more advanced features or have suggestions, feel free to open an issue or pull request.

---

**Author:** [Ayberk Gezer](https://github.com/ayberkgezer)
