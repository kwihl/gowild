ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Migrate the database up
goose_up:
	goose -dir=migrations postgres "user=$${DB_USERNAME} host=$${DB_HOST} port=$${DB_PORT} dbname=$${DB_NAME} password=$${DB_PASSWORD} sslmode=disable" up

# Migrate the database down
goose_down:
	goose -dir=migrations postgres "user=$${DB_USERNAME} host=$${DB_HOST} port=$${DB_PORT} dbname=$${DB_NAME} password=$${DB_PASSWORD} sslmode=disable" up

# Clear database volume
clean_pg_data:
	rm -r pg_data/*