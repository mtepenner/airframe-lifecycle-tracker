COMPOSE := docker compose

up:
	$(COMPOSE) up --build -d

down:
	$(COMPOSE) down -v

logs:
	$(COMPOSE) logs -f

migrate-up:
	$(COMPOSE) run --rm migrations

test-go:
	cd analytics_engine && go test ./...

test-dashboard:
	cd fleet_dashboard && npm run build
