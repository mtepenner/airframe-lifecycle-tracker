# Airframe Lifecycle Tracker

Airframe Lifecycle Tracker is a full-stack analytics system for tracking immutable airframe facts, SCD Type 2 configuration history, and maintenance downtime trends for aircraft models such as MD-11, Boeing 737, and DC-8.

## Blueprint Coverage

The repository now includes the complete architecture described in the blueprint:

- Go analytics API with concurrent service handlers
- PostgreSQL schema and migration set for SCD Type 2 and maintenance facts
- Materialized CUBE analytics view
- React plus TypeScript dashboard with timeline scrubber, downtime data grid, and maintenance heatmap
- Docker and Docker Compose local environment
- Kubernetes manifests for database, API, dashboard, and ingress
- GitHub Actions workflows for API tests, migration validation, and dashboard build

## Project Structure

- analytics_engine: Go API, repository queries, service logic, models, and router
- db_migrations: ordered SQL migrations with seed data and materialized view creation
- fleet_dashboard: React plus TypeScript dashboard
- infrastructure: Kubernetes manifests for runtime deployment
- .github/workflows: CI pipelines

## Local Development

### Prerequisites

- Docker
- Docker Compose

### Start everything

```bash
make up
```

Services:

- Dashboard: http://localhost:5173
- API: http://localhost:8080
- Health: http://localhost:8080/health

### Stop everything

```bash
make down
```

### Stream logs

```bash
make logs
```

## API Endpoints

- GET /health
- GET /api/v1/fleet?year=2025&model=MD-11&operator=Atlas%20Air
- GET /api/v1/fleet/{tailNumber}/history
- GET /api/v1/analytics/downtime?from=2025-01-01&to=2025-12-31&model=MD-11

## Database and Analytics

Migrations create and seed:

- airframes immutable base table
- airframe_configurations SCD Type 2 history table
- maintenance_logs fact table
- mv_fleet_cube materialized view for heavy analytical rollups

## CI Pipelines

- test-go-api.yml: runs Go module download and go test
- test-migrations.yml: applies migrations to Postgres and validates the analytics view
- build-dashboard.yml: installs frontend dependencies and runs production build

## License

MIT. See LICENSE.
