package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(h *Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))

	r.Get("/health", h.Health)
	r.Get("/api/v1/fleet", h.FleetAnalytics)
	r.Get("/api/v1/fleet/{tailNumber}/history", h.AirframeHistory)
	r.Get("/api/v1/analytics/downtime", h.DowntimeAnalytics)

	return r
}
