package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const requestIDKey contextKey = "request_id"

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.NewString()

			ctx := context.WithValue(
				r.Context(),
				requestIDKey,
				requestID,
			)

			r = r.WithContext(ctx)

			w.Header().Set(
				"X-Request-ID",
				requestID,
			)

			next.ServeHTTP(w, r)
		},
	)
}

func GetRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(requestIDKey).(string)

	if !ok {
		return ""
	}

	return requestID
}