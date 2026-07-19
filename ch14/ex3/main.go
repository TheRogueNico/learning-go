package main

import (
	"context"
	"fmt"
	"net/http"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

type key struct{}

func WithLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, key{}, level)
}

// LevelFromContext extracts the logging level from the context.
// The second return value reports whether a valid level was present.
func LevelFromContext(ctx context.Context) (Level, bool) {
	level, ok := ctx.Value(key{}).(Level)
	return level, ok
}

func LevelMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch Level(r.URL.Query().Get("log_level")) {
		case Debug:
			r = r.WithContext(WithLevel(r.Context(), Debug))
		case Info:
			r = r.WithContext(WithLevel(r.Context(), Info))
		}

		h.ServeHTTP(w, r)
	})
}

func Log(ctx context.Context, level Level, message string) {
	inLevel, ok := LevelFromContext(ctx)
	if !ok {
		return
	}

	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

func main() {
}
