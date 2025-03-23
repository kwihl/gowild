package handlers

import (
	"log/slog"
	"net/http"

	"gowild.com/internal"
)

type ForestHandler struct {
	service internal.ForestService
}

func NewForestHandler(service internal.ForestService) ForestHandler {
	return ForestHandler{
		service: service,
	}
}

var _ internal.ForestHandler = (*ForestHandler)(nil)

func (h *ForestHandler) Animals(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	defer func() {
		if r := recover(); r != nil {
			slog.ErrorContext(ctx, "recovered from panic in handler")
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}()

	switch req.Method {
	case http.MethodGet:
		_, err := h.service.ListForestAnimals(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "error when listing forest animals", "error", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		// Serve the resource.
	case http.MethodPost:
		// Create a new record.
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *ForestHandler) Plants(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	defer func() {
		if r := recover(); r != nil {
			slog.ErrorContext(ctx, "recovered from panic in handler")
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}()

	switch req.Method {
	case http.MethodGet:
		// Serve the resource.
	case http.MethodPost:
		// Create a new record.
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
