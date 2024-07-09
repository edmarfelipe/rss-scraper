package server

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/edmarfelipe/rss-scraper/internal/server/openapi"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"
)

func errorHandleFunc(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, validate.ErrBodyRequired) {
		WriteJSON(w, http.StatusBadRequest, openapi.ErrorResponse{Error: "Body is required"})
		return
	}

	if errors.Is(err, validate.ErrFieldRequired) {
		WriteJSON(w, http.StatusBadRequest, openapi.ErrorResponse{Error: "Field is required"})
		return
	}

	if errors.Is(err, ogenerrors.ErrSecurityRequirementIsNotSatisfied) {
		WriteJSON(w, http.StatusUnauthorized, openapi.ErrorResponse{Error: "Authentication required"})
		return
	}

	if ogenerrors.ErrorCode(err) == http.StatusInternalServerError {
		slog.Error("Internal Server Error", "error", err)
	}

	ogenerrors.DefaultErrorHandler(ctx, w, r, err)
}

// WriteJSON writes a JSON response.
func WriteJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
