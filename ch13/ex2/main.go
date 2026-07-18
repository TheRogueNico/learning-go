package main

import (
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

func logging(logger *slog.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}

		logger.Info("request received", "ip", ip)
		h.ServeHTTP(w, r)
	})
}

func main() {
	options := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stderr, options)
	mySlog := slog.New(handler)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().Format(time.RFC3339)))
	})

	s := &http.Server{
		Addr:         ":8080",
		Handler:      logging(mySlog, mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
