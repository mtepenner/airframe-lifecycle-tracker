package repository

import (
	"context"
	"fmt"

	"github.com/mtepenner/airframe-lifecycle-tracker/analytics_engine/internal/models"
)

func (r *Repository) GetAirframeHistory(ctx context.Context, tailNumber string) ([]models.AirframeSnapshot, error) {
	query := `
		SELECT tail_number, model, operator, livery, seating_capacity, valid_from, valid_to
		FROM airframe_configurations
		WHERE tail_number = $1
		ORDER BY valid_from ASC`

	rows, err := r.pool.Query(ctx, query, tailNumber)
	if err != nil {
		return nil, fmt.Errorf("query airframe history: %w", err)
	}
	defer rows.Close()

	out := make([]models.AirframeSnapshot, 0)
	for rows.Next() {
		var row models.AirframeSnapshot
		if err := rows.Scan(
			&row.TailNumber,
			&row.Model,
			&row.Operator,
			&row.Livery,
			&row.SeatingCapacity,
			&row.ValidFrom,
			&row.ValidTo,
		); err != nil {
			return nil, fmt.Errorf("scan airframe history row: %w", err)
		}
		out = append(out, row)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate airframe history rows: %w", err)
	}
	return out, nil
}
