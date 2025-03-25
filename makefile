ifneq (,$(wildcard ./.env))
    include .env
    export
endif

goose_up:
	goose -dir=migrations postgres "user=$${DB_USERNAME} host=$${DB_HOST} port=$${DB_PORT} dbname=$${DB_NAME} password=$${DB_PASSWORD} sslmode=disable" up

goose_down:
	goose -dir=migrations postgres "user=$${DB_USERNAME} host=$${DB_HOST} port=$${DB_PORT} dbname=$${DB_NAME} password=$${DB_PASSWORD} sslmode=disable" up

clean_pg_data:
	rm -r pg_data/*