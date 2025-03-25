package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"runtime/debug"
)

func recoveryFunc(ctx context.Context, w http.ResponseWriter) {
	if r := recover(); r != nil {
		slog.ErrorContext(ctx, "recovered from panic in handler", "r", r, "stacktrace", debug.Stack())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
