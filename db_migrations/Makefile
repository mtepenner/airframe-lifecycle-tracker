MIGRATIONS_PATH := ./db_migrations
DATABASE_URL ?= postgres://postgres:postgres@localhost:5432/airframe?sslmode=disable

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1

migrate-force:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" force $(VERSION)
