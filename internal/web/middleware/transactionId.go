package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func WithRequestId(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "request-id", uuid.New().String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
