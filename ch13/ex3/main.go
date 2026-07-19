package main

import (
	"encoding/json"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type timeResponse struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

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

func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	if strings.Contains(r.Header.Get("Accept"), "application/json") {
		resp := timeResponse{
			DayOfWeek:  now.Weekday().String(),
			DayOfMonth: now.Day(),
			Month:      now.Month().String(),
			Year:       now.Year(),
			Hour:       now.Hour(),
			Minute:     now.Minute(),
			Second:     now.Second(),
		}
		data, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return
	}

	w.Write([]byte(now.Format(time.RFC3339)))
}

func main() {
	options := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stderr, options)
	mySlog := slog.New(handler)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", timeHandler)

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
