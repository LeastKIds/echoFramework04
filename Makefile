-include .env

GO ?= go
MIGRATE ?= migrate
AIR ?= air

POSTGRESQL_URL = ${DB_DRIVER}://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}
MIGRATIONS_DIR = ./db/migrations

start:
	@$(AIR)

migrations-up:
	$(MIGRATE) -database $(POSTGRESQL_URL) -path $(MIGRATIONS_DIR) up

migrations-down:
	@$(MIGRATE) -database $(POSTGRESQL_URL) -path $(MIGRATIONS_DIR) down

migrations-create:
	@$(MIGRATE) create -ext sql -dir ./db/migrations -seq init_schema