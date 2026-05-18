package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mtepenner/airframe-lifecycle-tracker/analytics_engine/internal/models"
	"github.com/mtepenner/airframe-lifecycle-tracker/analytics_engine/internal/service"
)

type Handler struct {
	service *service.FleetService
}

func NewHandler(service *service.FleetService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) FleetAnalytics(w http.ResponseWriter, r *http.Request) {
	year := 0
	if rawYear := r.URL.Query().Get("year"); rawYear != "" {
		parsed, err := strconv.Atoi(rawYear)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "invalid year"})
			return
		}
		year = parsed
	}

	model := r.URL.Query().Get("model")
	operator := r.URL.Query().Get("operator")

	response, err := h.service.GetFleetAnalytics(r.Context(), year, model, operator)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, response)
}

func (h *Handler) DowntimeAnalytics(w http.ResponseWriter, r *http.Request) {
	from := time.Now().UTC().AddDate(-1, 0, 0)
	to := time.Now().UTC()

	if rawFrom := r.URL.Query().Get("from"); rawFrom != "" {
		parsed, err := time.Parse("2006-01-02", rawFrom)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "invalid from date, expected YYYY-MM-DD"})
			return
		}
		from = parsed
	}

	if rawTo := r.URL.Query().Get("to"); rawTo != "" {
		parsed, err := time.Parse("2006-01-02", rawTo)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "invalid to date, expected YYYY-MM-DD"})
			return
		}
		to = parsed.Add(24 * time.Hour)
	}

	if !from.Before(to) {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "from must be before to"})
		return
	}

	model := r.URL.Query().Get("model")
	response, err := h.service.GetDowntimeAnalytics(r.Context(), from, to, model)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, response)
}

func (h *Handler) AirframeHistory(w http.ResponseWriter, r *http.Request) {
	tailNumber := chi.URLParam(r, "tailNumber")
	if tailNumber == "" {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "tail number is required"})
		return
	}

	response, err := h.service.GetAirframeHistory(r.Context(), tailNumber)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, response)
}

func writeJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(payload)
}
