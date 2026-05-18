package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	api "github.com/mtepenner/airframe-lifecycle-tracker/analytics_engine/internal/api"
	"github.com/mtepenner/airframe-lifecycle-tracker/analytics_engine/internal/repository"
	"github.com/mtepenner/airframe-lifecycle-tracker/analytics_engine/internal/service"
)

func main() {
	databaseURL := getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/airframe?sslmode=disable")
	port := getEnv("PORT", "8080")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("failed to create pool: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := repository.New(pool)
	svc := service.New(repo)
	h := api.NewHandler(svc)
	router := api.NewRouter(h)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("analytics engine listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
