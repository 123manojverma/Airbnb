MIGRATIONS_FOLDER=db/migrations
DB_URL=auth_user:Manoj123@tcp(127.0.0.1:3306)/auth_dev
GOOSE=$(HOME)/go/bin/goose

# Create a new migration -> gmake migrate-create name="create_entity_table"
migrate-create:
	$(GOOSE) -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" create $(name) sql

migrate-up: # gmake migrate-up
	$(GOOSE) -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up

migrate-down: # gmake migrate-down
	$(GOOSE) -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down

# Rollback all migrations and reset database # gmake migrate-reset
migrate-reset:
	$(GOOSE) -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" reset

# Show current migration status # gmake migrate-status
migrate-status:
	$(GOOSE) -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" status

# Redo last migration (Down then Up) # gmake migrate-redo
migrate-redo:
	$(GOOSE) -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" redo

# Run specific migration version # gmake migrate-to version=20200101120000
migrate-to:
	$(GOOSE) -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up-to $(version)

# Rollback to a specific migration version # gmake migrate-down-to version=20200101120000
migrate-down-to:
	$(GOOSE) -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down-to $(version)

# Force a specific migration version # gmake migrate-force version=20200101120000
migrate-force:
	$(GOOSE) -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" force $(version)

# Print Goose help # gmake migrate-help
migrate-help:
	$(GOOSE) -h


# Tasks 1: Write a service layer code, that can encrypt the user password before passing to the repository layer.
# Task 2: Write a service layer function that can create a new JWT token for a user