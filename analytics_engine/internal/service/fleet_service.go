package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mtepenner/airframe-lifecycle-tracker/analytics_engine/internal/models"
	"github.com/mtepenner/airframe-lifecycle-tracker/analytics_engine/internal/repository"
)

type FleetService struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *FleetService {
	return &FleetService{repo: repo}
}

func (s *FleetService) GetFleetAnalytics(ctx context.Context, year int, model, operator string) (models.FleetAnalyticsResponse, error) {
	rows, err := s.repo.GetFleetCube(ctx, year, model, operator)
	if err != nil {
		return models.FleetAnalyticsResponse{}, err
	}

	events, err := s.repo.FleetEventCount(rows)
	if err != nil {
		return models.FleetAnalyticsResponse{}, err
	}

	return models.FleetAnalyticsResponse{
		At:          time.Now().UTC(),
		Rows:        rows,
		TotalEvents: events,
	}, nil
}

func (s *FleetService) GetDowntimeAnalytics(ctx context.Context, from, to time.Time, model string) (models.DowntimeResponse, error) {
	var (
		points []models.DowntimePoint
		total  int
	)

	errCh := make(chan error, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		res, err := s.repo.GetDowntimePoints(ctx, from, to, model)
		if err != nil {
			errCh <- fmt.Errorf("fetch downtime points: %w", err)
			return
		}
		points = res
	}()

	go func() {
		defer wg.Done()
		res, err := s.repo.GetDowntimeTotal(ctx, from, to, model)
		if err != nil {
			errCh <- fmt.Errorf("fetch downtime total: %w", err)
			return
		}
		total = res
	}()

	wg.Wait()
	close(errCh)
	for err := range errCh {
		if err != nil {
			return models.DowntimeResponse{}, err
		}
	}

	return models.DowntimeResponse{
		From:       from,
		To:         to,
		Points:     points,
		TotalHours: total,
	}, nil
}

func (s *FleetService) GetAirframeHistory(ctx context.Context, tailNumber string) ([]models.AirframeSnapshot, error) {
	return s.repo.GetAirframeHistory(ctx, tailNumber)
}
