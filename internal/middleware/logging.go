package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		recorder := &ResponseRecorder{
			writer: w,
			status: http.StatusOK,
		}

		next(recorder, r)

		duration := time.Since(start)

		fmt.Printf(
			"%s %s %s %d %v\n",
			time.Now().UTC().Format(time.RFC3339),
			r.Method,
			r.URL.Path,
			recorder.status,
			duration,
		)
	}
}