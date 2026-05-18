package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mtepenner/airframe-lifecycle-tracker/analytics_engine/internal/models"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) GetFleetCube(ctx context.Context, year int, model, operator string) ([]models.FleetCubeRow, error) {
	query := `
		SELECT model, operator, year, total_downtime_hours, maintenance_events
		FROM mv_fleet_cube
		WHERE ($1::int = 0 OR year = $1)
		  AND ($2 = '' OR model = $2)
		  AND ($3 = '' OR operator = $3)
		ORDER BY year DESC, model, operator`

	rows, err := r.pool.Query(ctx, query, year, model, operator)
	if err != nil {
		return nil, fmt.Errorf("query fleet cube: %w", err)
	}
	defer rows.Close()

	out := make([]models.FleetCubeRow, 0)
	for rows.Next() {
		var row models.FleetCubeRow
		if err := rows.Scan(&row.Model, &row.Operator, &row.Year, &row.TotalDowntimeHours, &row.MaintenanceEvents); err != nil {
			return nil, fmt.Errorf("scan fleet cube row: %w", err)
		}
		out = append(out, row)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate fleet cube rows: %w", err)
	}
	return out, nil
}

func (r *Repository) GetDowntimePoints(ctx context.Context, from, to time.Time, model string) ([]models.DowntimePoint, error) {
	query := `
		SELECT date_trunc('day', performed_at)::date, downtime_hours, tail_number, model
		FROM maintenance_logs
		WHERE performed_at >= $1
		  AND performed_at < $2
		  AND ($3 = '' OR model = $3)
		ORDER BY performed_at ASC`

	rows, err := r.pool.Query(ctx, query, from, to, model)
	if err != nil {
		return nil, fmt.Errorf("query downtime points: %w", err)
	}
	defer rows.Close()

	out := make([]models.DowntimePoint, 0)
	for rows.Next() {
		var row models.DowntimePoint
		if err := rows.Scan(&row.Date, &row.Hours, &row.TailNumber, &row.Model); err != nil {
			return nil, fmt.Errorf("scan downtime point: %w", err)
		}
		out = append(out, row)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate downtime points: %w", err)
	}
	return out, nil
}

func (r *Repository) GetDowntimeTotal(ctx context.Context, from, to time.Time, model string) (int, error) {
	query := `
		SELECT COALESCE(SUM(downtime_hours), 0)
		FROM maintenance_logs
		WHERE performed_at >= $1
		  AND performed_at < $2
		  AND ($3 = '' OR model = $3)`

	var total int
	if err := r.pool.QueryRow(ctx, query, from, to, model).Scan(&total); err != nil {
		return 0, fmt.Errorf("query downtime total: %w", err)
	}
	return total, nil
}

func (r *Repository) FleetEventCount(rows []models.FleetCubeRow) (int, error) {
	if rows == nil {
		return 0, errors.New("rows cannot be nil")
	}

	total := 0
	for _, row := range rows {
		total += row.MaintenanceEvents
	}
	return total, nil
}
