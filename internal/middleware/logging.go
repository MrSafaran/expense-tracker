package middleware

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf(
			"%s %s\n",
			r.Method,
			r.URL.Path,
		)

		next(w, r)
	}
}