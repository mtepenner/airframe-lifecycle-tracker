package models

import "time"

type AirframeSnapshot struct {
	TailNumber      string    `json:"tailNumber"`
	Model           string    `json:"model"`
	Operator        string    `json:"operator"`
	Livery          string    `json:"livery"`
	SeatingCapacity int       `json:"seatingCapacity"`
	ValidFrom       time.Time `json:"validFrom"`
	ValidTo         time.Time `json:"validTo"`
}

type FleetCubeRow struct {
	Model              string `json:"model"`
	Operator           string `json:"operator"`
	Year               int    `json:"year"`
	TotalDowntimeHours int    `json:"totalDowntimeHours"`
	MaintenanceEvents  int    `json:"maintenanceEvents"`
}

type DowntimePoint struct {
	Date       time.Time `json:"date"`
	Hours      int       `json:"hours"`
	TailNumber string    `json:"tailNumber"`
	Model      string    `json:"model"`
}

type FleetAnalyticsResponse struct {
	At          time.Time      `json:"at"`
	Rows        []FleetCubeRow `json:"rows"`
	TotalEvents int            `json:"totalEvents"`
}

type DowntimeResponse struct {
	From       time.Time       `json:"from"`
	To         time.Time       `json:"to"`
	Points     []DowntimePoint `json:"points"`
	TotalHours int             `json:"totalHours"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
