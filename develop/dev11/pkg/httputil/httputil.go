package httputil

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ensiouel/apperror"
	"github.com/ensiouel/apperror/codes"
	"golang.org/x/exp/slog"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (writer *statusWriter) WriteHeader(statusCode int) {
	writer.status = statusCode
	writer.ResponseWriter.WriteHeader(statusCode)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := &statusWriter{ResponseWriter: w}

		now := time.Now()
		next.ServeHTTP(writer, r)

		slog.Info("request",
			slog.Int("status", writer.status),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Duration("latency", time.Since(now)),
		)
	})
}

func JSON(w http.ResponseWriter, code int, data any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return fmt.Errorf("encoding json: %w", err)
	}

	return nil
}

func Error(w http.ResponseWriter, err error) error {
	code := http.StatusInternalServerError

	var apperr *apperror.Error
	if errors.As(err, &apperr) {
		switch apperr.Code {
		case codes.Internal:
			code = http.StatusInternalServerError
		case codes.BadRequest:
			code = http.StatusBadRequest
		case codes.NotFound:
			code = http.StatusNotFound
		default:
			code = http.StatusTeapot
		}

		return JSON(w, code, map[string]any{"error": apperr})
	}

	return JSON(w, code, map[string]any{"error": err.Error()})
}

func GET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			JSON(w, http.StatusMethodNotAllowed, map[string]any{
				"error": "method not allowed",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func POST(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			JSON(w, http.StatusMethodNotAllowed, map[string]any{
				"error": "method not allowed",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
