include ./.env
DBURL=postgresql://$(PG_USER):$(PG_PWD)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=disable
MIGRATIONPATH=db/migrations
SEEDSPATH=db/seeds

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONPATH) -seq create_$(NAME)_table

migrate-up:
	migrate -database $(DBURL) -path $(MIGRATIONPATH) up

migrate-down:
	migrate -database $(DBURL) -path $(MIGRATIONPATH) down $(s)