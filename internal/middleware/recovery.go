package middleware

import (
	"fmt"
	"net/http"
)

func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {

			if err := recover(); err != nil {

				fmt.Println("panic recovered:", err)

				http.Error(
					w,
					"internal server error",
					http.StatusInternalServerError,
				)
			}
		}()

		next(w, r)
	}
}