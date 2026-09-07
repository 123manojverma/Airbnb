MIGRATIONS_FOLDER=db/migrations
DB_URL=auth_user:Manoj123@tcp(127.0.0.1:3306)/auth_dev

# Create a new migration -> gmake migrate-create name="create_entity_table"
migrate-create:
	goose -dir $(MIGRATIONS_FOLDER) create $(name) sql
	
migrate-up: # gmake migrate-up
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up
	
migrate-up: # gmake migrate-down
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down

# Rollback all migrations and reset database # gmake migrate-reset
migrate-reset:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" reset

# Show current migration status # gmake migrate-status
migrate-status:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" status

# Redo last migration (Down then Up) # gmake migrate-redo
migrate-redo:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" redo

# Run specific migration version # gmake migrate-version version=20200101120000
migrate-to:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" up-to $(version)

# Rollback to a specific migration version # gmake migrate-down-to version=20200101120000
migrate-down-to:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" down-to $(version)

# Force a specific migration version # gmake migrate-force version=20200101120000
migrate-force:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" force $(version)

# Print Goose help # gmake migrate-help
migrate-help:
	goose -h


# Tasks 1: Write a service layer code, that can encrypt the user password before passing to the repository layer.
# Task 2: Write a service layer function that can create a new JWT token for a user