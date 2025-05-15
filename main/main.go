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
