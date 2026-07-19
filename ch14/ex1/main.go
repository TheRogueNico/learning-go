package main

import (
	"context"
	"net/http"
	"time"
)

func TimeoutMiddleware(ms int) func(http.Handler) http.Handler {
	d := time.Duration(ms) * time.Millisecond

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), d)
			defer cancel()

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func main() {
}
