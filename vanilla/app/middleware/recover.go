package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// Recover is a middleware that recovers panics and maps them to errors.
func Recover(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				var msg string
				switch x := r.(type) {
				case string:
					msg = fmt.Sprintf("panic: %s", x)
				case error:
					msg = fmt.Sprintf("panic: %s", x)
				default:
					msg = "unknown panic"
				}
				const size = 64 << 10 // 64KB
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				lines := strings.Split(string(buf), "\n")
				log.Printf("%s\n%s", msg, strings.Join(lines, "\n"))

				w.WriteHeader(500)
			}
		}()
		h(w, req)
	}
}
