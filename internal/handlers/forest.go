package handlers

import (
	"encoding/json"
	"io"
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
	slog.InfoContext(ctx, "/v1/animals received request")

	defer recoveryFunc(ctx, w)

	// Read data and check if there's payload
	data, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	var bodyIsEmpty bool = (len(data) == 0)

	if err != nil {
		slog.ErrorContext(ctx, "error reading request body", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// Determine request type and call appropriate handler
	switch req.Method {
	case http.MethodGet:
		// Check if it's a a get by ID or a get all
		if bodyIsEmpty {
			animals, err := h.service.ListForestAnimals(ctx)
			if err != nil {
				slog.ErrorContext(ctx, "error when listing forest animals", "error", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			resDTOs := animalsFromDomain(animals)
			res, err := json.Marshal(resDTOs)
			if err != nil {
				slog.ErrorContext(ctx, "error marshalling response", "error", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			w.Write(res)
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
	defer recoveryFunc(ctx, w)

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
